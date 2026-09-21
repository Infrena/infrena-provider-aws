# Amazon RDS

This guide covers creating and managing RDS databases with infrena, including DB instances, subnet groups, and parameter groups.

## DB Instance

Create and manage relational database instances in RDS. RDS handles backups, patching, and replication.

**Resource type:** `aws.rds.dbinstance`

**Key attributes:**
- `engine`: Database engine (mysql, postgres, oracle-ee, sqlserver-ex, mariadb, db2)
- `instance_class` (also: `db_instance_class`): Instance size (e.g., db.t3.micro, db.m5.large)
- `identifier` (also: `db_instance_identifier`): Name for the instance
- `username` (also: `master_username`): Master user login
- `password` (also: `master_user_password`): Master user password (sensitive; use a variable)
- `allocated_storage`: Storage size in GiB (string, e.g., "20")
- `storage_type`: gp2, gp3, io1, or io2
- `db_subnet_group_name`: Name of the subnet group (for VPC instances)
- `vpc_security_groups`: List of security group IDs to attach
- `backup_retention_period`: Days to retain automated backups
- `availability_zone`: Specific AZ (optional; AWS chooses if omitted)
- `publicly_accessible`: true to enable internet access
- `storage_encrypted`: true to encrypt storage with KMS
- `deletion_protection`: true to prevent accidental deletion
- `tags`: Map of tag key-value pairs

Passwords and other sensitive values must use variables, not literals. AWS marks them as sensitive in the reference.

**Example:**
```yaml
db:
  type: aws.rds.dbinstance
  engine: postgres
  instance_class: db.t3.micro
  identifier: app-db
  username: admin
  password: ${var.db_password}
  allocated_storage: "20"
  storage_type: gp3
  db_subnet_group_name: ${db_subnets.db_subnet_group_name}
  db_parameter_group_name: ${db_params.db_parameter_group_name}
  vpc_security_groups:
    - ${db_sg.group_id}
  backup_retention_period: 7
  deletion_protection: true
  tags:
    Environment: dev
```

## DB Subnet Group

Define subnets for RDS instances in a VPC. AWS requires at least two subnets in different availability zones.

**Resource type:** `aws.rds.dbsubnetgroup`

**Key attributes:**
- `description` (also: `db_subnet_group_description`): Human-readable name for the group
- `subnet_ids`: List of subnet IDs (required; must be at least 2 in different AZs)
- `tags`: Map of tag key-value pairs

**Example:**
```yaml
db_subnets:
  type: aws.rds.dbsubnetgroup
  description: Database subnets for RDS
  subnet_ids:
    - ${private_a}
    - ${private_b}
  tags:
    Environment: dev
```

## DB Parameter Group

Customize database engine parameters. Parameter groups apply to all instances using them; changing parameters can require a restart.

**Resource type:** `aws.rds.dbparametergroup`

**Key attributes:**
- `description`: Human-readable name (required on create)
- `family`: Engine family (e.g., postgres15, mysql8.0, oracle-ee-21)
- `parameters`: Map of parameter name to value (optional)
- `tags`: Map of tag key-value pairs

**Example:**
```yaml
db_params:
  type: aws.rds.dbparametergroup
  description: Custom parameters for Postgres
  family: postgres15
  parameters:
    log_statement: ddl
    shared_preload_libraries: pg_stat_statements
  tags:
    Environment: dev
```

## Common Patterns

**Secure database access:**
- Use `password: ${var.db_password}` or `${var.db_admin_password}` for all database passwords
- Store real passwords in infrena variables, environment-specific secrets, or AWS Secrets Manager
- Set `deletion_protection: true` for production databases
- Use `vpc_security_groups` to restrict access to application subnets only

**High availability:**
- Use `multi_az: true` for production to enable automatic failover
- Set `backup_retention_period` to at least 7 days for production
- Distribute DB subnet groups across at least 2 availability zones

**Cost optimization:**
- Start with `db.t3.micro` or `db.t4g.micro` for development
- Use `storage_type: gp3` for balanced cost and performance
- Lower `backup_retention_period` for non-critical databases

## Common Pitfalls

- **Passwords in config:** Passwords written as literals leak in commit history and logs. Always use variables.
- **Subnet group missing subnets:** DB subnet groups require at least 2 subnets; AWS returns an error if you provide only 1.
- **Engine version mismatches:** The parameter group family must match the instance engine. Use `family: postgres15` with `engine: postgres` version 15.x.
- **Changing replaces-on-change attributes:** `instance_class`, `storage_type`, and others trigger replacement (downtime). Plan before changing.
- **Encryption after creation:** `storage_encrypted` cannot be added to an unencrypted instance; you must recreate it.

## References

- [aws.rds.dbinstance reference](../reference/rds/rds-dbinstance.md)
- [aws.rds.dbsubnetgroup reference](../reference/rds/rds-dbsubnetgroup.md)
- [aws.rds.dbparametergroup reference](../reference/rds/rds-dbparametergroup.md)
- [RDS User Guide](https://docs.aws.amazon.com/rds/)
