from diagrams import Cluster, Diagram
from diagrams.aws.network import VPC
from diagrams.aws.network import Endpoint
from diagrams.aws.network import InternetGateway
from diagrams.aws.network import NATGateway
from diagrams.aws.network import PrivateSubnet
from diagrams.aws.network import PublicSubnet
from diagrams.aws.network import RouteTable

with Diagram("terraform-aws-stackx-network", outformat="png", filename="screenshot1", show=False):
    with Cluster("VPC"):
        vpc = VPC("vpc")
        with Cluster("Private Subnets"):
            privatesubnet_group =  [PrivateSubnet("private-subnet-a") >> RouteTable("PrivateRT-a") >> NATGateway("NAT Gateway"),
                                    PrivateSubnet("private-subnet-b") >> RouteTable("PrivateRT-b") >> NATGateway("NAT Gateway"),
                                    PrivateSubnet("private-subnet-c") >> RouteTable("PrivateRT-c") >> NATGateway("NAT Gateway")]
            s3endpoint = Endpoint("S3 VPC Endpoint")

        with Cluster("Public Subnets"):
            publicsubnet_group =  [PublicSubnet("public-subnet-a"),
                                   PublicSubnet("public-subnet-b"),
                                   PublicSubnet("public-subnet-c")]  >> RouteTable("PublicRT") >> InternetGateway("Internet Gateway")

        with Cluster("Cache Subnets"):
            cachesubnet_group = [PrivateSubnet("cache-subnet-a"),
                                 PrivateSubnet("cache-subnet-b"),
                                 PrivateSubnet("cache-subnet-c")]

        with Cluster("DB Subnets"):
            dbsubnet_group = [PrivateSubnet("db-subnet-a"),
                              PrivateSubnet("db-subnet-b"),
                              PrivateSubnet("db-subnet-c")]
