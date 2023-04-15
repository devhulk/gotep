To create a new VPC in AWS and a VM in that VPC using Terraform, you would need to create two files: `main.tf` and `variables.tf`.

In `main.tf`, you would include the following HCL code:

```
# Declare the provider
provider "aws" {
  region     = var.aws_region
}

# Create a VPC
resource "aws_vpc" "my_vpc" {
  cidr_block = var.vpc_cidr_block
  tags = {
    Name = "my_vpc"
  }
}

# Create an internet gateway
resource "aws_internet_gateway" "my_igw" {
  vpc_id = aws_vpc.my_vpc.id
  tags = {
    Name = "my_igw"
  }
}

# Attach the internet gateway to the VPC
resource "aws_vpc_attachment" "my_vpc_attachment" {
  vpc_id             = aws_vpc.my_vpc.id
  internet_gateway_id = aws_internet_gateway.my_igw.id
}

# Create a subnet
resource "aws_subnet" "my_subnet" {
  vpc_id     = aws_vpc.my_vpc.id
  cidr_block = var.subnet_cidr_block
  tags = {
    Name = "my_subnet"
  }
}

# Create a security group
resource "aws_security_group" "my_sg" {
  name        = "my_sg"
  description = "Allow all inbound traffic"
  vpc_id      = aws_vpc.my_vpc.id

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "icmp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 0
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 0
    to_port     = 65535
    protocol    = "udp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "icmp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 65535
    protocol    = "udp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Create an EC2 instance
resource "aws_instance" "my_instance" {
  ami           = var.ami_id
  instance_type = var.instance_type
  vpc_security_group_ids = [aws_security_group.my_sg.id]
  subnet_id     = aws_subnet.my_subnet.id
  key_name      = var.key_name
  tags = {
    Name = "my_instance"
  }
}
```

In `variables.tf`, you would include the following HCL code to define the variables used in `main.tf`:

```
variable "aws_region" {
  description = "The AWS region in which to provision infrastructure"
  type        = string
}

variable "vpc_cidr_block" {
  description = "The CIDR block for the VPC"
  type        = string
}

variable "subnet_cidr_block" {
  description = "The CIDR block for the subnet"
  type        = string
}

variable "ami_id" {
  description = "The ID of the AMI to use for the EC2 instance"
  type        = string
}

variable "instance_type" {
  description = "The type of EC2 instance to create"
  type        = string
}

variable "key_name" {
  description = "The name of the key pair to use to SSH into the EC2 instance"
  type        = string
}
```

To run Terraform, you would follow these steps:

1. Create a new directory for your Terraform files.
2. Create a new file named `provider.tf` and include the following code to declare the AWS provider:

```
provider "aws" {}
```

3. Create the `main.tf` and `variables.tf` files in the same directory.
4. In `main.tf`, replace the values of the `aws_region`, `vpc_cidr_block`, `subnet_cidr_block`, `ami_id`, `instance_type`, and `key_name` variables with the appropriate values for your use case.
5. Open your terminal or command prompt and navigate to the directory containing your Terraform files.
6. Run `terraform init` to initialize your Terraform environment and download the necessary provider plugins.
7. Run `terraform plan` to see a preview of the changes that Terraform will make.
8. If the plan looks correct, run `terraform apply` to create the infrastructure.
9. After Terraform finishes creating the infrastructure, it will output the public IP address of the EC2 instance. Save this value for future use.
10. To clean up your infrastructure, run `terraform destroy`. This will remove all resources created by Terraform.

One gotcha to watch out for is making sure that your AWS credentials are properly configured. Terraform will look for your AWS access key and secret access key in environment variables or in the `~/.aws/credentials` file. If you have not configured your credentials properly, Terraform may fail to create the necessary infrastructure.