package templates

import "html/template"

var IndexTempl = template.Must(template.ParseFiles("index.html"))
