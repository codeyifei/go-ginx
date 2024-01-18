package user

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/codeyifei/go-ginx"
)

type ListRequest struct {
	Keywords string `form:"keywords" binding:"required"`
	ginx.PaginationRequest
}

type ListResponse ginx.PaginationWrapper[*ListItem]

type ListItem struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

func List(c *gin.Context, req *ListRequest) (*ListResponse, error) {
	slog.Info("UserList请求参数", "keywords", req.Keywords)

	return &ListResponse{
		List: []*ListItem{
			{Id: 1, Username: "Username1"},
			{Id: 2, Username: "Username2"},
		},
		Meta: req.ToMeta(2),
	}, nil
}
