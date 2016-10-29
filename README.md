# GoFury

composable, mvc library for `fasthttp`. Inspired [dropwizard] and [utron]

## Features

- Offer a set of reusable libraries similar to `gorilla` libraries for `net/http`
- Designed to be 12factor compliant from the ground up
- Both server and router are zero memory allocation and byte slice optimised 
- No lock in. Purely a glue library that is non intrusive. Freely compatible with any `fasthttp.RequestHandler` based code or Middleware

## Motivation
I wrote this library whilst looking for a go library that's:

- similar to Rails and [dropwizard][dropwizard] in intuitiveness
- pays proper respect to MVC pattern, especially "thin Controller with helper Services" rather than "throws everything into a fat handler".

The closest match I found was [utron][utron], which was impressive. But still suffers the same lock-in symptoms as more heavy weight web frameworks such as [iris][iris], [gin][gin] or [echo][echo]. Having an wrapper layer of `utron.Ctx`, `iris.Context`, `gin.Context` or `echo.Context` adds an extra layer of custom vendor complexity and straight away throws out any standard `net/http` or `fasthttp` compatible middlewares.

`gofury` on the other hand is a "library and not a "framework". It offers some opinions but does not enforce them. All libraries and interfaces can be replaced with alternatives to suit developer needs. Helpful but with zero lock-in.

## Friends
GoFury out of the box forms a composition glue between the following libraries:

- [fasthttp][fasthttp] zero mem alloc http server, sensible `RequestCtx` that is easy to test
- [fasthttprouter][fasthttprouter] zero mem alloc router
- [fusion][fusion] middleware chaining inspired by [alice][alice]
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

[dropwizard]:   https://github.com/dropwizard/dropwizard
[utron]:        https://github.com/gernest/utron
[iris]:         https://github.com/kataras/iris
[echo]:         https://github.com/labstack/echo
[gin]:          https://github.com/gin-gonic/gin

[fasthttp]:     https://github.com/valyala/fasthttp
[fasthttprouter]: https://github.com/buaazp/fasthttprouter
[fusion]:       https://github.com/gofury/fusion
[fastjsonapi]:  https://github.com/gofury/fastjsonapi
[envconfig]:    https://github.com/kelseyhightower/envconfig
[easyjson]:     https://github.com/mailru/easyjson
[log]:          https://github.com/apex/log     
[glide]:        https://github.com/Masterminds/glide
[alice]:        https://github.com/justinas/alice

[testify]:      https://github.com/stretchr/testify/assert
[httpexpect]:   https://github.com/gavv/httpexpect
[jsonbenchmark]:https://github.com/buger/jsonparser
