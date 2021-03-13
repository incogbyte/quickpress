# quickpress

Scan urls or a single URL against XMLRPC wordpress issues.

usage:

##### Install

Compiling by yourself

```bash
git clone https://github.com/v4lak/quickpress.git
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


#Todo
--> Tracking where ssrf request come from, when vulnerable.
