# goplay
cmd tool to interact with go playground server, see https://github.com/golang/playground to learn more about playground server.

default go playground server is play.golang.org if no address provided. 

## install
install cmd tool
```bash
$ go install github.com/246859/goplay/cmd/goplay@latest
```
you can also import goplay to use the client directly
```bash
$ go get -u github.com/246859/goplay@latest
```

## usage

help
```bash
$ goplay -h
cmd tool to interact with go playground server,
see https://github.com/golang/playground to learn more about playground server.

Usage:
  goplay [commands] [flags]
  goplay [command]

Available Commands:
  compile     compile and run code snippet
  completion  Generate the autocompletion script for the specified shell
  fmt         fmt code snippet
  health      check whether the playground server is healthy
  help        Help about any command
  share       share your code to go playground
  version     get go version of the playground server
  view        view the specified code snippet

Flags:
  -d, --address string     specified the go playground address (default "https://play.golang.org")
  -h, --help               help for goplay
  -p, --proxy string       proxy address
  -t, --timeout duration   http request timeout (default 20s)
  -v, --version            show version

Use "goplay [command] --help" for more information about a command.
```

health check
```bash
$ goplay health
ok
```

specify target
```bash
$ goplay -d https://play.golang.org health
ok
```

check version
```bash
$ goplay version
name: Go 1.21
release: go1.21
version: go1.21.4
```

share your local code to playground
```bash
$ goplay share -f main.go
T9_8fv9CyRh
```

view specified code snippet in playground
```bash
$ goplay view T9_8fv9CyRh
package main

import "fmt"

func main() {
        fmt.Println("hello goplay!")
}
```

compile your local code in playground
```bash
$ goplay compile -f main.go
hello goplay!
```

you can compile with snippet id
```bash
$ goplay compile -i T9_8fv9CyRh
hello goplay!
```

use client directly
```go
package main

import (
	"github.com/246859/goplay"
	"fmt"
)

func main() {
	client, err := goplay.NewClient(goplay.Options{
		Address: Address,
		Proxy:   Proxy,
		Timeout: Timeout,
	})
	if err != nil {
		panic(err)
	}
	bytes, err := client.View("T9_8fv9CyRh")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

```