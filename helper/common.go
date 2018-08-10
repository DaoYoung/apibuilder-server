package helper

import (
	"reflect"
	"log"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

func Contains(intSlice []string, searchInt string) bool {
	if len(intSlice) == 0 {
		return false
	}
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}

func SetForbidUpdateFields(fs ...string) []string {
	res := []string{"id", "created_at", "updated_at", "deleted_at"}
	for _, value := range fs {
		res = append(res, value)
	}
	return res
}
func clone(i interface{}) interface{} {
	// Wrap argument to reflect.Value, dereference it and return back as interface{}
	return reflect.Indirect(reflect.ValueOf(i)).Interface()
}

func clearObj(obj interface{}) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	log.Println(111, v, v.Kind(), v.NumField())
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		//if f.Kind() == reflect.Struct {
		if f.Type().Name() == "BaseFields" {
			clearObj((&f).Addr().Interface())
			continue
		}
		if f.CanSet() == false {
			continue
		}
		if f.Kind() == reflect.Int {
			f.SetInt(0)
		}
		if f.Kind() == reflect.String {
			f.SetString("")
		}
		if f.Kind() == reflect.Slice {
			f.Set(reflect.Value{})
		}
	}
}


type JsonSuccess struct {
	Data interface{} `json:"data"`
}

func ReturnSuccess(c *gin.Context, code int, data interface{}) {
	js := new(JsonSuccess)
	js.Data = data
	c.JSON(code, js)
}

func MapUrlQuery(query url.Values, obj interface{})  map[string]interface{}{
	condition := make(map[string]interface{})
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Struct {
			for j := 0; j < f.NumField(); j++ {
				nm := snakeString(t.FieldByIndex([]int{i, j}).Name)
				log.Println(nm)
				if p := query.Get(nm);p != ""{
					condition[nm] = p
				}
			}
			continue
		}
		s := t.Field(i)
		nm := snakeString(s.Name)
		log.Println(nm)
		if p := query.Get(nm);p != ""{
			condition[nm] = p
		}
	}
	log.Println(condition)
	return  condition
}

// 驼峰式写法转为下划线写法
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	num := len(s)
	if num<3 {
		return strings.ToLower(s)
	}
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z'  {
			data = append(data, '_')
		}

		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}
// 下划线写法转为驼峰写法
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}