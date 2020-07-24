module main

go 1.13

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/carlescere/scheduler v0.0.0-20170109141437-ee74d2f83d82
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-gota/gota v0.10.1
	github.com/gocarina/gocsv v0.0.0-20200330101823-46266ca37bd3 // indirect
	github.com/jszwec/csvutil v1.3.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/valyala/fasttemplate v1.1.0 // indirect
	gonum.org/v1/gonum v0.7.0 // indirect
	handler v0.0.0-00010101000000-000000000000
	types v0.0.0-00010101000000-000000000000
)

replace handler => ./handler

replace types => ./types
