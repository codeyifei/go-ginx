package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/codeyifei/go-ginx/example/http"
)

func main() {
	r := gin.Default()
	http.RegisterRouter(r)
	if err := r.Run(); err != nil {
		panic(errors.Wrap(err, "启动服务失败"))
	}
}
