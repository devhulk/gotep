To create a new VPC in AWS and create a VM in that VPC, you will need to create two files: a main.tf and an output.tf. 

The main.tf file will contain the code to create the VPC and the VM. The code should look like this:

```
# Create a VPC
resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support = true
  tags = {
    Name = "example-vpc"
  }
}

# Create a subnet
resource "aws_subnet" "example" {
  vpc_id = aws_vpc.example.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "us-east-1a"
  tags = {
    Name = "example-subnet"
  }
}

# Create a security group
resource "aws_security_group" "example" {
  name = "example-security-group"
  description = "Allow SSH and HTTP access"
  vpc_id = aws_vpc.example.id

  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Create an instance
resource "aws_instance" "example" {
  ami = "ami-0d5d9d301c853a04a"
  instance_type = "t2.micro"
  subnet_id = aws_subnet.example.id
  security_groups = [aws_security_group.example.id]
  tags = {
    Name = "example-instance"
  }
}
```

The output.tf file will contain the code to output the public IP address of the VM. The code should look like this:

```
output "public_ip" {
  value = aws_instance.example.public_ip
}
```

Once you have created the two files, you can run Terraform to create the VPC and the VM. The steps are as follows:

1. Initialize Terraform: `terraform init`
2. Validate the configuration: `terraform validate`
3. Create an execution plan: `terraform plan`
4. Apply the execution plan: `terraform apply`

If you are running Terraform for the first time, you will be prompted to enter a "yes" to confirm the execution plan. After you enter "yes", Terraform will create the VPC and the VM.

Once the VPC and the VM have been created, you can run `terraform output` to view the public IP address of the VM.

There are a few gotchas to be aware of when running Terraform. First, make sure that you have the correct AWS credentials configured in your environment. Secondly, make sure that you have the correct AWS region configured in your environment. Finally, make sure that you have the correct AMI ID configured in your main.tf file.