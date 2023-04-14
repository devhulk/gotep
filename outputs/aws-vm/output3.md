You will need to create two files:

1. main.tf

This file will contain the HCL code to create the VPC and the VM.

```hcl
provider "aws" {
  region = "us-east-1"
}

resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "example" {
  vpc_id     = aws_vpc.example.id
  cidr_block = "10.0.1.0/24"
}

resource "aws_instance" "example" {
  ami           = "ami-a1b2c3d4"
  instance_type = "t2.micro"
  subnet_id     = aws_subnet.example.id
}

output "public_ip" {
  value = aws_instance.example.public_ip
}
```

2. terraform.tfvars

This file will contain the variables for the HCL code.

```hcl
region = "us-east-1"
```

To create the VPC and the VM, you will need to follow these steps:

1. Create the two files mentioned above.
2. Initialize the Terraform working directory by running the command `terraform init`.
3. Run `terraform plan` to preview the changes that will be applied to the infrastructure.
4. Run `terraform apply` to apply the changes.
5. Run `terraform output` to output the public IP address of the VM.

Note: Before running `terraform apply`, make sure that you have the necessary AWS credentials configured.