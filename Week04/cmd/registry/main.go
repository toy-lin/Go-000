package registry

import (
	"Week04/api/registry"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	e := gin.New()
	server := &http.Server{
		Addr:    ":8000",
		Handler: e,
	}

	registry.RegisterLoginHTTPServer(e)

	Go(func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		<-c

		if server != nil {
			err := server.Shutdown(context.TODO())
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	if err := server.ListenAndServe(); err != nil {
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
