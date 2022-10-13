package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

// gorilla/sessions为自定义session后端提供cookie和文件系统session以及基础结构
// 主要功能是：
//	简单的API：将其用作设置签名（以及可选的加密）cookie的简便方法。
//	内置的后端可将session存储在cookie或文件系统中。
//	Flash消息：一直持续读取的session值。
//	切换session持久性（又称“记住我”）和设置其他属性的便捷方法。
//	旋转身份验证和加密密钥的机制。
//	每个请求有多个session，即使使用不同的后端也是如此。
//	自定义session后端的接口和基础结构：可以使用通用API检索并批量保存来自不同商店的session
// Cookie和Session的区别：
// 	Session是在服务端保存的一个数据结构，用来跟踪用户的状态，这个数据可以保存在集群、数据库、文件中
//  Cookie是客户端保存用户信息的一种机制，用来记录用户的一些信息，也是实现Session的一种方式

// 初始化一个cookie存储对象
// fyyxdp一个你自己的密匙，只要不被别人知道就行
var store = sessions.NewCookieStore([]byte("fyyxdp"))

func main() {
	r := gin.Default()
	r.GET("/save", SaveSession)
	r.GET("/get", GetSession)
	r.GET("/delete", DeleteSession)
	r.Run()
}

func SaveSession(c *gin.Context) {
	// 获取一个session对象
	session, err := store.Get(c.Request, "fyy")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	// 在session中存储值
	session.Values["fyy"] = "xdp"
	session.Values[1118] = 1114
	// 保存更改
	session.Save(c.Request, c.Writer)
}

func GetSession(c *gin.Context) {
	session, err := store.Get(c.Request, "fyy")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	fyy := session.Values["fyy"]
	xdp := session.Values[1118]
	fmt.Println(fyy, xdp)
}

func DeleteSession(c *gin.Context) {
	session, err := store.Get(c.Request, "fyy")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	// 将session的最大存储时间设置为小于零的数即为删除
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
}
