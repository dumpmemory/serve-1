# serve

[![codecov](https://codecov.io/gh/chyroc/serve/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/serve)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/serve "go report card")](https://goreportcard.com/report/github.com/chyroc/serve)
[![test status](https://github.com/chyroc/serve/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/serve/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/serve)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Fserve.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Fserve)

![](./header.png)

## Install

```shell
brew install chyroc/tap/serve
```

```shell
go instll github.com/chyroc/serve@latest
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/chyroc/serve"
)

func main() {
	res := go_project_template.Incr(1)
	fmt.Println(res) // output: 2
}
```
