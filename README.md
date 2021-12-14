# Go walking skeleton
[![Go Report Card](https://goreportcard.com/badge/github.com/aurelius15/go-skeleton?style=flat-square)](https://goreportcard.com/report/github.com/aurelius15/go-skeleton)
[![Build Status](https://app.travis-ci.com/aurelius15/go-skeleton.svg?token=JGY9GZm9cXWUmfpm1c5C&branch=master)](https://app.travis-ci.com/aurelius15/go-skeleton)
[![codecov](https://codecov.io/gh/aurelius15/go-skeleton/branch/master/graph/badge.svg?token=M708NM3HOP)](https://codecov.io/gh/aurelius15/go-skeleton)

## Features
- logging via [zap-logging](go.uber.org/zap)
- configuration via [go-arg](github.com/alexflint/go-arg)
- REST API via [gin](github.com/gin-gonic/gin) with "access logging" and "correlation id" middlewares
- redis as a storage via [go-redis](github.com/go-redis/redis/v8)
- autoloading new routes and commands
- testing via [github.com/stretchr/testify](testify) and [apitest](github.com/steinfletcher/apitest) 
- profiling
- managing via makefile

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.