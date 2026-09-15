# Elastic Compute Cloud (EC2)

The AWS EC2 service provides computing resources from simple virtual machines to complex auto-scaling fleets. This guide covers launching instances, defining templates, and managing storage.

## aws.ec2.instance

Launches an EC2 instance — a virtual server in the cloud.

**What it's for:** Provides a computing environment where you run applications and services.

**Key attributes:**
- `image_id` (or `ami`) — the Amazon Machine Image ID to launch (e.g., `ami-0c55b159cbfafe1f0`); required unless using a launch template
- `instance_type` — the instance type (e.g., `t3.micro`, `t3.small`, `m5.large`); specifies CPU, memory, and network performance
- `subnet_id` — the subnet where the instance launches; required for VPC instances
- `key_name` — the EC2 key pair to use for SSH access (reference the key pair's name)
- `security_group_ids` — list of security group IDs to attach (e.g., `[${web_sg}]`)
- `private_ip_address` — the private IP address within the subnet; AWS assigns one if omitted
- `user_data` — a script or cloud-init data to run at startup
- `tags` — resource tags

**Example:**
```yaml
web_server:
  type: aws.ec2.instance
  image_id: ami-0c55b159cbfafe1f0
  instance_type: t3.small
  subnet_id: ${public_subnet}
  key_name: ${keypair.KeyName}
  security_group_ids:
    - ${web_sg}
  user_data: |
    #!/bin/bash
    echo "Hello from $(hostname)" > /var/www/html/index.html
  tags:
    Name: web-server-1
```

## aws.launchtemplate

Creates a launch template — a reusable blueprint for launching EC2 instances.

**What it's for:** Encapsulates instance configuration (AMI, type, security groups, storage, etc.) so you can launch identical instances without repeating settings.

**Key attributes:**
- `launch_template_name` — the name of the template
- `launch_template_data` — a map containing the configuration; see AWS documentation for the full schema, but commonly includes:
  - `ImageId` — the AMI ID
  - `InstanceType` — the instance type
  - `KeyName` — the key pair name
  - `SecurityGroupIds` — list of security group IDs
  - `UserData` — base64-encoded startup script (for YAML, use `#base64:<content>` syntax or pass raw; see CloudFormation documentation)

**Example:**
```yaml
app_template:
  type: aws.launchtemplate
  launch_template_name: app-server-template
  launch_template_data:
    ImageId: ami-0c55b159cbfafe1f0
    InstanceType: t3.small
    KeyName: ${keypair.KeyName}
    SecurityGroupIds:
      - ${app_sg.GroupId}
```

## aws.keypair

Creates or imports an EC2 key pair — a credential for accessing EC2 instances via SSH.

**What it's for:** Generates or stores a public key that AWS uses to verify your private key during SSH authentication.

**Key attributes:**
- `key_name` — the name of the key pair (required, unique per region)
- `key_type` — the type (`rsa` or `ed25519`); default is `rsa`
- `public_key_material` — if importing an existing key, the public key material in OpenSSH format; if omitted, AWS generates a new key pair (and you cannot retrieve the private key later)
- `tags` — resource tags

**Example (create a new key):**
```yaml
keypair:
  type: aws.keypair
  key_name: my-key
  key_type: ed25519
  tags:
    Name: my-key
```

**Example (import an existing public key):**
```yaml
imported_keypair:
  type: aws.keypair
  key_name: my-existing-key
  public_key_material: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKz... user@host
```

**Note:** After creating a new key pair via infrena, you must download the private key material separately from the AWS console or use `aws ec2 get-console-output`. Consider using SSM Session Manager or a bastion host as an alternative.

## aws.ec2.volume

Creates an EBS volume — block storage that can be attached to EC2 instances.

**What it's for:** Provides persistent storage independent of instance lifecycle; data survives instance termination if the volume is not set to delete on termination.

**Key attributes:**
- `size` — the volume size in GiB
- `volume_type` — the volume type (`gp2`, `gp3`, `io1`, `io2`, `st1`, `sc1`, `standard`); `gp3` is general-purpose and recommended for most workloads
- `availability_zone` — the AZ where the volume is created; must match the instance's AZ
- `iops` — I/O operations per second; required for `io1`/`io2`, optional for `gp3`
- `throughput` — throughput in MiB/s; for `gp3` only, default 125
- `encrypted` — whether to encrypt the volume; default false
- `tags` — resource tags

**Example:**
```yaml
data_volume:
  type: aws.ec2.volume
  size: 100
  volume_type: gp3
  availability_zone: us-east-1a
  iops: 3000
  throughput: 125
  tags:
    Name: app-data-vol
```

**Note:** To attach a volume to an instance, create an `aws.volumeattachment` resource or specify volumes on the instance's configuration. The examples here show stand-alone volumes.

## Common Pitfalls

- **Missing AMI ID:** You must specify an `image_id` unless using a launch template. Placeholder IDs like `ami-0123456789abcdef0` may not exist in your region; use actual AMI IDs or filter by name.
- **Mismatched AZs:** An instance's subnet (which has a fixed AZ) and any EBS volumes must be in the same AZ. Infrena will reject or AWS will fail the request if they don't match.
- **Lost key pairs:** If you create a new key pair via `aws.keypair` and do not download the private key before infrena completes, the private key is lost and cannot be recovered. Use an existing key or save the generated key immediately.
- **Security group IDs vs. names:** In a VPC, you must reference security group IDs, not names. Use `${sg.GroupId}` or pass the whole resource where the reference is declared.
- **Instance size mismatch:** Ensure the instance type you choose is available in the subnet's AZ and region.

## Reference Pages

- [aws.ec2.instance](../reference/ec2/ec2-instance.md)
- [aws.launchtemplate](../reference/ec2/launchtemplate.md)
- [aws.keypair](../reference/ec2/keypair.md)
- [aws.ec2.volume](../reference/ec2/ec2-volume.md)
