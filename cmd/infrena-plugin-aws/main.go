// Command infrena-plugin-aws is infrena's AWS provider. Infrena runs it; you do not.
package main

import (
	"github.com/infrena/infrena-provider-aws/internal/awsprov"
	"github.com/infrena/infrena/pkg/pluginsdk"
)

func main() { pluginsdk.Main(awsprov.NewPlugin()) }
