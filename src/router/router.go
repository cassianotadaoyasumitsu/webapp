package router

import "github.com/gorilla/mux"

// Retorna todas as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
