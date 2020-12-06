package test

import (
	// "strings"
	// "testing"

	// "github.com/gruntwork-io/terratest/tree/master/modules/terraform"
	// "github.com/gruntwork-io/terratest/modules/azure"

	"fmt"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
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

	TestVMCreated(t, terraformOptions, subscriptionId)

}

func TestVMCreated(t *testing.T, terraformOptions * terraform.Options, subscriptionId string) {
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	expectedWebserverVMName := terraform.Output(t, terraformOptions, "webserver_vm_name")
	expectedClientVMName := terraform.Output(t, terraformOptions, "client_vm_name")
	expectedVMSize := compute.VirtualMachineSizeTypes(terraform.Output(t, terraformOptions, "cleint_vm_size"))

	// Check against all VM names in a Resource Group.
	vmList := azure.ListVirtualMachinesForResourceGroup(t, resourceGroupName, subscriptionID)
	expectedVMCount := 3
	assert.Equal(t, expectedVMCount, len(vmList))
	assert.Contains(t, vmList, expectedWebserverVMName)
	assert.Contains(t, vmList, expectedClientVMName)

	// Check that the all the VM IPs are correct
	var clientIp := terraform.Output(t, terraformOptions, "client_public_ip")
	var webserver0Ip := terraform.Output(t, terraformOptions, "webserver0_public_ip")
	var webserver1Ip := terraform.Output(t, terraformOptions, "webserver1_public_ip")

	vmsByRef := azure.GetVirtualMachinesForResourceGroup(t, resourceGroupName, subscriptionID)
	thisVM := vmsByRef[expectedVMName]
	assert.Equal(t, expectedClientVMName, thisVM.HardwareProfile.VMSize)
}
