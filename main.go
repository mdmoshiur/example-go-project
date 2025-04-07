// service documentation
// ...
package main

import (
	"github.com/mdmoshiur/example-go/cmd"

	_ "net/http/pprof"

	_ "github.com/davecgh/go-spew/spew"
)

func main() {
	cmd.Execute()
}
