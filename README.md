# FastMVC

composable, mvc library for `fasthttp`. Inspired [dropwizard] and [utron]

## Features

- Offer a set of reusable libraries similar to `gorilla` libraries for `net/http`
- Designed to be 12factor compliant from the ground up
- Both server and router are zero memory allocation and byte slice optimised 
- Non intrusive and freely compatible with any `fasthttp.RequestHandler` based code or Middleware

## Friends

FastMVC is a composition glue between the following awesome libraries:
- [fasthttp][fasthttp] for zero mem alloc http server
- [furyroad][furyroad] zero mem alloc router
- [fusion][fusion] middleware chaining
- [envconfig][envconfig] environment variable based config
- [apex/log][log] simple, handler based structured logging
- [glide][glide] for depedency management that makes sense

Testing is made easier with:
- [testify][testify] assertions and mocks
- [httpexpect][httpexpect] http testing library with `fasthttp` support

Other options libraries:
- [fastjsonapi][fastjsonapi] jsonapi conversion using `fasthttp` and `easyjson`
- [easyjson][easyjson] [fastest][jsonbenchmark] non-reflection struct based json parsing

## Usage
To import the project, use either `glide` 

    glide get github.com/gofury/gofury
    
or `go get`

    go get github.com/gofury/gofury

Build the project

    go test $(glide novendor)

## Todo

- db migration
- metrics and healthcheck
- views

## Alternatives
Some other `fasthttp` based frameworks
- [iris][iris]
- [echo][echo]

[dropwizard]:   https://github.com/dropwizard/dropwizard
[utron]:        https://github.com/gernest/utron
[iris]:         https://github.com/kataras/iris
[echo]:         https://github.com/labstack/echo

[fasthttp]:     https://github.com/valyala/fasthttp
[furyroad]:     https://github.com/gofury/furyroad
[fusion]:       https://github.com/gofury/fusion
[fastjsonapi]:  https://github.com/gofury/fastjsonapi
[envconfig]:    https://github.com/kelseyhightower/envconfig
[easyjson]:     https://github.com/mailru/easyjson
[log]:          https://github.com/apex/log     
[glide]:        https://github.com/Masterminds/glide

[testify]:      https://github.com/stretchr/testify/assert
[httpexpect]:   https://github.com/gavv/httpexpect
[jsonbenchmark]:https://github.com/buger/jsonparser