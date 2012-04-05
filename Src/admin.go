// admin.go
package admin

import (
	"web"
	"utils"
)

func defaultpage(ctx *web.Context) string{
	sidebar := utils.Loadmustache("admin.mustache", &map[string]string{})

	mapping := map[string]string{"css": "../inc/site.css", "title": "Proggin: Admin panel",
	"sidebar":sidebar,
	"content": "Welcome to the admin panel, use the control box on your right to control the site content",
	"script":"<script type=\"text/javascript\" src=\"../inc/adminref.js\"></script>"}

	return utils.Loadmustache("frame.mustache", &mapping)
}

func getnewpost() string{
	content := utils.Loadmustache("newpost.mustache", &map[string]string{})
	return content
}

func getnewuser() string {
	content := utils.Loadmustache("newuser.mustache", &map[string]string{})
	return content

}

func LoadGet(ctx *web.Context, val string) {
	reg := map[string]string{"newpost":getnewpost(),
	"newuser":getnewuser()}

	cookie, _bcook := ctx.GetSecureCookie("id")
	print(cookie)
	output := ""
	if _bcook == false {
		{}
	}
	if len(ctx.Params) == 0 {
		output = defaultpage(ctx)
	} else {
		output = reg[ctx.Params["mode"]]
	}
	ctx.WriteString(output)
}
