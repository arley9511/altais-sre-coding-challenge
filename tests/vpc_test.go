package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// An example of how to test the Terraform module in examples/terraform-aws-network-example using Terratest.
func TestTerraformAwsNetworkExample(t *testing.T) {
	t.Parallel()

	jsonData := `[
	  {
		name = "altais-test-net"
		cidr_block = "10.26.0.0/24"
		instance_tenancy = "default"
		enable_dns_support = true
		enable_dns_hostnames = true
		subnets = [
		  {
			name = "altais-test-public-1"
			cidr_block = "10.26.0.0/26"
			availability_zone = "us-east-2a"
		  },
		  {
			name = "altais-test-public-2"
			cidr_block = "10.26.0.64/26"
			availability_zone = "us-east-2a"
		  },
		  {
			name = "altais-test-private-1"
			cidr_block = "10.26.0.128/26"
			availability_zone = "us-east-2a"
		  },
		  {
			name = "altais-test-private-2"
			cidr_block = "10.26.0.192/26"
			availability_zone = "us-east-2a"
		  }
		]
		nat_gateways = [
		  {
			name = "altais-nat-public-1"
			subnet = "altais-test-public-1"
			elastic_ip = {
			  name = "altais-eip-public-1"
			  vpc = true
			}
		  }
		]
		router_tables = [
		  {
			name = "altais-test-public"
			subnets = [
			  "altais-test-public-1",
			  "altais-test-public-2"
			]
			routes = [
			  {
				cidr_block = "0.0.0.0/0"
				nat_gateway = ""
				gateway_name = "altais-test-net-internet-gateway"
				vpc_peering_connection_id = ""
			  }
			]
		  },
		  {
			name = "altais-test-private"
			subnets = [
			  "altais-test-private-1",
			  "altais-test-private-2"
			]
			routes = [
			  {
				cidr_block = "0.0.0.0/0"
				nat_gateway = "altais-nat-public-1"
				gateway_name = ""
				vpc_peering_connection_id = ""
			  }
			]
		  }
		]
		security_groups = [
		  {
			name = "altais-test-security-group"
			ingress = [
			  {
				from_port = 0
				to_port = 0
				protocol = "-1"
				cidr_blocks = [
				  "0.0.0.0/0"
				]
			  }
			]
			egress = [
			  {
				from_port = 0
				to_port = 0
				protocol = "-1"
				cidr_blocks = [
				  "0.0.0.0/0"
				]
			  }
			]
		  }
		]
		acl = [
		  {
			name    = "altais-test-acl-public"
			subnets = [
			  "altais-test-public-1",
			  "altais-test-public-2"
			]
			egress = [
			  {
				protocol   = "-1"
				rule_no    = 100
				action     = "allow"
				cidr_block = "0.0.0.0/0"
				from_port  = 0
				to_port    = 0
			  }
			]
			ingress = [
			  {
				protocol   = "-1"
				rule_no    = 100
				action     = "allow"
				cidr_block = "0.0.0.0/0"
				from_port  = 0
				to_port    = 0
			  }
			]
		  },
		  {
			name    = "altais-test-acl-private"
			subnets = [
			  "altais-test-private-1",
			  "altais-test-private-2"
			]
			egress = [
			  {
				protocol   = "-1"
				rule_no    = 100
				action     = "allow"
				cidr_block = "0.0.0.0/0"
				from_port  = 0
				to_port    = 0
			  }
			]
			ingress = [
			  {
				protocol   = "tcp"
				rule_no    = 100
				action     = "allow"
				cidr_block = "10.26.0.0/26"
				from_port  = 22
				to_port    = 22
			  },
			  {
				protocol   = "tcp"
				rule_no    = 101
				action     = "allow"
				cidr_block = "10.26.0.64/26"
				from_port  = 22
				to_port    = 22
			  },
			  {
				protocol   = "tcp"
				rule_no    = 102
				action     = "allow"
				cidr_block = "10.26.0.0/26"
				from_port  = 80
				to_port    = 80
			  },
			  {
				protocol   = "tcp"
				rule_no    = 103
				action     = "allow"
				cidr_block = "10.26.0.64/26"
				from_port  = 80
				to_port    = 80
			  },
			  {
				protocol   = "tcp"
				rule_no    = 104
				action     = "allow"
				cidr_block = "0.0.0.0/0"
				from_port  = 1024
				to_port    = 65535
			  },
			  {
				protocol   = "udp"
				rule_no    = 105
				action     = "allow"
				cidr_block = "0.0.0.0/0"
				from_port  = 123
				to_port    = 123
			  }
			]
		  }
		]
	  }
	]`

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../aws/stack",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"region":          "us-east-2",
			"profile":         "arley_tests",
			"vpc":             jsonData,
			"ec2":             `[]`,
			"s3_with_trigger": `[]`,
			"load_balancers": `[]`,
		},
	})

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndPlan(t, terraformOptions)
}
