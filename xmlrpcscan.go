package main

import (
	"flag"
	"fmt"
	"os"

	"./core"
	"./utils"
)

func main() {
	banner := utils.UglyBanner()
	fmt.Println(banner)
	target := flag.String("target", "", "[-] (e.g: https://wordpress.site.com)")
	server := flag.String("server", "", "[-] (e.g: http://159.89.121.20 or http://mydomain.com")
	flag.Parse()

	serverUser := *server
	targetUser := *target

	if serverUser == "" {
		fmt.Println("[+] Server is required to test ssrf..")
		fmt.Println("./xmlrpcscan -server http://burpcollaborator.net")
		os.Exit(1)
	}

	targetStruct := core.New(targetUser, serverUser)

	//verify if data is from stdin
	file, _ := os.Stdin.Stat()
	if (file.Mode() & os.ModeCharDevice) == 0 {
		targetStruct.FromStdin()
	} else {
		if targetStruct.IsAlive(targetUser) {
			targetStruct.VerifyMethods(targetUser)
		}
	}

}
