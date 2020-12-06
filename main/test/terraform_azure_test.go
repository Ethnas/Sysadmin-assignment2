package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/tree/master/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/azure"
)

// Medium-level tests for the servers set up by Terraform
func TestTerraformAzureMidLevel(t *testing.T) {
	t.Parallel()

	subscriptionId := "93d67d1d-09d3-4cca-9b39-7cd1ef68c9dd"

	// Variables for Terraform
	var numberPublicIp := 3
	var numberWebserver := 2
	var numberClient := 1
	var numberNetworkInterface := numberClient + numberWebserver

	terraformOptions := &terraform.Options{
		// The path where the Terraform code is located
		TerraformDir: "../",
		// The  variables needed to run the Terraform code
		Vars: map[string]interface{}{
			"publicip_number":           numberPublicIp,
			"webserver_instance_number": numberWebserver,
			"client_instance_number":    numberClient,
			"network_interface_number":  numberNetworkInterface,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	var clientIp := terraform.Output(t, terraformOptions, "client_public_ip")
	var webserver0Ip := terraform.Output(t, terraformOptions, "webserver0_public_ip")
	var webserver1Ip := terraform.Output(t, terraformOptions, "webserver1_public_ip")

}
