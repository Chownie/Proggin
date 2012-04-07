// login.go
package login

import (
	"crypto/sha256"
	"fmt"
	sqlite "gosqlite.googlecode.com/hg/sqlite"
	"hash"
	"strconv"
	"utils"
	"web"
)

func LoadGet(ctx *web.Context, val string) {
	ctx.WriteString(generic(ctx, ""))
}

func generic(ctx *web.Context, previous string) string {
	form := utils.Loadmustache("login.mustache", &map[string]string{})
	mapping := map[string]string{"css": "inc/site.css", "title": "Proggin: Login",
		"sidebar": "bacon", "content": previous + form}

	return utils.Loadmustache("frame.mustache", &mapping)
}

func LoadPost(ctx *web.Context, val string) {
	output := ""
	conn, _dberr := sqlite.Open("dbs/blog.db")
	utils.ReportErr(_dberr)

	s, _err := conn.Prepare("SELECT * FROM users WHERE name=?")
	utils.ReportErr(_err)

	password := ctx.Params["password"]
	username := ctx.Params["username"]

	s.Exec(username)
	if s.Next() {
		var user utils.User
		err := s.Scan(&user.Id, &user.Name, &user.Password, &user.Salt)
		utils.ReportErr(err)

		var h hash.Hash = sha256.New()
		h.Write([]byte(password + user.Salt))
		if string(h.Sum()) == user.Password {
			fmt.Println("Matching")
			ctx.Redirect(307, "/admin/")
			ctx.SetSecureCookie("id", strconv.Itoa(user.Id), 1*24*60*60)
		}
	} else {
		output = generic(ctx, "")
	}

	conn.Close()

	ctx.WriteString(output)
}
