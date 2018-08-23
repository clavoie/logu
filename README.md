# logu [![GoDoc](https://godoc.org/github.com/clavoie/logu?status.svg)](http://godoc.org/github.com/clavoie/logu) [![Build Status](https://travis-ci.org/clavoie/logu.svg?branch=master)](https://travis-ci.org/clavoie/logu) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/logu)](https://goreportcard.com/report/github.com/clavoie/logu)

Injectable wrappers around Google AppEngine logging for Go. 

logu is an interface around the Google AppEngine [log package](https://cloud.google.com/appengine/docs/standard/go/logs/reference) that integrates well with dependency injection. An implementation of the logging interface is also provided for [testing](https://godoc.org/github.com/clavoie/logu#NewTestLogger), as well as a [null logger](https://godoc.org/github.com/clavoie/logu#NewNullLogger) implementation which throws away all log messages.

A full example of using logu along with [erru](https://github.com/clavoie/erru) and [di](https://github.com/clavoie/di) is available [here](https://godoc.org/github.com/clavoie/logu#example-Logger). 
