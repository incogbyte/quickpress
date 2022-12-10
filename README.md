# quickpress

Scan urls or a single URL against XMLRPC wordpress issues.

usage:

##### Install

```bash

$ go install github.com/incogbyte/quickpress@latest

```
Compiling by yourself

```bash
git clone https://github.com/incogbyte/quickpress.git
cd quickpress
go build -o quickpress
./quickpress
```

##### Usage

* List of URLS

```bash
# urls without / at the end of URL!
cat urls.txt | quickpress -server http://burpcollaborator.net
```

* Single URL
```bash
quickpress -target https://target.com -server http://burpcollaborator.net
```

