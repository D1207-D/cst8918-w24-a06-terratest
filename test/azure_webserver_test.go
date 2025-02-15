package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var subscriptionID string = "b1b9bb0f-d71b-49d6-9f77-7c9d5cacccce"

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where your Terraform code is located
		TerraformDir: "../",
		// Override the default Terraform variables
		Vars: map[string]interface{}{
			"labelPrefix": "dani0197",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs from Terraform
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	nicName := terraform.Output(t, terraformOptions, "nic_name") // Add output for NIC if defined in Terraform

	// Test 1: Confirm VM exists
	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID), "The VM does not exist")

	// Test 2: Confirm NIC exists and is connected to VM
	assert.True(t, azure.NetworkInterfaceExists(t, nicName, resourceGroupName, subscriptionID), "The NIC does not exist")

	// Fetch the NIC and handle potential errors
	nic, err := azure.GetNetworkInterfaceE(nicName, resourceGroupName, subscriptionID)
	assert.NoError(t, err, "Failed to fetch the NIC")
	assert.NotNil(t, nic.VirtualMachine, "The NIC is not connected to a VM")

	// Test 3: Confirm VM is running the correct Ubuntu version
	vm := azure.GetVirtualMachine(t, vmName, resourceGroupName, subscriptionID)
	assert.Equal(t, "Linux", string(vm.StorageProfile.OsDisk.OsType), "The VM is not running a Linux OS")

	ubuntuVersion := "18.04-LTS" // Replace with your expected version
	assert.Contains(t, *vm.StorageProfile.ImageReference.Sku, ubuntuVersion, "The VM is not running the correct Ubuntu version")
}
