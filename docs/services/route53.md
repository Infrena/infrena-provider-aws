# Amazon Route 53

This guide covers creating DNS hosted zones and records with Route 53, AWS's managed DNS service.

## Hosted Zone

A hosted zone contains all DNS records for a single domain. Route 53 can host public zones (internet-facing) or private zones (within VPCs).

**Resource type:** `aws.hostedzone`

**Key attributes:**
- `name`: Domain name (e.g., example.com or example.com.) — Route 53 treats trailing dots as optional
- `vpcs`: (Private zones only) List of VPC IDs and regions to associate
- `hosted_zone_config`: Map containing optional `comment` field for notes
- `hosted_zone_tags` (also: `tags`): Map of tag key-value pairs

Route 53 is a global service; hosted zones have no region attribute.

**Example (public zone):**
```yaml
zone:
  type: aws.hostedzone
  name: example.com
  hosted_zone_tags:
    Environment: dev
```

**Example (private zone):**
```yaml
private_zone:
  type: aws.hostedzone
  name: internal.example.com
  vpcs:
    - id: ${vpc_id}
      region: ${var.aws_region}
  hosted_zone_tags:
    Environment: dev
```

## Record Set

DNS records that route traffic for domain names. Records can point to AWS resources (alias records) or external IP addresses and hostnames.

**Resource type:** `aws.recordset`

**Key attributes:**
- `name`: Record domain (e.g., www.example.com or example.com for the apex)
- `type_value` (also: `type`): DNS record type (A, AAAA, CNAME, MX, TXT, SRV, NS, etc.)
- `hosted_zone_id`: ID of the hosted zone (or use `hosted_zone_name` with trailing dot)
- `resource_records`: List of values for the record (e.g., IP addresses or hostnames)
- `ttl`: Time to live in seconds (300 to 86400; default 300)
- `alias_target`: (Alias records only) Map pointing to an AWS resource (CloudFront, ELB, S3, etc.)
- `set_identifier`: For weighted or failover records; must be unique within the record set
- `weight`: (Weighted routing) Relative proportion for load balancing
- `failover`: PRIMARY or SECONDARY for active-passive failover
- `geo_location`: (Geolocation routing) Map specifying continent, country, or subdivision
- `health_check_id`: Reference to a health check for failover decisions
- `multi_value_answer`: true for multivalue answer records

**Example (A record):**
```yaml
www:
  type: aws.recordset
  name: www.example.com
  type_value: A
  hosted_zone_id: ${zone}
  resource_records:
    - 192.0.2.1
  ttl: 300
```

**Example (CNAME record):**
```yaml
blog:
  type: aws.recordset
  name: blog.example.com
  type_value: CNAME
  hosted_zone_id: ${zone}
  resource_records:
    - blog-platform.example.com
  ttl: 300
```

**Example (Alias to CloudFront):**
```yaml
cdn:
  type: aws.recordset
  name: cdn.example.com
  type_value: A
  hosted_zone_id: ${zone}
  alias_target:
    dns_name: d123.cloudfront.net
    hosted_zone_id: Z2FDTNDATAQYW2  # CloudFront hosted zone ID
    evaluate_target_health: false
```

**Example (Weighted record for load balancing):**
```yaml
api_primary:
  type: aws.recordset
  name: api.example.com
  type_value: A
  hosted_zone_id: ${zone}
  set_identifier: primary
  weight: 100
  resource_records:
    - 192.0.2.10
  ttl: 60

api_secondary:
  type: aws.recordset
  name: api.example.com
  type_value: A
  hosted_zone_id: ${zone}
  set_identifier: secondary
  weight: 50
  resource_records:
    - 192.0.2.11
  ttl: 60
```

## Common Patterns

**Zone delegation:**
- Create a hosted zone for a subdomain (ns.example.com)
- Create NS records in the parent zone pointing to the subdomain's nameservers
- The parent zone's `nameservers` are available after creation

**Load balancing with weighted records:**
- Create multiple recordsets with the same name and type
- Assign `weight` values proportional to desired traffic split (e.g., 70 primary, 30 backup)
- Route 53 distributes queries based on relative weights

**Health-checked failover:**
- Create PRIMARY and SECONDARY records with unique `set_identifier` values
- Point both to health checks via `health_check_id`
- Route 53 fails over to SECONDARY only when PRIMARY fails its health check

**Alias records to AWS resources:**
- Use `alias_target` for CloudFront, ELB, S3, API Gateway, and other AWS services
- Aliases do not incur Route 53 query charges (like CNAME but for apex domains)
- Each AWS service has its own hosted zone ID (e.g., CloudFront is Z2FDTNDATAQYW2)

## Common Pitfalls

- **Missing trailing dot in zone name:** The `hosted_zone_name` attribute for creating records requires a trailing dot (e.g., example.com.), but `name` can omit it.
- **CNAME at apex:** You cannot create a CNAME for the apex domain (example.com). Use an alias record or an A record instead.
- **TTL too high:** A TTL over 86400 (1 day) is rarely useful and causes slow propagation of updates. Keep it at 300–3600 for most records.
- **Record name case sensitivity:** DNS is case-insensitive, but consistency in configuration helps readability.
- **CloudFront alias to www:** Aliases to CloudFront distributions for www subdomains require the distribution's domain, not the original origin domain.
- **Missing zone ID in records:** Every recordset requires either `hosted_zone_id` (ID of the zone) or `hosted_zone_name` (name with trailing dot); omitting both causes an error.

## References

- [aws.hostedzone reference](../reference/route53/hostedzone.md)
- [aws.recordset reference](../reference/route53/recordset.md)
- [Route 53 User Guide](https://docs.aws.amazon.com/route53/)
- [CloudFront hosted zone IDs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-returned.html)
