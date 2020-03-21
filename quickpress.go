package main

import (
	"flag"
	"fmt"
	"os"

	"./core"
	"./utils"
)

func main() {

	utils.Banner()
	target := flag.String("target", "", ">> target url (e.g: https://wordpress.com)")
	server := flag.String("server", "", ">> server to receive connections SSRF test")
	flag.Parse()

	if *target == "" {
		fmt.Printf(">> Usage: quickpress -target https://wordpress.com -server myserver.com\n")
		os.Exit(1)
	}

	if *server == "" {
		fmt.Printf("Server is required \n")
	}

	quickpressTarget := core.New(*target)
	quickpressTarget.XMLRPCTest()
	quickpressTarget.TestSSRF(*server)
	quickpressTarget.TestSSRFProxy(*server)
}
