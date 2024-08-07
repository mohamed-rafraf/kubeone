disable_kubeapi_loadbalancer = true
subnets_cidr                 = 27

# Use smaller instances in Ireland for E2E tests
aws_region                = "eu-west-1"
control_plane_type        = "t3a.small"
control_plane_volume_size = 25
worker_type               = "t3a.small"
worker_volume_size        = 25
bastion_type              = "t3a.small"
ami_filters = {
  ubuntu = {
    owners       = ["099720109477"] # Canonical
    image_name   = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
    ssh_username = "ubuntu"
    worker_os    = "ubuntu"
  }
}
