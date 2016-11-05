# GoFury

composable, open minded and light weight mvc library. Inspired [dropwizard] and [utron]. Usese `fasthttp` by default but can also switch to standard `net/http` in the future.

## Features

- Offer a set of reusable libraries similar to excellent `gorilla` libraries.
- Designed to be 12factor compliant from the ground up
- Both server and router are zero memory allocation and byte slice optimised 
- No lock in and non intrusive. Freely compatible with any `fasthttp.RequestHandler` based code or Middleware
- Not opinionated about how you structure your business logic

## Motivation
For a while I looked for a light weight go mvc library that was intuitive to use like Rails and [dropwizard][dropwizard]. The closest match I found was [utron][utron], which was impressive. But still suffers the same lock-in symptoms as more heavy weight web frameworks such as [iris][iris], [gin][gin] or [echo][echo]. Having an wrapper layer of `utron.Ctx`, `iris.Context`, `gin.Context` or `echo.Context` adds an extra layer of custom vendor complexity and straight away throws out any standard `net/http` or `fasthttp` compatible middlewares.

I wrote `gofury` to be a "library and not a "framework". It offers some opinions but does not enforce them. All libraries and interfaces can be replaced with alternatives to suit developer needs. Helpful but with zero lock-in.

```
func main() {
	app := gofury.BaseApplication{}

	// load configuration
	cfg := &gofury.HTTPConfig{}
	app.LoadConfig(cfg)

	// register services
	app.RegisterServices(createFastHTTPService(cfg, &log.Logger{}), createQueueService())

	// start up application
	app.StartUp()

	// shudown cleanly when application exits
	defer app.ShutDown()
}
```


## Friends
GoFury out of the box forms a composition glue between the following libraries:

- [fasthttp][fasthttp] zero mem alloc http server, sensible `RequestCtx` that is easy to use and test
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

See `examples/example_test.go`

## Build 

    go test $(glide novendor)

## Todo

- db migration
- metrics and healthcheck
- views
- standard net/http module

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
