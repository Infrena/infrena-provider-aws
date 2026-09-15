# Amazon ECS (Elastic Container Service)

Amazon ECS is a fully managed container orchestration service that makes it easy to deploy, manage, and scale containerized applications. This guide covers three core ECS resources: clusters, task definitions, and services.

## ECS Cluster

An ECS cluster is a logical grouping of resources where you run tasks or services. Clusters are regional resources and can span multiple availability zones.

### Key attributes

- **name**: A unique name for the cluster (CloudFormation generates one if omitted)
- **capacity_providers**: List of capacity provider short names to associate with the cluster (optional)
- **default_capacity_provider_strategy**: Default capacity provider strategy when tasks are launched without explicit strategy (optional)
- **service_connect_defaults**: Default Service Connect namespace for new services (optional)
- **tags**: Metadata tags to categorize and organize the cluster

### Example

```yaml
ecs_cluster:
  type: aws.ecs.cluster
  name: production
  tags:
    Environment: prod
    Team: platform
```

## ECS Task Definition

A task definition describes how Docker containers should run within ECS. It specifies the container image, CPU, memory, environment variables, and logging configuration for your tasks.

### Key attributes

- **family**: Family name for the task definition (required for registration)
- **container_definitions**: List of container definitions in JSON format (required) specifying images, ports, volumes, and environment variables
- **cpu**: Number of CPU units for the task (required for Fargate)
- **memory**: Amount of memory in MiB (required for Fargate)
- **network_mode**: Docker networking mode (`bridge`, `host`, `awsvpc`, or `none`)
- **execution_role_arn**: ARN of the task execution role that grants ECS permissions to pull images and write logs (required for Fargate with ECR images)
- **task_role_arn**: ARN of the IAM role that containers assume to call AWS services
- **requires_compatibilities**: List of launch types this definition is compatible with (`EC2`, `FARGATE`, etc.)
- **tags**: Metadata tags for the task definition

### Example

```yaml
task_def:
  type: aws.ecs.taskdefinition
  family: app-api
  cpu: "256"
  memory: "512"
  network_mode: awsvpc
  requires_compatibilities: [FARGATE]
  execution_role_arn: ${ecs_task_execution_role.Arn}
  container_definitions:
    - name: api
      image: 123456789.dkr.ecr.us-east-1.amazonaws.com/app:latest
      portMappings:
        - containerPort: 8080
          hostPort: 8080
          protocol: tcp
      logConfiguration:
        logDriver: awslogs
        options:
          awslogs-group: /ecs/app
          awslogs-region: us-east-1
          awslogs-stream-prefix: ecs
```

## ECS Service

An ECS service runs and maintains a specified number of tasks. Services manage the lifecycle of tasks, handle load balancing, and update task definitions.

### Key attributes

- **service_name**: Unique name for the service within the cluster (optional, auto-generated if omitted)
- **cluster**: Cluster name or ARN where the service runs (defaults to `default` cluster if omitted)
- **task_definition**: Family and revision (`family:revision`) or full ARN of the task definition
- **desired_count**: Number of instantiations to run (optional)
- **launch_type**: Launch type for tasks (`EC2`, `FARGATE`, `EXTERNAL`)
- **platform_version**: Platform version for Fargate tasks (defaults to `LATEST`)
- **network_configuration**: VPC configuration including subnets and security groups (required for Fargate with `awsvpc` network mode)
- **load_balancers**: List of load balancers to associate (optional)
- **deployment_configuration**: Deployment parameters controlling task count during updates (optional)
- **scheduling_strategy**: Scheduling strategy (`REPLICA` or `DAEMON`)
- **enable_ecs_managed_tags**: Enable ECS managed tags for tasks (optional)
- **propagate_tags**: Propagate tags from task definition or cluster to tasks (optional)
- **tags**: Metadata tags for the service

### Example

```yaml
service:
  type: aws.ecs.service
  cluster: ${ecs_cluster.ClusterName}
  service_name: api-service
  task_definition: ${task_def.TaskDefinitionArn}
  desired_count: 2
  launch_type: FARGATE
  platform_version: "1.4.0"
  network_configuration:
    awsvpcConfiguration:
      subnets:
        - ${private_subnet_1}
        - ${private_subnet_2}
      securityGroups:
        - ${ecs_security_group}
      assignPublicIp: DISABLED
  deployment_configuration:
    maximumPercent: 200
    minimumHealthyPercent: 100
  tags:
    Environment: prod
```

## Common pitfalls

- **Missing execution role**: Fargate tasks pulling from ECR must have an execution role with permissions to pull images and write logs to CloudWatch. See the example for details.
- **Network configuration required for awsvpc**: Services using `awsvpc` network mode require `network_configuration` with VPC subnets and security groups.
- **Task definition not updated**: Changing the task definition doesn't automatically roll out new tasks. Use `force_new_deployment` to trigger a new deployment with the same task definition.
- **Desired count mismatch**: If `desired_count` isn't set, the service won't launch tasks. Explicitly set the count you want.
- **Container port mismatch**: Ensure port mappings in the container definition match what your application listens on.

## Reference pages

- [aws.ecs.cluster](../reference/ecs/ecs-cluster.md) — Cluster configuration
- [aws.ecs.taskdefinition](../reference/ecs/ecs-taskdefinition.md) — Task definition schema
- [aws.ecs.service](../reference/ecs/ecs-service.md) — Service configuration
