package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// Adicionado o verbo na propria rota especifica sem necessidade realizar validações para saber o verbo da request
	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	//Pegar parametros direto da rota
	mux.HandleFunc("GET /test/{id}", func(w http.ResponseWriter, r *http.Request) {
		idTest := r.PathValue("id")
		w.Write([]byte("Hello World" + idTest))
	})

	http.ListenAndServe(":8080", mux)

	//// Comparando versão e checando se é valida
	//old := "go1.19"
	//new := "go1.22"
	//
	//for _, v := range []string{old, new} {
	//	if version.IsValid(v) {
	//		fmt.Printf("%s is valid\n", v)
	//	} else {
	//		fmt.Printf("%s is invalid\n", v)
	//
	//	}
	//}
	//
	//switch version.Compare(old, new) {
	//case 0:
	//	fmt.Printf("%s and %s are equal\n", old, new)
	//case 1:
	//	fmt.Printf("%s is greater than %s\n", old, new)
	//case -1:
	//	fmt.Printf("%s is less than %s\n", old, new)
	//
	//}
	//
	////Diferente da versão inferior agora o comprimido apenas com as posições restantes
	//sliceA := []uint{1, 5, 10}
	//sliceB := []uint{4, 20, 16}
	//
	//join := slices.Concat(sliceA, sliceB)
	//fmt.Printf("%v joined slice\n", join)
	//
	//join = slices.Delete(join, 0, 2)
	//fmt.Printf("%v deleted slice\n", join)

}
