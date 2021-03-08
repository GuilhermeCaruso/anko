package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sou um retorno atualizado automagicamente. De versdade! U22AU \n"))
	})
	fmt.Println("Novo")
	log.Println("To vivo!")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("To morto!")

	}

}
