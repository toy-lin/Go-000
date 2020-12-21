package registry

import (
    "Week04/api/registry"
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    e := gin.New()

    registry.RegisterLoginHTTPServer(e)

    if err := e.Run(":8000"); err != nil {
        fmt.Println(err)
    }
}
