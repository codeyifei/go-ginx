package ginx

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type NilResponse struct{}

var NewNilResponse = &NilResponse{}

type IResponseWrapper interface {
	Response(resp any, err error) (code int, render render.Render)
}

var ResponseWrapper IResponseWrapper = &DefaultResponseWrapper{}

type DefaultResponseWrapper struct{}

func (w *DefaultResponseWrapper) Response(resp any, err error) (int, render.Render) {
	var r gin.H
	if err != nil {
		var code = -1
		var e *Error
		if errors.As(err, &e) {
			code = e.code
		}
		r = gin.H{
			"code":    code,
			"message": err.Error(),
		}
	} else {
		r = gin.H{
			"code":    0,
			"message": "Success!",
			"data":    resp,
		}
	}
	return http.StatusOK, render.JSON{Data: r}
}
