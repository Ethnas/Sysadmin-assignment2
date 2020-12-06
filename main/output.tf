output "webserver0_public_ip" {
    value = azurerm_linux_virtual_machine.webserver.0.public_ip_address
}

output "webserver1_public_ip" {
    value = azurerm_linux_virtual_machine.webserver.1.public_ip_address
}

output "client_public_ip" {
    value = azurerm_linux_virtual_machine.client.*.public_ip_address
}