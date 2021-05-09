package callbacks

import (
	"reflect"
)

var typeRegistry = make(map[string]reflect.Type)

func Register(mRecord interface{})  {
	constTypes := []interface{}{mRecord}
	for _, v := range constTypes {
		typeRegistry[reflect.TypeOf(v).Name()] = reflect.TypeOf(v)
	}
}

// SetInstance("model.WfTranstion").(model.WfTranstion).UnassignmentCallback
func SetInstance(name string) interface{} {
	v := reflect.New(typeRegistry[name]).Elem()
	return v.Interface()
}
