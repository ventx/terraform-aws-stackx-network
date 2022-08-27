package tests

import (
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"net"
	"os"
	"strings"
	"testing"
)

func runAwsNetworkAllTest(t *testing.T) {

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
		log.Println("Copy localstack.tf provider configuration file to examples/all/localstack-provider.tf (will be removed at the end)")

		defer os.Remove("../examples/all/localstack-provider.tf")
		err := files.CopyFile("localstack.tf", "./../examples/all/localstack-provider.tf")
		if err != nil {
			log.Fatalf("Failed to copy localstack.tf to examples/all/localstack-provider.tf: %s", err)
		}
	}

	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	// approvedRegions match with stackx ones
	approvedRegions := []string{"eu-central-1", "eu-west-2", "us-east-1", "us-west-2", "ap-east-1", "ap-southeast-1", "ap-southeast-2"}
	awsRegion := aws.GetRandomStableRegion(t, approvedRegions, nil)
	log.Printf("awsRegion: %s", awsRegion)

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "./../examples/all",

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
	publicSubnetIds := terraform.Output(t, terraformOptions, "public_subnet_ids")
	privateSubnetIds := terraform.Output(t, terraformOptions, "private_subnet_ids")
	vpcId := terraform.Output(t, terraformOptions, "vpc_id")

	publicSubnetFields := strings.Fields(publicSubnetIds)
	publicSubnetsCount := len(publicSubnetFields)
	assert.True(t, publicSubnetsCount == numberAzs)

	privateSubnetFields := strings.Fields(privateSubnetIds)
	privateSubnetsCount := len(privateSubnetFields)
	assert.True(t, privateSubnetsCount == numberAzs)

	if listenerErr == nil {
		log.Println("Running against AWS - Continue with additional tests")

		// Closing the listener we created to check if Localstack is running
		errClose := listener.Close()
		if errClose != nil {
			log.Fatalf("Error while closing port %s for testing if Localstack is running: %v", port, errClose)
		}
		subnets := aws.GetSubnetsForVpc(t, vpcId, awsRegion)

		require.Equal(t, 2*numberAzs, len(subnets))

		replacer := strings.NewReplacer("[", "", "]", "", "\"", "", "\n", "", " ", "")
		subnetPublID := replacer.Replace(publicSubnetIds)
		arrayPublSubnets := strings.Split(subnetPublID, ",")
		subnetPrivID := replacer.Replace(privateSubnetIds)
		arrayPrivSubnets := strings.Split(subnetPrivID, ",")

		// Verify if the network that is supposed to be private is really private
		for i := 0; i < len(arrayPrivSubnets)-1; i++ {
			assert.False(t, aws.IsPublicSubnet(t, arrayPrivSubnets[i], awsRegion))
		}

		// Verify if the network that is supposed to be public is really public
		for i := 0; i < len(arrayPublSubnets)-1; i++ {
			assert.True(t, aws.IsPublicSubnet(t, arrayPublSubnets[i], awsRegion))
		}
	}
}
