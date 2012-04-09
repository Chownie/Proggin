// admin.go
package main

import (
	"fmt"
	"github.com/chownplusx/web"
)

func defaultpage(ctx *web.Context) string {
	sidebar := Loadmustache("admin.mustache", &map[string]string{})

	mapping := map[string]string{"css": "../inc/site.css", "title": "Proggin: Admin panel",
		"sidebar": sidebar,
		"content": "Welcome to the admin panel, use the control box on your right to control the site content",
		"script":  "<script type=\"text/javascript\" src=\"../inc/adminref.js\"></script>"}

	return Loadmustache("frame.mustache", &mapping)
}

func postnewpost() string {
	content := Loadmustache("newpost.mustache", &map[string]string{})
	return content
}

func getnewpost() string {
	content := Loadmustache("newpost.mustache", &map[string]string{})
	return content
}

func getnewuser() string {
	content := Loadmustache("newuser.mustache", &map[string]string{})
	return content
}

func AdminLoadGet(ctx *web.Context, val string) {
	reg := map[string]string{"newpost": getnewpost(),
		"newuser": getnewuser()}

	output := ""
	if len(ctx.Params) == 0 {
		output = defaultpage(ctx)
	} else {
		output = reg[ctx.Params["mode"]]
	}
	ctx.WriteString(output)
}

func AdminNewPost(ctx *web.Context, val string) {
	//reg := map[string]string{"newpost": getnewpost(),
	//"newuser": getnewuser()}
	fmt.Println("We're in!")
	fmt.Println(ctx.Params)
	ctx.Redirect(304, "/")
}
