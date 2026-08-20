package main

import (
	"github.com/voska/venancio"
	"github.com/voska/vtexkit/cli"
)

// version is injected at build time via -ldflags.
var version = "dev"

func main() {
	cli.Main(cli.App{
		Store:       venancio.Store,
		Version:     version,
		Description: "Drogaria Venancio pharmacy CLI for humans and AI agents.",
	})
}
