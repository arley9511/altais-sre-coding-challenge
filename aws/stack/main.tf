module "vpc" {
  source = "./../services/network/vpc"

  vpc = var.vpc
}

module "subnet" {
  source = "./../services/network/subnets"

  depends_on = [module.vpc]

  vpc = var.vpc

  # using the outputs form the vpc to create the required subnets
  gateway_info = module.vpc.gateway_info
  vpc_info = module.vpc.vpc_info
}

module ec2 {
  source = "./../services/compute/ec2"

  depends_on = [module.subnet]

  ec2 = var.ec2

  security_groups = module.vpc.security_groups
  subnets = module.subnet.subnets
}


module "s3_with_trigger" {
  source = "./../services/storage/s3_with_trigger"

  s3_with_trigger = var.s3_with_trigger
}

module "cls_lb" {
  source = "./../services/compute/load_balancing/clasic"

  depends_on = [module.ec2]

  subnets = module.subnet.subnets
  instances = module.ec2.instances
  load_balancers = var.load_balancers
  security_groups = module.vpc.security_groups
}
