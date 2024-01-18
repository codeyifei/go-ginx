package ginx

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"
)

const (
	metaKey = "Meta"
)

type Meta struct{}

type requestMetaData struct {
	Method       string
	Path         string
	BindingTypes BindingTypes
}

func requestMeta(obj any, metaKeys ...string) (data requestMetaData) {
	rt := reflect.TypeOf(obj)

	data.Method, data.Path = requestRouterMeta(rt, metaKeys...)

	data.BindingTypes = requestBindingTypesMeta(rt, metaKeys...)

	return
}

func requestRouterMeta(rt reflect.Type, metaKeys ...string) (method, path string) {
	if len(metaKeys) == 0 {
		metaKeys = []string{metaKey}
	}

	metaField, ok := rt.FieldByName(metaKeys[0])
	if !ok {
		panic(errors.New(fmt.Sprintf("结构体%s未设置Meta信息", rt)))
	}

	method, ok = metaField.Tag.Lookup("method")
	if ok {
		method = strings.ToUpper(method)
	} else {
		panic(errors.New(fmt.Sprintf("结构体%s未设置method tag", rt)))
	}

	path, ok = metaField.Tag.Lookup("path")
	if !ok {
		panic(errors.New(fmt.Sprintf("结构体%s未设置path tag", rt)))
	}

	return
}

func requestBindingTypesMeta(rt reflect.Type, metaKeys ...string) (types BindingTypes) {
	if len(metaKeys) == 0 {
		metaKeys = []string{metaKey}
	}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)

		if f.Name == metaKeys[0] {
			continue
		}

		if f.Type.Kind() == reflect.Struct {
			types = append(types, requestBindingTypesMeta(f.Type)...)
		}
		tag := f.Tag
		if _, ok := tag.Lookup("json"); ok {
			types = append(types, BindingTypeJson)
		}
		if _, ok := tag.Lookup("uri"); ok {
			types = append(types, BindingTypePath)
		}
		if _, ok := tag.Lookup("form"); ok {
			types = append(types, BindingTypeQuery)
		}
		if _, ok := tag.Lookup("header"); ok {
			types = append(types, BindingTypeHeader)
		}
	}
	types = lo.Uniq(types)
	sort.Sort(types)

	return
}
