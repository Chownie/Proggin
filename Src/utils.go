// utils.go
package utils

import (
	"fmt"
	"web"
	"github.com/hoisie/mustache.go"
	"io/ioutil"
	"strings"
	"os"
)

type User struct {
	Id			int
	Name		string
	Password	string
	Salt		string
}

func ReportErr(s os.Error) {
	if s != nil {
		fmt.Println(s)
	}
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

func Loadmustache(filename string, args *map[string]string) string {
	file, _err := ioutil.ReadFile("Mst/" + filename)
	if _err != nil {
		fmt.Println(_err)
		return "File not found"
	}
	data := mustache.Render(string(file), args)
	return data
}
