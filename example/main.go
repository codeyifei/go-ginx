package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/codeyifei/go-ginx"
)

type GetRequest struct {
	Source                 string `form:"source"` // 使用form tag，可以自动绑定及验证query参数
	Id                     uint   `uri:"id"`      // 使用uri tag，可以自动绑定及验证path参数
	ginx.PaginationRequest        // 合并分页请求，会自动绑定query中的page和page_size参数，并设置默认值，page = 1, page_size = 20
}

type GetResponseItem struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

type GetResponse ginx.PaginationWrapper[*GetResponseItem]

func GetHandleFunc(c *gin.Context, req *GetRequest) (*GetResponse, error) {
	slog.Info("Get Request", "source", req.Source, "id", req.Id, "page", req.Page, "pageSize", req.PageSize)

	return &GetResponse{
		List: []*GetResponseItem{
			{Id: 1, Username: "Username1"},
			{Id: 2, Username: "Username2"},
		},
		Meta: req.ToMeta(2),
	}, nil
}

type PostRequest struct {
	Authorization string `header:"Authorization"` // 使用header tag，可以自动绑定及验证header参数
	Username      string `json:"username"`        // 使用json tag，可以自动绑定及验证json参数
}

func PostHandleFunc(c *gin.Context, req *PostRequest) (*ginx.NilResponse, error) {
	slog.Info("Post Request", "authorization", req.Authorization, "username", req.Username)

	return ginx.NewNilResponse, nil
}

type HelloRequest struct {
	ginx.Meta     `method:"get" path:"/hello/:name"` // 通过ginx.Meta设置接口的method和path
	Authorization string                             `header:"Authorization"` // 使用header tag，可以自动绑定及验证header参数
	Name          string                             `uri:"name"`
	Greetings     string                             `form:"greetings" default:"Hello"` // 可以通过default tag设置默认值，用法参考https://github.com/creasty/defaults
}

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloHandleFunc(c *gin.Context, req *HelloRequest) (*HelloResponse, error) {
	slog.Info("Hello Request", "authorization", req.Authorization, "name", req.Name, "greetings", req.Greetings)

	return &HelloResponse{Message: fmt.Sprintf("%s, %s!", req.Greetings, req.Name)}, nil
}

// func init() {
// 	gin.SetMode(gin.ReleaseMode)
// }

func main() {
	r := gin.Default()
	ginx.Get(r, "/", GetHandleFunc)
	ginx.Post(r, "/", PostHandleFunc)
	ginx.Handle(r, HelloHandleFunc)
	if err := r.Run(); err != nil {
		panic(errors.Wrap(err, "启动服务失败"))
	}
}
