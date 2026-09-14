// Command infrata-plugin-aws is infrata's AWS provider. Infrata runs it; you do not.
package main

import (
	"github.com/infrata/infrata-provider-aws/internal/awsprov"
	"github.com/infrata/infrata/pkg/pluginsdk"
)

func main() { pluginsdk.Main(awsprov.NewPlugin()) }
