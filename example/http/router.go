package http

import (
	"github.com/gin-gonic/gin"

	"github.com/codeyifei/go-ginx"
	"github.com/codeyifei/go-ginx/example/http/user"
)

func RegisterRouter(r gin.IRouter) {
	ginx.Group(r, "/users", func(ug gin.IRouter) {
		ginx.Get(ug, "", user.List)
		ginx.Get(ug, "/:id", user.Detail)
		ginx.Post(ug, "", user.Create)
	})
}
