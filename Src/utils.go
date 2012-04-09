// utils.go
package main

import (
	"fmt"
	"github.com/chownplusx/mustache"
	"github.com/chownplusx/web"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

/* func MostRecent(limit int) []map[string]string {
	// FORMAT: map[string]string{"title": title, "content": example, "details": details}
}

func ByName() []map[string]string {
	// FORMAT: map[string]string{"title": title, "content": example, "details": details}

} */

func CreatePost(owner string, title string, content string) {
	date := time.Now().Format("2006-01-02 15:04")
	os.Mkdir(string(date), 0755)
	//ioutil.WriteFile(filename, data, perm)
}

func Loadmustache(filename string, args *map[string]string) string {
	file, _err := ioutil.ReadFile("Mst/" + filename)
	if _err != nil {
		fmt.Println(_err)
		return "File not found"
	}
	data := mustache.Render(string(file), args)
	return data
}

func Sendstatic(ctx *web.Context, val string) {
	file, _err := ioutil.ReadFile("inc/" + val)
	if _err != nil {
		return
	}
	filetype := strings.Split(val, ".")
	ctx.ContentType(filetype[len(filetype)-1])
	ctx.WriteString(string(file))
}
