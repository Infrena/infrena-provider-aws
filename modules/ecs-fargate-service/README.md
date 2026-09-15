# ECS Fargate Service Module

Creates an ECS cluster with a Fargate service running a containerized application.

## Resources

- ECS Cluster
- CloudWatch Logs log group
- IAM role for ECS task execution
- ECS Fargate task definition
- ECS service

## Inputs

| Input | Type | Default | Description |
|-------|------|---------|-------------|
| `name` | string | (required) | Name for the cluster and service |
| `image` | string | `nginx:latest` | Docker image URI |
| `cpu` | string | `256` | Task CPU units (256, 512, 1024, 2048, 4096) |
| `memory` | string | `512` | Task memory in MB (512-30720) |
| `subnet_ids` | list | (required) | List of subnet IDs for task placement |
| `security_group_ids` | list | (required) | List of security group IDs for tasks |
| `region` | string | `us-east-1` | AWS region for CloudWatch logs |

## Outputs

| Output | Description |
|--------|-------------|
| `cluster_name` | ECS cluster name |
| `service_name` | ECS service name |

## Usage

```yaml
modules:
  - ./modules/vpc
  - ./modules/ecs-fargate-service

resources:
  network:
    type: module.vpc
    name: demo

  web:
    type: module.ecs_fargate_service
    name: web-app
    image: nginx:latest
    cpu: "256"
    memory: "512"
    subnet_ids:
      - ${network.private_subnet_1_id}
      - ${network.private_subnet_2_id}
    security_group_ids:
      - sg-12345678
```

Note: This module creates a basic single-container task definition. For more complex configurations, define task definitions directly in your project.
