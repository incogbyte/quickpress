package utils

import "fmt"

//Banner prints banner
func Banner() {

	banner := `

	author: t0gu
	twitter: @t0guu

	\ʕ◔ϖ◔ʔ/

	Usage: quickpress -target https://foo.bar -server evil.com
	`

	fmt.Println(banner)
}
