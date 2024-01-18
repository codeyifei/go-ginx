package ginx

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	logFormatter = "[GINXdebug] func(*gin.Context, *%s) (*%s, error)\n"
)

type HandlerFunc[Req any, Resp any] func(c *gin.Context, req *Req) (resp *Resp, err error)

func adapter[Req any, Resp any](handlerFunc HandlerFunc[Req, Resp], types BindingTypes) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Req
		defaults.MustSet(&req)
		for _, t := range types {
			var err error
			switch t {
			case BindingTypeDefault:
				err = c.ShouldBind(&req)
			case BindingTypeQuery:
				err = c.ShouldBindQuery(&req)
			case BindingTypePath:
				err = c.ShouldBindUri(&req)
			case BindingTypeJson:
				err = c.ShouldBindJSON(&req)
			case BindingTypeHeader:
				err = c.ShouldBindHeader(&req)
			default:
			}
			if err != nil {
				c.Render(ResponseWrapper.Response(nil, errors.Wrap(err, "参数错误")))
				return
			}
		}

		resp, err := handlerFunc(c, &req)
		c.Render(ResponseWrapper.Response(resp, err))
	}
}

func Handle[Req any, Resp any](r gin.IRouter, handlerFunc HandlerFunc[Req, Resp]) {
	var req Req
	metaData := requestMeta(req)
	log[Req, Resp]()
	r.Handle(metaData.Method, metaData.Path, adapter(handlerFunc, metaData.BindingTypes))
}

func Get[Req any, Resp any](r gin.IRouter, path string, handlerFunc HandlerFunc[Req, Resp]) {
	handle(r, http.MethodGet, path, handlerFunc)
}

func Head[Req any, Resp any](r gin.IRouter, path string, handlerFunc HandlerFunc[Req, Resp]) {
	handle(r, http.MethodHead, path, handlerFunc)
}

func Post[Req any, Resp any](r gin.IRouter, path string, handlerFunc HandlerFunc[Req, Resp]) {
	handle(r, http.MethodPost, path, handlerFunc)
}

func Put[Req any, Resp any](r gin.IRouter, path string, handlerFunc HandlerFunc[Req, Resp]) {
	handle(r, http.MethodPut, path, handlerFunc)
}

func Patch[Req any, Resp any](r gin.IRouter, path string, handlerFunc HandlerFunc[Req, Resp]) {
	handle(r, http.MethodPatch, path, handlerFunc)
}

func Delete[Req any, Resp any](r gin.IRouter, path string, handlerFunc HandlerFunc[Req, Resp]) {
	handle(r, http.MethodDelete, path, handlerFunc)
}

func Options[Req any, Resp any](r gin.IRouter, path string, handlerFunc HandlerFunc[Req, Resp]) {
	handle(r, http.MethodOptions, path, handlerFunc)
}

func Group(r gin.IRouter, groupName string, fn func(gin.IRouter)) {
	fn(r.Group(groupName))
}

func handle[Req any, Resp any](r gin.IRouter, method, path string, handlerFunc HandlerFunc[Req, Resp]) {
	var req Req
	log[Req, Resp]()
	rt := reflect.TypeOf(req)
	r.Handle(method, path, adapter(handlerFunc, requestBindingTypesMeta(rt)))
}

func log[Req, Resp any]() {
	if gin.IsDebugging() {
		var (
			req  Req
			resp Resp
		)
		fmt.Printf(logFormatter, reflect.TypeOf(req), reflect.TypeOf(resp))
	}
}
