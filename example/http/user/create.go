package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/codeyifei/go-ginx"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Create(c *gin.Context, req *CreateRequest) (*ginx.NilResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "生成密码失败")
	}
	_ = hashedPassword
	return ginx.NewNilResponse, nil
}
