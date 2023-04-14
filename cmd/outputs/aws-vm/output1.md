You would need to create two files: a main.tf file and a variables.tf file.

In the main.tf file, you would need to include the following HCL code:

```
# Configure the AWS Provider
provider "aws" {
  region = var.region
}

# Create the VPC
resource "aws_vpc" "vpc" {
  cidr_block = var.vpc_cidr_block
  tags = {
    Name = "MyVPC"
  }
}

# Create the Subnet
resource "aws_subnet" "subnet" {
  vpc_id     = aws_vpc.vpc.id
  cidr_block = var.subnet_cidr_block
  tags = {
    Name = "MySubnet"
  }
}

# Create the Security Group
resource "aws_security_group" "sg" {
  name        = "MySecurityGroup"
  description = "Allow SSH access"
  vpc_id      = aws_vpc.vpc.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Create the EC2 Instance
resource "aws_instance" "vm" {
  ami           = var.ami_id
  instance_type = var.instance_type
  subnet_id     = aws_subnet.subnet.id
  vpc_security_group_ids = [aws_security_group.sg.id]

  tags = {
    Name = "MyVM"
  }
}

# Output the Public IP Address
output "public_ip" {
  value = aws_instance.vm.public_ip
}
```

In the variables.tf file, you would need to include the following HCL code:

```
variable "region" {
  description = "AWS Region"
  default     = "us-east-1"
}

variable "vpc_cidr_block" {
  description = "VPC CIDR Block"
  default     = "10.0.0.0/16"
}

variable "subnet_cidr_block" {
  description = "Subnet CIDR Block"
  default     = "10.0.1.0/24"
}

variable "ami_id" {
  description = "AMI ID"
  default     = "ami-12345678"
}

variable "instance_type" {
  description = "Instance Type"
  default     = "t2.micro"
}
```

To create the files and run Terraform, you would need to follow these steps:

1. Create a directory for your Terraform project.
2. Create the main.tf and variables.tf files in the directory.
3. Copy and paste the HCL code from above into the respective files.
4. Initialize Terraform in the project directory by running the command `terraform init`.
5. Validate your Terraform configuration by running the command `terraform validate`.
6. Plan your Terraform configuration by running the command `terraform plan`.
7. Apply your Terraform configuration by running the command `terraform apply`.
8. Output the public IP address of the VM by running the command `terraform output public_ip`.

Note: Before running the `terraform apply` command, you will need to make sure that you have the correct AWS credentials configured.
