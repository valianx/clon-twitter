package main

import (
	"github.com/valianx/clon-twitter/bd"
	"github.com/valianx/clon-twitter/handlers"
	"log"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin coneccion a la base de datos")
		return
	}

	handlers.Manejadores()
}
