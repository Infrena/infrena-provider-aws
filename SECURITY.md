# Security policy

## Reporting a vulnerability

Report it privately through GitHub, using
[**Report a vulnerability**](https://github.com/infrena/infrena-provider-aws/security/advisories/new)
on this repository's Security tab. That opens a private advisory visible only to you and
the maintainers.

Please do not open a public issue for a security problem, and please do not disclose it
publicly until a fix is released.

Include the plugin version (`infrena plugins list`), the engine version (`infrena version`),
what you did and what happened. A minimal configuration that reproduces it is worth more
than a description.

If the problem is in the engine rather than this provider, report it against
[infrena/infrena](https://github.com/infrena/infrena/security/advisories/new). If you are
unsure, report it in either and it will be routed.

## Supported versions

Only the most recent release. This plugin is pre-1.0 and tracks the engine closely; check
`plugin.yaml` for the engine versions a release supports.

## What is in scope

This plugin holds AWS credentials and creates, changes and deletes real infrastructure.
In particular:

- **Credential handling.** Anything that causes credentials to be logged, written to disk,
  or used against an account other than the one the configuration names.
- **Secret handling.** Attributes the catalog marks sensitive must not appear in plans,
  state, diagnostics or reports.
- **Discovery and import.** Anything that causes a resource to be adopted, or a
  cloud-owned resource to be claimed, when it should not be.

## What is known, and not a vulnerability

- **Values reach AWS in cleartext**, and are recorded in state in cleartext. That is the
  engine's documented position, not this plugin's choice.
- **A resource type the catalog does not cover cannot be managed.** That is a gap, and a
  normal issue rather than a security report.
