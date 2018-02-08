provider "wormly" {
  api_key = ""
}

data "wormly_hosts" "test_hosts" {}

# output "test" {
#   value = "${data.wormly_hosts.test_hosts.id}"
# }

resource "null_resource" "test" {
  provisioner "local-exec" {
    command = "echo ${data.wormly_hosts.test_hosts.id}"
  }
}
