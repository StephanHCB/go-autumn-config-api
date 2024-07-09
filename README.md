# go-autumn-config-api

A library that handles configuration for enterprise microservices (API only package).

## About go-autumn-config-api

This package declares a few public types that you will need for your package to collaborate
with go-autumn-config implementations. 

No actual code or dependencies are added if you include this package.

Import only this package when working on a library.

## Available implementations

You have a number of choices that implement this API. Pick one for your application.

  * full yaml/env/command line configuration with support for multiple profiles: [go-autumn-config](https://github.com/StephanHCB/go-autumn-config)
  * containerized 15 factor microservice configuration (environment vars + local yaml file only): [go-autumn-config-env](https://github.com/StephanHCB/go-autumn-config-env)

Do **not** import any of these implementation packages when working on a library, or you'll prevent your
users from making their choice of configuration library.
