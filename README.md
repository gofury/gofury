# FastMVC

composable, mvc library for `fasthttp`. Inspired [dropwizard] and [utron]

## Why
- offer a set of reusable libraries similar to `gorilla` libraries for `net/http`
- `fastmvc` is not a framework and allows you to swap out any existing components with 
`fasthttp.RequestHandler` based middleware

## Features

- 12factor compliant
- Both server and router are zero memory allocation and byte slice optimised 
- Compatible with any `fasthttp.RequestHandler` or Middleware

## Friends

FastMVC is a composition glue between the following awesome libraries:
- [fasthttp][fasthttp] for zero http server
- [furyroad][furyroad] handler and middleware chaining
- [fusion][fusion] handler and middleware chaining
- [envconfig][envconfig] environment variable based config
- [apex/log][log] simple, handler based structured logging

Testing is made easier with:
- [testify][testify] assertions and mocks
- [httpexpect][httpexpect] http testing library with `fasthttp` support

Other options libraries:
- [fastjsonapi][fastjsonapi] jsonapi conversion using `fasthttp` and `easyjson`
- [easyjson][easyjson] [fastest][jsonbenchmark] non-reflection struct based json parsing

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

[testify]:      https://github.com/stretchr/testify/assert
[httpexpect]:   https://github.com/gavv/httpexpect
[jsonbenchmark]:https://github.com/buger/jsonparser