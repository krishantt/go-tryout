package main

import (
    "context"
    "flag"
	"fmt"
	"github.com/krishantt/go-tryout/grpc/client"
)

var(
	addr = flag.String("addr", "127.0.0.1:80", "The address to connect to.")
	author = flag.String("author", "", "The author whoose quote to get")
)

func main() {
	flag.Parse()
    c, err := client.New(*addr)
    if err!= nil {
        panic(err)
    }
    author, quote, err := c.QOTD(context.Background(), *author)
    if err!= nil {
        panic(err)
    }
    fmt.Println(author, quote)
}