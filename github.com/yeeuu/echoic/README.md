# Echoic
[![Build Status](https://travis-ci.org/yeeuu/echoic.svg)](https://travis-ci.org/yeeuu/echoic)
[![Coverage Status](https://coveralls.io/repos/yeeuu/echoic/badge.svg?branch=master&service=github)](https://coveralls.io/github/yeeuu/echoic?branch=master)
[![License](https://img.shields.io/github/license/yeeuu/echoic.svg)](https://github.com/yeeuu/echoic/blob/master/LICENSE)

Echoic is a fast and unfancy micro web framework for Go. It based on [Echo](http://labstack.com/echo).

Echoic choose [Pongo2](https://github.com/flosch/pongo2) as offical support template engine (middleware) to work with [Django](https://www.djangoproject.com/). And it still has a zero alloc router. ;)

Echoic is using in Yeeuu Inc. production environment, we will fix mostly bugs we have meet.

**HTTP/2 support and Websocket support has removed.** HTTP/2 support will return when Golang 1.6 release.

## Installation

```sh
$ go get github.com/yeeuu/echoic
```

or

```sh
$ go get -u github.com/yeeuu/echoic
```

## Features

- Fast HTTP router which smartly prioritize routes.
- Extensible middleware, supports:
	- `echo.MiddlewareFunc`
	- `func(echo.HandlerFunc) echo.HandlerFunc`
	- `echo.HandlerFunc`
	- `func(*echo.Context) error`
	- `func(http.Handler) http.Handler`
	- `http.Handler`
	- `http.HandlerFunc`
	- `func(http.ResponseWriter, *http.Request)`
- Extensible handler, supports:
    - `echo.HandlerFunc`
    - `func(*echo.Context) error`
    - `http.Handler`
    - `http.HandlerFunc`
    - `func(http.ResponseWriter, *http.Request)`
- Sub-router/Groups
- Handy functions to send variety of HTTP response:
    - HTML
    - HTML via templates
    - String
    - JSON
    - JSONP
    - XML
    - File
    - NoContent
    - Redirect
    - Error
- Build-in support for:
	- Favicon
	- Index file
	- Static files
- Centralized HTTP error handling.
- Customizable HTTP request binding function.
- Customizable HTTP response rendering function, allowing you to use any HTML template engine.

##

## Performance

Echoic is based on Echo, mostly modifications is not related to Performance.

Echo benchmark:

> June 5, 2015    Echo: 38662 ns/op, 0 B/op, 0 allocs/op

## License
MIT License.
