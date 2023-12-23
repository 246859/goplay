# goplay
cmd tool to interact with go playground server, see https://github.com/golang/playground to learn more about playground server.

use server `play.golang.org` by default if no address provided. 

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
see https://github.com/246859/goplay to learn more about goplay

Usage:
  goplay [flags]
  goplay [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  fmt         fmt code snippet
  health      check whether the playground server is healthy
  help        Help about any command
  run         compile and run code snippet in playground
  share       share your code to go playground
  version     get go version of the playground server
  view        view the specified code snippet

Flags:
  -d, --address string     specified the go playground address (default "https://play.golang.org")
  -h, --help               help for goplay
  -p, --proxy string       proxy address
  -t, --timeout duration   http request timeout (default 20s)
  -v, --version            show goplay version

Use "goplay [command] --help" for more information about a command.
```

health check
```bash
$ goplay health
ok
```

specify target server
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
$ goplay share main.go
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
$ goplay run main.go
#1
hello goplay!
```

you can compile with snippet id
```bash
$ goplay run -i T9_8fv9CyRh
#1
hello goplay!
```

and work with pipeline
```bash
$ cat main.go | goplay run
#1
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
output
```
package main

import "fmt"

func main() {
        fmt.Println("hello goplay!")
}
```