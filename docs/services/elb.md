# Elastic Load Balancing (ELB)

The AWS ELB service distributes incoming application traffic across multiple targets to ensure high availability and fault tolerance. This guide covers Elastic Load Balancing v2 (ALB and NLB).

## aws.elasticloadbalancingv2.loadbalancer

Creates a load balancer — an Application Load Balancer (ALB) or Network Load Balancer (NLB) that distributes traffic.

**What it's for:** Routes incoming requests to a fleet of targets (instances, containers, or IPs) based on rules and protocols.

**Key attributes:**
- `subnets` — list of subnet IDs where the load balancer is deployed (one per AZ; required, or use `subnet_mappings` to assign static IPs)
- `security_groups` — list of security group IDs to apply to the load balancer
- `type_value` (or `type`) — the load balancer type: `application` (ALB, default), `network` (NLB), or `gateway` (GLB)
- `scheme` — `internet-facing` (default, public IP) or `internal` (private IP only)
- `name` — the load balancer name (32 characters max, alphanumeric and hyphens only, cannot start/end with hyphen)
- `tags` — resource tags

**Example (Application Load Balancer):**
```yaml
alb:
  type: aws.elasticloadbalancingv2.loadbalancer
  type_value: application
  subnets:
    - ${public_subnet_a.SubnetId}
    - ${public_subnet_b.SubnetId}
  security_groups:
    - ${alb_sg.GroupId}
  scheme: internet-facing
  tags:
    Name: main-alb
```

## aws.elasticloadbalancingv2.targetgroup

Creates a target group — a logical grouping of targets that receive traffic from the load balancer.

**What it's for:** Defines how the load balancer checks target health, what protocol and port to use, and groups instances or services together.

**Key attributes:**
- `port` — the port on targets receive traffic (e.g., 80, 443, 8080)
- `protocol` — the protocol: `HTTP`, `HTTPS`, `TCP`, `TLS`, `UDP`, `TCP_UDP`, `GRPC`, or `GENEVE`
- `vpc_id` — the VPC where targets run
- `target_type` — `instance` (EC2 instances), `ip` (IP addresses), `lambda` (Lambda functions), or `alb` (another ALB)
- `health_check_enabled` — whether health checks are performed (default: true for instance/ip/alb, false for lambda)
- `health_check_protocol` — protocol for health checks (e.g., `HTTP`, `TCP`)
- `health_check_path` — for HTTP/HTTPS checks, the endpoint path (default: `/`)
- `health_check_port` — the port for health checks (default: the traffic port)
- `health_check_interval_seconds` — seconds between checks (default: 30)
- `health_check_timeout_seconds` — seconds to wait for a response (default: 5)
- `healthy_threshold_count` — consecutive successful checks to mark healthy (default: 5)
- `unhealthy_threshold_count` — consecutive failed checks to mark unhealthy (default: 2)
- `name` — the target group name (32 characters max, alphanumeric, hyphens, and underscores)
- `tags` — resource tags

**Example:**
```yaml
app_targets:
  type: aws.elasticloadbalancingv2.targetgroup
  protocol: HTTP
  port: 80
  vpc_id: ${vpc}
  target_type: instance
  health_check_protocol: HTTP
  health_check_path: /health
  health_check_interval_seconds: 30
  healthy_threshold_count: 3
  unhealthy_threshold_count: 2
  tags:
    Name: app-targets
```

## aws.elasticloadbalancingv2.listener

Creates a listener — a rule that checks for incoming traffic on a specific port and protocol, then forwards it to a target group.

**What it's for:** Binds a port and protocol on the load balancer to a target group and defines how traffic is routed.

**Key attributes:**
- `load_balancer_arn` — the ARN of the load balancer (reference the load balancer resource)
- `port` — the listener port on the load balancer (e.g., 80, 443)
- `protocol` — the listener protocol: `HTTP`, `HTTPS`, `TCP`, `TLS`, `UDP`, `TCP_UDP`, `GRPC`, or `GENEVE`
- `default_actions` — list of actions to perform on matching traffic; typically one action:
  - `type` — `forward` (most common), `redirect`, `fixed-response`, or `authenticate-oidc`
  - `target_group_arn` — the ARN of the target group (for `forward` action)
- `ssl_policy` — for `HTTPS` or `TLS`, the security policy (e.g., `ELBSecurityPolicy-TLS-1-2-2017-01`)
- `certificates` — for `HTTPS` or `TLS`, list of certificate ARNs
- `tags` — resource tags

**Example (HTTP listener):**
```yaml
http_listener:
  type: aws.elasticloadbalancingv2.listener
  load_balancer_arn: ${alb}
  port: 80
  protocol: HTTP
  default_actions:
    - type: forward
      target_group_arn: ${app_targets.TargetGroupArn}
```

**Example (HTTPS listener with redirect from HTTP):**
```yaml
https_listener:
  type: aws.elasticloadbalancingv2.listener
  load_balancer_arn: ${alb.LoadBalancerArn}
  port: 443
  protocol: HTTPS
  ssl_policy: ELBSecurityPolicy-TLS-1-2-2017-01
  certificates:
    - certificate_arn: ${certificate.CertificateArn}
  default_actions:
    - type: forward
      target_group_arn: ${app_targets.TargetGroupArn}

http_redirect:
  type: aws.elasticloadbalancingv2.listener
  load_balancer_arn: ${alb.LoadBalancerArn}
  port: 80
  protocol: HTTP
  default_actions:
    - type: redirect
      redirect_config:
        protocol: HTTPS
        port: "443"
        status_code: HTTP_301
```

## Common Pitfalls

- **Load balancer in wrong subnets:** The load balancer must span multiple AZs for high availability. Specify subnets from different AZs.
- **Target group without targets:** A target group with no registered targets will mark all targets as unhealthy. After creating the target group, list the instances or IP addresses in the target group's `targets` attribute.
- **Health check timeouts:** If health checks repeatedly fail, check that the target port and path are correct, the security group allows traffic from the load balancer, and instances are running.
- **Missing listener:** A load balancer with no listeners cannot route traffic. Create at least one `aws.elasticloadbalancingv2.listener`.
- **Certificate mismatches:** For HTTPS listeners, the certificate's domain must match the request hostname, or clients will see warnings. Use a wildcard certificate or multiple certificates for different domains.
- **Protocol mismatch:** Ensure the listener protocol matches what your application supports. ALBs support HTTP/HTTPS, NLBs support TCP/TLS/UDP/TCP_UDP, and GLBs support GENEVE.

## Reference Pages

- [aws.elasticloadbalancingv2.loadbalancer](../reference/elasticloadbalancingv2/elasticloadbalancingv2-loadbalancer.md)
- [aws.elasticloadbalancingv2.targetgroup](../reference/elasticloadbalancingv2/elasticloadbalancingv2-targetgroup.md)
- [aws.elasticloadbalancingv2.listener](../reference/elasticloadbalancingv2/elasticloadbalancingv2-listener.md)
