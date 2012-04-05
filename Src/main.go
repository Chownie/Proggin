// BlogGo project main.go
package main

import (
	"web"

	//These are self-made files
	"utils"
	"login"
	"index"
	"admin"
	"newuser"
)

/* Passwording
original := input
var h hash.Hash = sha1.New()
salt := "HziNn!'*'()./A pretty and stable young pony"
h.Write([]byte(original+salt))
fmt.Printf("%s: %x\n", original, h.Sum())
*/

func main() {
	web.Get("/()", index.Load)
	web.Get("/(login)", login.LoadGet)
	web.Post("/(login)", login.LoadPost)
	web.Get("/admin/(.*)", admin.LoadGet)
	web.Post("/admin/newuser(.*)", newuser.LoadPost)
	web.Get("/inc/(.*)", utils.Sendstatic)
	web.Run("0.0.0.0:9999")
}
