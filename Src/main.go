// BlogGo project main.go
package main

import (
	"web"

	//These are self-made files
	"admin"
	"index"
	"login"
	"newuser"
	"utils"
)

func main() {
	web.Get("/()", index.Load)
	web.Get("/(login)", login.LoadGet)
	web.Post("/(login)", login.LoadPost)
	web.Get("/admin/(.*)", admin.LoadGet)
	web.Post("/admin/newuser(.*)", newuser.LoadPost)
	web.Get("/inc/(.*)", utils.Sendstatic)
	web.Run("0.0.0.0:9999")
}
