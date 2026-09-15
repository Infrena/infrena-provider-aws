# Virtual Private Cloud (VPC)

The AWS VPC service provides networking components to isolate and control your infrastructure. This guide covers the core resources for creating and configuring VPCs, subnets, gateways, routing, and network security.

## aws.vpc

Creates a virtual private cloud (VPC) — the foundation for your AWS network infrastructure.

**What it's for:** Defines the IP address space for your entire network in a region.

**Key attributes:**
- `cidr` — the IPv4 CIDR block (e.g., `10.0.0.0/16`); AWS canonicalizes the value
- `enable_dns_hostnames` — whether instances get DNS hostnames (default: false for non-default VPCs)
- `enable_dns_support` — whether DNS resolution is supported (default: true)
- `tags` — resource tags as a map

**Example:**
```yaml
vpc:
  type: aws.vpc
  cidr: 10.0.0.0/16
  enable_dns_support: true
  enable_dns_hostnames: true
  tags:
    Name: my-app
```

## aws.subnet

Creates a subnet within a VPC — a subdivision of the VPC's IP address space in a single availability zone.

**What it's for:** Segments a VPC into smaller networks, each tied to one AZ and managed independently.

**Key attributes:**
- `vpc` — the VPC this subnet belongs to (reference the whole resource)
- `cidr` — the IPv4 CIDR block allocated to this subnet
- `az` — the availability zone (e.g., `us-east-1a`); omit to let AWS choose
- `map_public_ip_on_launch` — whether instances in this subnet get public IPs (default: false)
- `tags` — resource tags

**Example:**
```yaml
public_subnet:
  type: aws.subnet
  vpc: ${vpc}
  cidr: 10.0.1.0/24
  az: us-east-1a
  map_public_ip_on_launch: true
  tags:
    Name: public-1a

private_subnet:
  type: aws.subnet
  vpc: ${vpc}
  cidr: 10.0.2.0/24
  az: us-east-1b
```

## aws.internetgateway

Creates an internet gateway — the resource that enables communication between a VPC and the public internet.

**What it's for:** Acts as a target for routes and allows instances in the VPC to reach the internet.

**Key attributes:**
- `tags` — resource tags

**Attachment:** An internet gateway must be attached to the VPC using a route that references it. See `aws.ec2.route` below.

**Example:**
```yaml
igw:
  type: aws.internetgateway
  tags:
    Name: main-igw
```

## aws.routetable

Creates a route table — a set of rules (called routes) that determine where network traffic is directed.

**What it's for:** Associates routes with subnets to control traffic flow.

**Key attributes:**
- `vpc_id` — the VPC this route table belongs to
- `tags` — resource tags

**Note:** Routes are added as separate `aws.ec2.route` resources.

**Example:**
```yaml
public_routes:
  type: aws.routetable
  vpc_id: ${vpc}
  tags:
    Name: public-routes
```

## aws.ec2.route

Creates a route — a rule that defines traffic destinations and targets.

**What it's for:** Directs traffic destined for specific IP ranges to specific targets (internet gateway, NAT gateway, instance, etc.).

**Key attributes:**
- `route_table_id` — the route table this route belongs to
- `destination_cidr_block` — the IPv4 CIDR range this route matches (e.g., `0.0.0.0/0` for all traffic)
- `gateway_id` — the internet gateway ID (for routes to the internet)
- `nat_gateway_id` — the NAT gateway ID (for routes through a NAT)
- Use only one destination and one target per route

**Example:**
```yaml
default_route:
  type: aws.ec2.route
  route_table_id: ${public_routes}
  destination_cidr_block: 0.0.0.0/0
  gateway_id: ${igw.InternetGatewayId}

private_nat_route:
  type: aws.ec2.route
  route_table_id: ${private_routes}
  destination_cidr_block: 0.0.0.0/0
  nat_gateway_id: ${nat}
```

## aws.natgateway

Creates a NAT (Network Address Translation) gateway — allows instances in private subnets to reach the internet while remaining unreachable from it.

**What it's for:** Provides outbound internet access for private instances without exposing them to inbound traffic.

**Key attributes:**
- `subnet_id` — the subnet where the NAT gateway runs (typically public)
- `allocation_id` — the allocation ID of an Elastic IP; required for public NAT gateways
- `connectivity_type` — `public` (default) or `private`
- `tags` — resource tags

**Example:**
```yaml
nat:
  type: aws.natgateway
  subnet_id: ${public_subnet}
  allocation_id: ${eip.AllocationId}
  tags:
    Name: nat-1a
```

## aws.eip

Creates an Elastic IP address — a public IPv4 address that can be allocated and associated with resources.

**What it's for:** Provides a static public IP for internet-facing resources or for use with NAT gateways.

**Key attributes:**
- `domain` — set to `vpc` for VPC-based resources (default)
- `tags` — resource tags

**Example:**
```yaml
eip:
  type: aws.eip
  domain: vpc
  tags:
    Name: nat-eip
```

## aws.securitygroup

Creates a security group — a virtual firewall that controls inbound and outbound traffic to resources.

**What it's for:** Defines allow rules for traffic entering (ingress) and leaving (egress) associated instances.

**Key attributes:**
- `description` — the security group description (required)
- `vpc_id` — the VPC this security group belongs to
- `ingress` — list of inbound rules, each with:
  - `ip_protocol` — the protocol (e.g., `tcp`, `udp`, `icmp`)
  - `from_port` — the start port (or -1 for ICMP)
  - `to_port` — the end port
  - `cidr_ip` — the source CIDR range (e.g., `0.0.0.0/0` for anywhere)
  - `source_security_group_id` — alternatively, reference another security group
- `egress` — list of outbound rules (same structure); by default, VPC security groups allow all outbound traffic
- `tags` — resource tags

**Example:**
```yaml
web_sg:
  type: aws.securitygroup
  description: Security group for web servers
  vpc_id: ${vpc}
  ingress:
    - ip_protocol: tcp
      from_port: 80
      to_port: 80
      cidr_ip: 0.0.0.0/0
    - ip_protocol: tcp
      from_port: 443
      to_port: 443
      cidr_ip: 0.0.0.0/0
  tags:
    Name: web-sg

app_sg:
  type: aws.securitygroup
  description: Security group for app servers
  vpc_id: ${vpc}
  ingress:
    - ip_protocol: tcp
      from_port: 8080
      to_port: 8080
      source_security_group_id: ${web_sg.GroupId}
  tags:
    Name: app-sg
```

## Common Pitfalls

- **Unattached internet gateway:** Creating an internet gateway without adding a route to it does nothing. Add an `aws.ec2.route` with `gateway_id: ${igw}` to the public route table.
- **Missing NAT gateway Elastic IP:** Public NAT gateways require an allocation ID. Create an `aws.eip` first and reference its `AllocationId`.
- **Routes without route table association:** Subnets are not automatically associated with route tables. Create a `aws.routetable` and then associate subnets with an `aws.subnetroutetableassociation` or (if using a subnet's `route_table_id` attribute, where supported) specify it directly on the subnet.
- **Incorrect CIDR blocks:** AWS canonicalizes CIDR blocks (e.g., `10.0.0.18/18` becomes `10.0.0.0/18`). Ensure your subnets' CIDR blocks fall within the VPC's range.
- **Default egress rule:** If you add any egress rules to a security group, the default allow-all is replaced. Explicitly allow traffic you need.

## Reference Pages

- [aws.vpc](../reference/ec2/vpc.md)
- [aws.subnet](../reference/ec2/subnet.md)
- [aws.internetgateway](../reference/ec2/internetgateway.md)
- [aws.routetable](../reference/ec2/routetable.md)
- [aws.ec2.route](../reference/ec2/ec2-route.md)
- [aws.natgateway](../reference/ec2/natgateway.md)
- [aws.eip](../reference/ec2/eip.md)
- [aws.securitygroup](../reference/ec2/securitygroup.md)
