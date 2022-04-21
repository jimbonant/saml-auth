package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/itzg/go-flagsfiller"
	"github.com/itzg/saml-auth-proxy/server"
	"log"
	"os"
)

var (
	version = "dev"
	commit  = "HEAD"
)

func main() {
	var serverConfig server.Config

	filler := flagsfiller.New(flagsfiller.WithEnv("SamlProxy"))
	err := filler.Fill(flag.CommandLine, &serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()

	if serverConfig.Version {
		fmt.Printf("%s %s (%s)\n", os.Args[0], version, commit)
		os.Exit(0)
	}

	checkRequired(serverConfig.BaseUrl, "base-url")
	checkRequired(serverConfig.BackendUrl, "backend-url")
	checkRequired(serverConfig.IdpMetadataUrl, "idp-metadata-url")

	ctx := context.Background()
	// server only returns when there's an error
	log.Fatal(server.Start(ctx, &serverConfig))
}

func checkRequired(value string, name string) {
	if value == "" {
		_, _ = fmt.Fprintf(os.Stderr, "%s is required\n", name)
		flag.Usage()
		os.Exit(2)
	}
}
