package utils

const author string = " [+] Author: r00td3v\n"
const version string = " [+] Version: 1.2\n\n"

//UglyBanner banner
func UglyBanner() string {

	var banner string = `
	

██╗  ██╗███╗   ███╗██╗     ██████╗ ██████╗  ██████╗ ███████╗ ██████╗ █████╗ ███╗   ██╗
╚██╗██╔╝████╗ ████║██║     ██╔══██╗██╔══██╗██╔════╝ ██╔════╝██╔════╝██╔══██╗████╗  ██║
 ╚███╔╝ ██╔████╔██║██║     ██████╔╝██████╔╝██║█████╗███████╗██║     ███████║██╔██╗ ██║
 ██╔██╗ ██║╚██╔╝██║██║     ██╔══██╗██╔═══╝ ██║╚════╝╚════██║██║     ██╔══██║██║╚██╗██║
██╔╝ ██╗██║ ╚═╝ ██║███████╗██║  ██║██║     ╚██████╗ ███████║╚██████╗██║  ██║██║ ╚████║
╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝╚═╝  ╚═╝╚═╝      ╚═════╝ ╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝                                                                       	


`
	returnBanner := banner + author + version
	return returnBanner
}
