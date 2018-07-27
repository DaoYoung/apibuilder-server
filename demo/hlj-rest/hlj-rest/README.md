# 婚礼纪GoLang Rest框架
一个基于[ant0ine/go-json-rest](https://github.com/ant0ine/go-json-rest)的Restful框架。让你简单快速的创建Restful接口。

## 特性
- 快速可扩展的路由（使用Trie数据结构）
- 丰富的中间件，如日志，Gzip，跨域，认证等
- 与Orm无关的资源，可简单快速地对资源进行增删改查。支持丰富的资源过滤，资源序列化功能

## 安装
将本项目放在`gopath/src`目录下

## 路由
支持Get, Post, Put, Delete路由，路由参数通过`:name`进行定义：如
```go
	route := rest.GetFunc("/user/:id", func(w rest.ResponseWriter, r *rest.Request) {
		userId := r.EnvParam("id").Int()
		...
	})
```
### 例子

curl:
``` sh
curl -i http://127.0.0.1:8080/hello
```

code:
```go
package main

import (
	"log"
	"net/http"
	"hlj-rest/rest"
)

func main() {
	route := rest.GetFunc("/hello", func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	})
	router, err := rest.MakeRouter(route)

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
```

## 请求与响应
### 请求
```go
func Test(w rest.ResponseWriter, r *rest.Request) {
    postId := r.PathParam("postId").Int() // 获取路由参数
    r.SetPathParam("postId", "1") // 设置路由参数
    
    id := r.QueryParam("page").Default("1").Int() // 获取GET参数page, 默认值为1
    r.SetQueryParam("perPage", "20") // 设置GET参数
    
    postId := r.PostParam("postId").Int() // 获取POST参数
    r.SetPostParam("postId", "1") // 设置Post参数
    
    postId := r.FormParam("postId").Int() // 获取Form参数
    r.SetFormParam("postId", "1") // 设置Form参数
    
    postId := r.EnvParam("postId").Int() // 获取Env参数（一般由中间件放入）
    r.EnvParam("postId", "1") // 设置Env参数
    
    // 反序列化POST参数
    var User struct {
    	Id int
    	Name string
    }
    var user User
    r.DecodeJsonPayload(&user) 
    // 设置反序列话的POST参数
    r.SetJsonPayload("Name", "guest")
}
```
### 响应
```go
func Test(w rest.ResponseWriter, r *rest.Request) {
    w.WriteHeader(200) // 写响应头
    w.WriteJson(user) // 返回Json数据
    w.Error(w, 500, "server error", 500) // 返回格式化错误
    w.WriteData(w, user, nil, 200) //返回格式化数据
    ...
}
```

## 中间件

| 名称 | 描述 |
|------|-------------|
| **AccessLogApache** | 基于Apache mod_log_config的访问日志中间件 |
| **AccessLogJson** | Json访问日志中间件 |
| **AuthBasic** | Basic HTTP认证 |
| **ContentTypeChecker** | 文档类型检查 |
| **Cors** | 跨域支持 |
| **Gzip** | Gzip压缩 |
| **If** | 条件中间件 |
| **Jsonp** | JSONP |
| **Recorder** | 统计 |
| **Status** | 请求状态 |
| **Timer** | 计时器 |


### 自定义中间件
curl demo:
```sh
curl -i http://127.0.0.1:8080/api/login
```

code:
```go
package main

import (
	"net/http"
	"log"
	"hlj-rest/rest"
	"hlj-rest/middleware"
)

type AuthMiddleware struct{}

func (auth *AuthMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		r.Env["USER"] = r.QueryParam("name").Default("anonymous").String()

		handler(w, r)
	}
}

func Login(w rest.ResponseWriter, r *rest.Request) {
	user := r.Env["USER"]
	name, _ := user.(string)

	rest.Ok(w, "Hi, " + name, nil)
}

func main() {
	auth := &AuthMiddleware{}
	routes := []*rest.Route{
		rest.Get("/login", rest.HandlerFunc(Login).Use(auth)),
	}

	router, _ := rest.MakeRouter(routes...)

	http.Handle("/api/", http.StripPrefix("/api", router.Use(middleware.DefaultDevStack...)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 资源
### 资源定义
任何类型的结构体都可以是一个资源
```go
// 用户资源
type User struct {
	Id   int
	Name string
}

// 文章资源
type Post struct {
	Id      int   
	Title   string
	Content string
	UserId  int
}

```
### 资源增删改查
当资源实现`resource.Resource`接口后，可通过以下方式对资源进行增删改查
```go
// 资源列表
func GetPosts(w rest.ResponseWriter, r *rest.Request) {
	handler := resource.NewHandler(&model.Post{}, w, r)
    handler.List(where)
}
func GetPost(w rest.ResponseWriter, r *rest.Request) {
	id := r.QueryParam("id").Int()
	handler := resource.NewHandler(&model.Post{}, w, r)
    handler.One(id, nil)
}
func AddPost(w rest.ResponseWriter, r *rest.Request) {
	handler := resource.NewHandler(&model.Post{}, w, r)
    handler.Add(nil)
}
func UpdatePost(w rest.ResponseWriter, r *rest.Request) {
	id := r.QueryParam("id").Int()
	handler := resource.NewHandler(&model.Post{}, w, r)
    handler.Update(nil)
}
func Delete(w rest.ResponseWriter, r *rest.Request) {
	id := r.QueryParam("id").Int()
	handler := resource.NewHandler(&model.Post{}, w, r)
    handler.Delete(nil)
}
```

### 资源序列化（即转map）
通过给资源定义以下函数，可以自定义resource.ToMap函数的返回结果
```go
// 文章关联的用户（当调用resource.ToMap函数，并且参数fields中包含user{}时会自动将结果放入ToMap的返回结果）
func (p Post) User() *User {
	return GetUserById(p.UserId) 
}

// 自定义ToMap返回的字段
func (Post) ToMap(data map[string]interface{}, fields []string) map[string]interface{} {
	// 可根据fields参数，修改resource.ToMap返回的数据
    return data
}
```

## 模版渲染
不支持模版渲染相关功能，若有SEO需要，[请使用node服务器进行渲染](https://www.zhihu.com/question/52235652)