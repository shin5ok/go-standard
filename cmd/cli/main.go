package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/shin5ok/go-standard/internal/service"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	p := service.New(name)

	ctx, cancel := context.WithCancel(context.Background())

	r, err := p.RunCore(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer cancel()

	result := make([]byte, 128)
	r.Read(result)

	fmt.Println(string(result))

}
