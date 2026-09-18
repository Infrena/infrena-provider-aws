# Contributing to infrena-provider-aws

Thanks for considering it. This is the AWS provider plugin for
[Infrena](https://github.com/infrena/infrena): one generic provider serving every resource type
AWS Cloud Control API supports, generated from AWS's published CloudFormation schemas. This file
covers how the repository works and why there is a CLA.

## Before you start

**Open an issue first for anything non-trivial.** Most of this codebase was argued before it was
written — `docs/specs/` and `docs/plans/` carry the reasoning, and `CLAUDE.md` carries the rules
that are easy to get wrong and expensive to get wrong. A pull request that cuts across one of
those is not a bad pull request, but it is a conversation, and it is cheaper to have before the
code exists.

Small fixes — a typo, a wrong comment, a genuinely broken thing — just send them.

## What this repository expects of code

These are not style preferences. They are the rules the codebase is built on, and a change that
breaks one will be sent back:

- **The catalog is generated. Never hand-edit `internal/catalog/catalog.json.gz`.** Change
  `gen/overlay.yaml` or the generator in `internal/gen`/`internal/cfn`, regenerate with
  `go run ./cmd/gen-cloudcontrol`, and commit the diff like any other reviewed change.
  `gen/names.lock.json` only grows: once a CloudFormation type is assigned a name, that name
  never moves to a different type.
- **A test must be able to fail.** Write it, then break the code it covers so it still compiles
  and behaves wrongly, and confirm the test fails for the right reason. A test that passes when
  you break the thing it guards is worse than no test, because it reports safety that is not
  there. This repository has found three of them, and every one was green for weeks: one that
  only covered half a condition, and one that had never compared anything at all. Do not add a
  fourth. Record the sabotage in the commit message.
- **stdout is the protocol.** One stray `fmt.Println` corrupts the plugin stream, and the symptom
  is an unrelated parse error much later. Log to stderr.
- **Never log a credential, a signed request, or a `Sensitive` value.** stderr reaches CI logs.
- **Nested values are reconciled** (`internal/ccprov/reconcile.go`). A change there is a change to
  whether plans converge, and the e2e suite is the check.
- **Dependencies are argued in an issue, not in a pull request.** Adding one AWS SDK service was
  measured: no new indirect modules, +503 KiB of binary, and a cold build going from 8.7 s to
  28.9 s. That cost is the thing to justify.
- **Document what you changed.** Work is not finished until the docs reflect it, and the generated
  reference pages are part of that — run `go run ./cmd/gen-docs` and commit the result.

## Running the tests

```bash
go test -count=1 ./...                  # unit, the fake Cloud Control, and the protocol tests
go vet -tags e2e,live ./...
gofmt -l .
```

`-count=1` is **mandatory**. Go caches test results, and a cached pass hides a fixture edit.

```bash
# what CI blocks on: builds against the exact infrena release go.mod requires
GOWORK=off GOPRIVATE='github.com/infrena/*' go test -count=1 ./...

# the end-to-end suite, against a real infrena binary
go test -tags e2e -count=1 ./e2e/
```

Two things about the e2e suite that have caught people out, including us:

- **`GOWORK=off` does not make it pinned.** It governs this module's own resolution, while the
  suite builds its infrena binary from `INFRENA_SRC`, defaulting to a sibling checkout at
  `../../infrena`. To exercise a real release, point `INFRENA_SRC` at a `git archive` of that tag.
- **A skip is not a pass.** If it cannot find a host it prints `E2E SKIPPED:` on stderr and
  reports `ok`. Check for that line. There is no `INFRENA_REQUIRE_PLUGIN` in this repository; that
  variable belongs to infrena, and setting it here does nothing.

**Do not run the live suite in a pull request.** `go test -tags live ./live/` creates real
resources in a real AWS account and costs real money. It never runs in CI, it needs
`INFRENA_AWS_LIVE_PROFILE` and `INFRENA_AWS_LIVE_ACCOUNT`, and it needs the maintainer's approval
each time. See `live/README.md`.

## Commit messages

Plain English, present tense, and explain **why** rather than what — the diff already says what.
No em-dashes. If a commit fixes something subtle, the message is the right place to record how it
was found.

**Stage explicit paths.** `git add -A`, `git add .` and `git commit -am` are not used here: they
sweep up scratch files and unrelated work, and a commit containing something its message does not
mention is a commit nobody can review.

```bash
git add path/one path/two
git commit -m "..." -- path/one path/two
```

## Why there is a CLA

Most projects this size use a lightweight sign-off instead. Infrena asks for a
[Contributor License Agreement](CLA.md), and it is worth being straight about the reason.

**What it is not:** it is not a copyright assignment. You keep ownership of everything you write.

**What it does, for you:** it makes your grant irrevocable, so your contribution cannot later be
pulled out from under people who have built on it.

**What it does, for the project:** there is a commercial platform planned around Infrena. The CLA
means the project has the rights it needs to build that without chasing every past contributor.

**The limit that makes this fair.** Clause 4 permits relicensing **only to other OSI-approved open
source licences**. The project cannot take your contribution — or the CLI — proprietary or
source-available. That restriction exists because it is precisely the mechanism by which Terraform
was moved to a non-open licence in 2023. You do not have to take anyone's word for it; the
agreement you sign does not permit it.

What is free and what is paid is written down in
[Infrena's open core boundary](https://github.com/infrena/infrena/blob/main/docs/open-core.md).

## Signing

Comment on your first pull request:

```
I have read the CLA document and I hereby sign the CLA.
```

Once, ever, across Infrena and its first-party plugins. Tell us in the pull request if you are
contributing on behalf of an employer.

## Writing your own provider

You need none of this to write a provider for Infrena. Plugins are separate programs in separate
repositories, talking to Infrena over a documented protocol, and they are yours — your repository,
your licence, your release schedule. This repository is a first-party plugin and is maintained
alongside Infrena itself, which is why it carries the same agreement.

If you are building one, `pkg/pluginsdk` in the Infrena repository is the whole of a plugin's
`main()`, and this repository is a worked example of a large one.
