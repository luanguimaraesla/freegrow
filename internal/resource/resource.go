package resource

import (
	"reflect"

	"github.com/luanguimaraesla/freegrow/internal/global"
)

type Tags map[string]string

type Metadata struct {
	Name string `yaml:"name" json:"name"`
	Tags Tags   `yaml:"tags" json:"tags"`
}

type ResourceList struct {
	Resources []interface{} `json:"resources"`
}

func NewResourceList(i interface{}) *ResourceList {
	s := reflect.ValueOf(i)
	if s.Kind() != reflect.Slice {
		global.Logger.Fatal("non-slice type")
	}

	rs := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		rs[i] = s.Index(i).Interface()
	}

	return &ResourceList{
		Resources: rs,
	}
}
