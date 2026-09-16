module github.com/infrena/infrena-provider-aws

go 1.27.0

require (
	github.com/aws/aws-sdk-go-v2 v1.47.0
	github.com/aws/aws-sdk-go-v2/config v1.33.4
	github.com/aws/aws-sdk-go-v2/credentials v1.20.4
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.38.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.332.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.50.0
	github.com/aws/smithy-go v1.28.1
	// infrena is private: fetch it with GOPRIVATE=github.com/infrena/* and git credentials. Local work
	// builds against ../infrena through a gitignored go.work; CI builds this exact version (GOWORK=off).
	github.com/infrena/infrena v0.7.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.20.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.5.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.5.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.14.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.10.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.38.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.43.0 // indirect
)
