// newuser.go
package newuser

import (
	"crypto/sha256"
	sqlite "gosqlite.googlecode.com/hg/sqlite"
	"hash"
	"strconv"
	"time"
	"utils"
	"web"
)

func LoadPost(ctx *web.Context, val string) {
	username := ctx.Params["username"]
	password := ctx.Params["password"]

	conn, _dberr := sqlite.Open("dbs/blog.db")
	utils.ReportErr(_dberr)

	salt := strconv.Itoa64(time.Nanoseconds()) + username

	var h hash.Hash = sha256.New()
	h.Write([]byte(password + salt))

	s, _err := conn.Prepare("INSERT INTO users VALUES(NULL, ?, ?, ?)")
	utils.ReportErr(_err)

	s.Exec(username, string(h.Sum()), salt)
	s.Finalize()
	conn.Close()
	sidebar := utils.Loadmustache("admin.mustache", &map[string]string{})

	//TESTING, REMOVE LATER
	script := "<script type=\"text/javascript\" src=\"../inc/adminref.js\"></script>"
	content := "Welcome to the admin panel, use the control box on your right to control the site content"
	//ENDTESTING

	mapping := map[string]string{"css": "../inc/site.css",
		"title":   "Proggin: Admin panel",
		"sidebar": sidebar,
		"content": content,
		"script":  script}

	output := utils.Loadmustache("frame.mustache", &mapping)
	ctx.WriteString(output)
}
