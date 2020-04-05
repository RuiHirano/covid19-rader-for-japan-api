module main

go 1.13

require (
	cloud.google.com/go/storage v1.6.0
	github.com/PuerkitoBio/goquery v1.5.1 // indirect
	github.com/aws/aws-lambda-go v1.15.0 // indirect
	github.com/aws/aws-sdk-go v1.29.32 // indirect
	github.com/carlescere/scheduler v0.0.0-20170109141437-ee74d2f83d82
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-gota/gota v0.10.1 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/valyala/fasttemplate v1.1.0 // indirect
	gonum.org/v1/gonum v0.7.0 // indirect
	handler v0.0.0-00010101000000-000000000000
)

replace handler => ./handler

replace types => ./types
