# VPC Module

Creates a foundational VPC with public and private subnets across two availability zones.

## Resources

- VPC with configurable CIDR
- 2 public subnets (one per AZ, with public IP assignment enabled)
- 2 private subnets (one per AZ)
- Internet Gateway
- Public route table with route to IGW
- Private route table
- Route table associations for all subnets

## Inputs

| Input | Type | Default | Description |
|-------|------|---------|-------------|
| `cidr` | string | `10.0.0.0/16` | CIDR block for the VPC |
| `azs` | list | `[us-east-1a, us-east-1b]` | List of two availability zones |
| `name` | string | `vpc` | Name prefix for resources |

## Outputs

| Output | Description |
|--------|-------------|
| `vpc_id` | The VPC ID |
| `public_subnet_1_id` | First public subnet ID |
| `public_subnet_2_id` | Second public subnet ID |
| `private_subnet_1_id` | First private subnet ID |
| `private_subnet_2_id` | Second private subnet ID |

## Usage

```yaml
modules:
  - ./modules/vpc

resources:
  network:
    type: module.vpc
    cidr: 10.0.0.0/16
    azs:
      - us-west-2a
      - us-west-2b
    name: my-vpc
```

Then reference outputs like `${network.vpc_id}` or `${network.public_subnet_1_id}`.
