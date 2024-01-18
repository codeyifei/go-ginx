# go-ginx: gin框架的增强工具

> gin框架是目前非常流行的http框架，但是由于gin本身设计的比较简单，以至于我们在开发项目的时候需要自行开发辅助函数或者使用一个封装好的脚手架。鉴于此，ginx出现了。

> ginx本身并不是一个框架，它只是一个基于gin框架的增加工具集，ginx存在的目的只是帮助开发者用最小的修改来更快速的开发gin项目

## ginx解决的痛点
- [x] 对泛型的支持
- [x] 请求参数自动解析
- [x] 响应包装
- [ ] 自动生成接口文档
- [ ] 封装常用的中间件
- [ ] 支持SSE

## 代码示例

```go
package main

import (
    "log/slog"

    "github.com/gin-gonic/gin"
    "github.com/pkg/errors"

    "github.com/codeyifei/go-ginx"
)

type GetRequest struct {
    Source                 string `form:"source"` // 使用form tag，可以自动绑定及验证query参数
    Id                     uint   `uri:"id"`      // 使用uri tag，可以自动绑定及验证path参数
    ginx.PaginationRequest                        // 合并分页请求，会自动绑定query中的page和page_size参数，并设置默认值，page = 1, page_size = 20
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

func main() {
    r := gin.Default()
    ginx.Get(r, "/", GetHandleFunc)
    ginx.Post(r, "/", PostHandleFunc)
    if err := r.Run(); err != nil {
        panic(errors.Wrap(err, "启动服务失败"))
    }
}

```
