package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// CarregaTemplates carrega todos os templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecutaTemplate executa um html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
