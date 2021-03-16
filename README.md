# quickpress

Scan urls or a single URL against XMLRPC wordpress issues.

usage:

##### Install

```bash
# dependencies
$ go get github.com/fatih/color

```
Compiling by yourself

```bash
git clone https://github.com/pownx/quickpress.git
cd quickpress
go build -o quickpress
./quickpress
```

##### Usage

* List of URLS
```bash
cat urls.txt | quickpress -server http://burpcollaborator.net
```

* Single URL
```bash
quickpress -target https://target.com -server http://burpcollaborator.net
```

