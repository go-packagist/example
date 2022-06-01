package main

import (
	"crypto/md5"
	"fmt"
	"github.com/go-packagist/example/bootstrap"
	"github.com/go-packagist/framework/support/facades"
)

func main() {
	bootstrap.App()

	fmt.Println(facades.MustMd5Hasher().MustMake("password"))
	fmt.Printf("%x", md5.Sum([]byte("password")))
	fmt.Println("\naa")
}
