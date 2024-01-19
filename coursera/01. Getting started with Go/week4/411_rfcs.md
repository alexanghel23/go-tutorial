# RFC (Requests for comments)

They are definitions of Internet protocols and formats

Examples:
 - HTML (RFC 1866)
 - URI (RFC 3986)
 - HTTP (RFC 2616)

 ## Protocol packages

 Golang provides packages for important RFCs
 The packages provide functions to encode and decode protocol format. 

 Encoding means you take some kind of Go object and transform it into a protocol. Decoding is the opposite operation.

Examples:
- "net/http" - http.Get(www.uci.edu)
- "net" - net.Dial("tcp", "uci.edu:80")

## JSON

Data format, not protocol (RFC 7159)

Go Struct
p1 := Person(name: "joe", addr: "a st", phone: "123")

Equivalent Json object
{"name": "joe", "addr": "a st", "phone": "123"}