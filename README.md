# quickpress
Small tool to automate SSRF wordpress and XMLRPC finder

* DNS resolution it's not a SSRF :P 



* Usage from src
  * go run quickpress.go -target https://www.target.com -server evil.com
* Usage from releases (download at: https://github.com/t0gu/quickpress/releases/tag/1.0 )
  * ./quickpress -target https://www.target.com -server evil.com
  
* Example
  ![](https://github.com/t0gu/quickpress/blob/master/qf.gif)


# TODO
1. Pass wordlist with targets to test
2. Maybe brute force module.. 
