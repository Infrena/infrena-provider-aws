# Amazon CloudWatch Logs

This guide covers creating and managing CloudWatch Logs log groups and metric filters for application and infrastructure monitoring.

## Log Group

A log group is a collection of log streams that share the same configuration, retention policy, and access control.

**Resource type:** `aws.loggroup`

**Key attributes:**
- `log_group_name`: Name for the log group (e.g., /aws/lambda/my-function or /app/backend)
- `retention_in_days`: Days to retain log events (1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557, 2922, 3288, 3653; omit for indefinite retention)
- `kms_key_id`: ARN of KMS key for encryption at rest (optional)
- `deletion_protection_enabled`: true to require explicit disablement before deletion
- `bearer_token_authentication_enabled`: true to allow bearer token auth on log operations
- `tags`: Map of tag key-value pairs

**Example:**
```yaml
app_logs:
  type: aws.loggroup
  log_group_name: /app/backend
  retention_in_days: 30
  tags:
    Application: backend
    Environment: dev
```

**Example (with encryption):**
```yaml
sensitive_logs:
  type: aws.loggroup
  log_group_name: /app/audit
  retention_in_days: 365
  kms_key_id: ${var.kms_key_arn}
  deletion_protection_enabled: true
  tags:
    Sensitivity: confidential
```

## Metric Filter

A metric filter extracts data from log events and transforms it into CloudWatch metrics. Use filters to count errors, measure latency, or trigger alarms on application events.

**Resource type:** `aws.metricfilter`

**Key attributes:**
- `log_group_name`: Name of the log group to filter (must exist; recommended to reference it)
- `filter_name`: Name of the metric filter
- `filter_pattern`: CloudWatch Logs filter syntax to match log events (e.g., "[ERROR]" or "[timestamp, request_id, status != 200, ...]")
- `metric_transformations`: List of transformations; each is a map with:
  - `metric_name`: CloudWatch metric name
  - `metric_namespace`: Namespace for the metric (e.g., ApplicationMetrics)
  - `metric_value`: Value to publish (e.g., 1 for count, or a field extracted from the log)
  - `default_value`: (optional) Value if no match

Filter patterns use CloudWatch Logs syntax. Common examples:
- `"[ERROR]"` — matches any line containing ERROR
- `"[msg, ...]"` — matches JSON with a msg field
- `"[*, status_code != 200, ...]"` — matches fields where status_code is not 200
- `"{ ($.statusCode = 400) || ($.statusCode = 401) }"` — matches JSON with specific status codes

**Example (count errors):**
```yaml
error_count:
  type: aws.metricfilter
  log_group_name: ${app_logs.log_group_name}
  filter_name: ErrorCount
  filter_pattern: "[ERROR]"
  metric_transformations:
    - metric_name: ErrorCount
      metric_namespace: ApplicationMetrics
      metric_value: "1"
```

**Example (count by HTTP status):**
```yaml
http_4xx:
  type: aws.metricfilter
  log_group_name: ${app_logs.log_group_name}
  filter_name: HTTP4xx
  filter_pattern: "{ ($.statusCode >= 400) && ($.statusCode < 500) }"
  metric_transformations:
    - metric_name: HTTP4xxCount
      metric_namespace: ApplicationMetrics
      metric_value: "1"

http_5xx:
  type: aws.metricfilter
  log_group_name: ${app_logs.log_group_name}
  filter_name: HTTP5xx
  filter_pattern: "{ ($.statusCode >= 500) && ($.statusCode < 600) }"
  metric_transformations:
    - metric_name: HTTP5xxCount
      metric_namespace: ApplicationMetrics
      metric_value: "1"
```

**Example (extract and average latency):**
```yaml
latency:
  type: aws.metricfilter
  log_group_name: ${app_logs.log_group_name}
  filter_name: RequestLatency
  filter_pattern: "[timestamp, request_id, latency_ms, ...]"
  metric_transformations:
    - metric_name: RequestLatencyMs
      metric_namespace: ApplicationMetrics
      metric_value: "$latency_ms"
      default_value: "0"
```

## Common Patterns

**Application logging:**
- Create one log group per application or service (e.g., /app/api, /app/worker)
- Use `retention_in_days` based on compliance and cost: 7 days for dev, 30 for staging, 90+ for production
- Add tags for cost allocation and service ownership

**Error monitoring:**
- Create a metric filter that matches ERROR or FATAL log levels
- Set `metric_value: "1"` to count errors
- Reference the metric in CloudWatch alarms to notify on error spikes

**Performance monitoring:**
- Use JSON-formatted logs and extract timing fields
- Create metric filters with `metric_value: "$field_name"` to publish the actual value
- Combine metrics in CloudWatch dashboards to correlate patterns

**Long-term archival:**
- Set `retention_in_days: 3653` (10 years) for logs requiring long-term retention
- Use `kms_key_id` for encryption at rest if logs contain sensitive data
- Export logs periodically to S3 for cheaper archival (outside infrena scope)

## Common Pitfalls

- **Retention day values:** AWS only accepts specific values (1, 3, 5, 7, 14, 30, etc.). Using other values causes an error; omit the attribute for indefinite retention.
- **Filter pattern syntax errors:** CloudWatch Logs filter syntax is strict. Test patterns in the CloudWatch console before committing.
- **Metric value must be a string:** Even numeric values like latency must be quoted in YAML (e.g., "$latency_ms" or "1").
- **Log group must exist:** The metric filter cannot be created if the log group does not exist. Always create the log group first or reference it from another resource.
- **Field extraction in JSON:** Use `$` prefix only when extracting fields from JSON (e.g., `$statusCode`). For space-delimited logs, use positional syntax (e.g., `$1`, `$2`).
- **Default value not applied on match:** If the filter matches but the field is missing or unparseable, the metric is not published unless a `default_value` is set.

## References

- [aws.loggroup reference](../reference/logs/loggroup.md)
- [aws.metricfilter reference](../reference/logs/metricfilter.md)
- [CloudWatch Logs User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/)
- [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html)
