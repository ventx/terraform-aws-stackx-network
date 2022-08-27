package tests

import (
	"context"
	"fmt"
	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"testing"
)

func runAwsNetworkTest(t *testing.T) {

	// Pick a random AWS region to test in.
	// approvedRegions match with stackx regions
	approvedRegions := []string{"eu-central-1", "eu-west-2", "us-east-1", "us-west-2", "ap-east-1", "ap-southeast-1", "ap-southeast-2"}
	awsRegion := aws.GetRandomStableRegion(t, approvedRegions, nil)
	log.Printf("awsRegion: %s", awsRegion)

	var awsConfig awssdk.Config

	// Localstack check
	//
	// Waiting for:
	// https://github.com/gruntwork-io/terratest/pull/495
	// to be implemented / merged.
	//
	// Check if we can listen on port 4566 (a localstack port)
	// If YES (no localstack port found) --> Continue with additional tests
	// If NO (we cant listen, therefore localstack port was found) --> Skip additional tests
	//

	port := "4566"
	listener, listenerErr := net.Listen("tcp", "127.0.0.1:"+port)

	if listenerErr != nil {
		log.Printf("Running Localstack found (can't listen on localstack port %q): %s\n", port, listenerErr)
		log.Println("Copy localstack.tf provider configuration file to examples/localstack-provider.tf (will be removed at the end)")

		defer os.Remove("../examples/localstack-provider.tf")
		err := files.CopyFile("localstack.tf", "./../examples/localstack-provider.tf")
		if err != nil {
			log.Fatalf("Failed to copy localstack.tf to examples/localstack-provider.tf: %s", err)
		}

		log.Println("Configuring aws-sdk with localstack endpoints")
		os.Setenv("AWS_ACCESS_KEY_ID", "mocktest")
		os.Setenv("AWS_SECRET_ACCESS_KEY", "mocktest")

		customResolver := awssdk.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (awssdk.Endpoint, error) {
			if service == ec2.ServiceID {
				return awssdk.Endpoint{
					PartitionID:       "aws",
					URL:               "http://localhost:4566",
					SigningRegion:     awsRegion,
					HostnameImmutable: true,
				}, nil
			}
			return awssdk.Endpoint{}, fmt.Errorf("localstack endpoint config in AWS-SDK failed")
		})

		awsConfig, err = config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))

		if err != nil {
			log.Fatal(err)
		}
	}

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "./../examples",

		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	// At the end of the test, run `terraform destroy`
	defer terraform.Destroy(t, terraformOptions)

	// Runs `terraform init` and `terraform apply` and fails the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	numberAzs := 3

	// Run `terraform output` to get the value of an output variable
	privateSubnetIds := terraform.Output(t, terraformOptions, "private_subnet_ids")
	publicSubnetIds := terraform.Output(t, terraformOptions, "public_subnet_ids")
	vpcId := terraform.Output(t, terraformOptions, "vpc_id")
	zones := terraform.Output(t, terraformOptions, "zones")

	publicSubnetFields := strings.Fields(publicSubnetIds)
	publicSubnetsCount := len(publicSubnetFields)
	assert.True(t, publicSubnetsCount == numberAzs)

	privateSubnetFields := strings.Fields(privateSubnetIds)
	privateSubnetsCount := len(privateSubnetFields)
	assert.True(t, privateSubnetsCount == numberAzs)

	replacer := strings.NewReplacer("[", "", "]", "", "\"", "", "\n", "", " ", "")
	subnetPublID := replacer.Replace(publicSubnetIds)
	arrayPublSubnets := strings.Split(subnetPublID, ",")
	subnetPrivID := replacer.Replace(privateSubnetIds)
	arrayPrivSubnets := strings.Split(subnetPrivID, ",")

	if listenerErr == nil {
		log.Println("Running against AWS - Continue with additional tests")

		// Closing the listener we created to check if Localstack is running
		errClose := listener.Close()
		if errClose != nil {
			log.Fatalf("Error while closing port %s for testing if Localstack is running: %v", port, errClose)
		}
		subnets := aws.GetSubnetsForVpc(t, vpcId, awsRegion)

		require.Equal(t, 2*numberAzs, len(subnets))

		// Verify if the network that is supposed to be private is really private
		for i := 0; i < len(arrayPrivSubnets)-1; i++ {
			assert.False(t, aws.IsPublicSubnet(t, arrayPrivSubnets[i], awsRegion))
		}

		// Verify if the network that is supposed to be public is really public
		for i := 0; i < len(arrayPublSubnets)-1; i++ {
			assert.True(t, aws.IsPublicSubnet(t, arrayPublSubnets[i], awsRegion))
		}

		// AWS config
		var err error
		awsConfig, err = config.LoadDefaultConfig(context.TODO())

		if err != nil {
			log.Fatal(err)
		}
	}

	svc := ec2.NewFromConfig(awsConfig)

	zonesInt, err := strconv.Atoi(zones)
	if err != nil {
		log.Fatal("Error converting output zones to int")
	}

	for i := 0; i < zonesInt; i++ {
		//log.Printf("\nTesting Private Subnet: %d", i+1)
		privateSubnetId := strings.TrimSuffix(strings.TrimPrefix(privateSubnetFields[i], "["), "]")

		reqPrivate, err := svc.DescribeSubnets(context.Background(), &ec2.DescribeSubnetsInput{SubnetIds: []string{*awssdk.String(privateSubnetId)}})
		if err != nil {
			log.Printf("Failed describing subnets by ID: %s", privateSubnetId)
			log.Fatal(err)
		}

		// Get Tags from subnet
		privateSubnetTags := reqPrivate.Subnets[0].Tags

		for _, tag := range privateSubnetTags {
			//log.Printf("Private Subnet Tag: %v(%v)", *tag.Key, *tag.Value)

			if *tag.Key == "kubernetes.io/cluster/stackx" {
				assert.True(t, *tag.Key == "kubernetes.io/cluster/stackx" && *tag.Value == "shared")
			}
			if *tag.Key == "kubernetes.io/role/internal-elb" {
				assert.True(t, *tag.Key == "kubernetes.io/role/internal-elb" && *tag.Value == "1")
			}
		}

		//log.Printf("\nTesting Public Subnet: %d", i+1)
		publicSubnetId := strings.TrimSuffix(strings.TrimPrefix(publicSubnetFields[i], "["), "]")

		reqPublic, err := svc.DescribeSubnets(context.Background(), &ec2.DescribeSubnetsInput{SubnetIds: []string{*awssdk.String(publicSubnetId)}})
		if err != nil {
			log.Printf("Failed describing subnets by ID: %s", publicSubnetId)
			log.Fatal(err)
		}

		// Get Tags from subnet
		publicSubnetTags := reqPublic.Subnets[0].Tags

		for _, tag := range publicSubnetTags {
			//log.Printf("Public Subnet Tag: %v(%v)", *tag.Key, *tag.Value)

			if *tag.Key == "kubernetes.io/cluster/stackx" {
				assert.True(t, *tag.Key == "kubernetes.io/cluster/stackx" && *tag.Value == "shared")
			}
			if *tag.Key == "kubernetes.io/role/elb" {
				assert.True(t, *tag.Key == "kubernetes.io/role/elb" && *tag.Value == "1")
			}
		}
	}
}
