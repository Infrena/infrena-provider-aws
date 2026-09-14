module github.com/infrata/infrata-provider-aws

go 1.27.0

// infrata is private: fetch it with GOPRIVATE=github.com/infrata/* and git credentials. Local work
// builds against ../infrata through a gitignored go.work; CI builds this exact version (GOWORK=off).
require github.com/infrata/infrata v0.2.0
