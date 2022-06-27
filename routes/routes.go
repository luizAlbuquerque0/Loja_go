package routes

import (
	"net/http"

	"github.com/luiz/loja/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
