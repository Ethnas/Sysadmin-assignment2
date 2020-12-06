variable "location_name" {
    type = string
    default = "North Europe"
}

variable "username" {
    type = string
    default = "azureuser"
}

variable "address_space" {
    type = string
    description = "The address space of the virual network"
    default = "10.0.0.0/16"
}

variable "subnet_address_prefixes" {
    type = string
    description = "The address prefixes for the subnet"
    default = "10.0.2.0/24"
}

variable "publicip_number" {
    type = number
    description = "The number of public ips you want"
    #default = 3
}

variable "webserver_instance_number" {
    type = number
    description = "The number of webserver VM instances you want created"
    #default = 2
}

variable "client_instance_number" {
    type = number
    description = "The number of client VM instances you want created"
    default = 1
}

variable "network_interface_number" {
    type = number
    description = "The number of network interfaces you want"
    #default = 3
}

variable "vm_size" {
    type = string
    description = "The size of the VM(s) to be created"
    default = "Standard_DS1_v2"
}