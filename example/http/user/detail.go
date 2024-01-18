package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type DetailRequest struct {
	Id uint `uri:"id"`
}

type DetailResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

func Detail(c *gin.Context, req *DetailRequest) (*DetailResponse, error) {
	return &DetailResponse{
		Id:       req.Id,
		Username: fmt.Sprintf("Username%d", req.Id),
	}, nil
}
