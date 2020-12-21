package registry

import (
	"Week04/api/registry"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
)

func main() {
	e := gin.New()

	registry.RegisterLoginHTTPServer(e)

	Go(func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		<-c
		if e != nil {

		}
	})

	if err := e.Run(":8000"); err != nil {
		fmt.Println(err)
	}
}

func Go(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		f()
	}()
}
