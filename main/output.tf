output "webserver0_public_ip" {
    value = azurerm_linux_virtual_machine.webserver.0.public_ip_address
}

output "webserver1_public_ip" {
    value = azurerm_linux_virtual_machine.webserver.1.public_ip_address
}

output "client_public_ip" {
    value = azurerm_linux_virtual_machine.client.*.public_ip_address
}

output "resource_group_name" {
  value = azurerm_resource_group.myterraformgroup.name
}

output "client_vm_name" {
  value = [azurerm_linux_virtual_machine.client.*.name]
}

output "webserver_vm_name" {
  value = [azurerm_linux_virtual_machine.webserver.*.name]
}

output "client_vm_size" {
  value = azurerm_linux_virtual_machine.client.0.name
}