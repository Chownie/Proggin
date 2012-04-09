// BlogGo project main.go
package main

import (
	"github.com/chownplusx/web"
)

func main() {
	print("V0.34\n")
	web.Get("/()", IndexLoad)
	//web.Get("/(login)", login.LoadGet)
	//web.Post("/(login)", login.LoadPost)
	web.Get("/admin/(.*)", AdminLoadGet)
	web.Post("/admin/newpost(.*)", AdminNewPost)
	web.Get("/inc/(.*)", Sendstatic)
	web.Run("0.0.0.0:9999")
}
