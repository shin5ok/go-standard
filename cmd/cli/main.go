package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/shin5ok/go-standard/internal/service"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	p := service.New(name)

	ctx, cancel := context.WithCancel(context.Background())
	r := p.RunCore(ctx)
	defer cancel()

	result := make([]byte, 128)
	r.Read(result)

	fmt.Println(string(result))

}
