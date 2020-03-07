# go-autumn-config-api

API for [go-autumn-config](https://github.com/StephanHCB/go-autumn-config).

## About go-autumn

A collection of libraries for [enterprise microservices](https://github.com/StephanHCB/go-mailer-service/blob/master/README.md) in golang that
- is heavily inspired by Spring Boot / Spring Cloud
- is very opinionated
- names modules by what they do
- unlike Spring Boot avoids certain types of auto-magical behaviour
- is not a library monolith, that is every part only depends on the api parts of the other components
  at most, and the api parts do not add any dependencies.  

Fall is my favourite season, so I'm calling it go-autumn.

## About go-autumn-config

A library that handles configuration for enterprise microservices.

See [go-autumn-config](https://github.com/StephanHCB/go-autumn-config) for details.

## About go-autumn-config-api

This package declares a few public types that you will need for your package to collaborate 
with go-autumn-config. No actual code or dependencies are added if you include this package.
