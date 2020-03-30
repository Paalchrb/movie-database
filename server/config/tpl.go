package config

import "html/template"

//TPL initializes templates from template folder
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
