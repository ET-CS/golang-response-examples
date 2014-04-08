golang-response-examples
========================

Golang examples with several client\server responses.

## Minimal

Simple minimal Golang web-server

> go run minimal.go

Visit *http://127.0.0.1:8080/My-Name*

## Header

Set the header on Golang web-server

> go run header.go

curl -i localhost:8080

## Redirect

Redirect response

> go run redirect.go

Visit *http://127.0.0.1:8080/redirect"

## JSON

Response with simple JSON from struct

> go run json.go

Visit *http://127.0.0.1:8080"

## XML

Response with simple XML from struct

> go run xml.go

Visit *http://127.0.0.1:8080"

## File

Serve a file

> go run file.go

Visit *http://127.0.0.1:8080"

## Template

Simple HTML Response using template

> go run template.go

Visit *http://127.0.0.1:8080"

## Nested Templates

Nested Templates response

> go run nested.go

Visit *http://127.0.0.1:8080"

## Form

HTML Form POST and Response

> go run form.go

Visit *http://127.0.0.1:8080"

## AJAX

Simple AJAX call to submit and recieve byte[] data

> go run ajax.go

Visit *http://127.0.0.1:8080"

## AJAX - JSON

Simple AJAX JSON request and response

> go run ajax-json.go

Visit *http://127.0.0.1:8080"

## AJAX - OCTET

Simple AJAX JSON Octet response

> go run ajax-octet.go

Visit *http://127.0.0.1:8080"