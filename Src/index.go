// index.go
package main

func IndexLoad(val string) string {
	content := ""
	example := Loadmustache("content.mustache", &map[string]string{})
	details := "Posted by chown at 22:40GMT 14/01/2012"
	title := "This Will Be A Title"
	perpost := map[string]string{"title": title, "content": example, "details": details}
	for i := 0; i < 5; i++ {
		print(i)
		content += Loadmustache("perpost.mustache", &perpost)
	}
	mapping := map[string]string{"css": "inc/site.css", "title": "Proggin: Index", "sidebar": "bacon", "content": content}
	return Loadmustache("frame.mustache", &mapping)
}
