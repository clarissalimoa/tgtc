package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/radityaqb/tgtc/backend/dictionary"
	"github.com/radityaqb/tgtc/backend/domain/product"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func AddProduct(w http.ResponseWriter, r *http.Request) {

	var p dictionary.Product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	p = product.AddProduct(context.Background(), p)
	fmt.Fprintf(w, fmt.Sprint("success, id product : ", p.ID))
}

// func GetProduct(w http.ResponseWriter, r *http.Request) {
// 	idstring := r.URL.Query().Get("id")

// 	idInt64, err := strconv.ParseInt(idstring, 10, 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	p, err := product.GetProduct(context.Background(), idInt64)
// 	if err != nil {
// 		// log.Fatal(err)
// 		fmt.Fprintf(w, err.Error())
// 	}

// 	val, err := json.Marshal(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Fprintf(w, string(val))
// }

//CONNECT TO DB
func GetProduct(w http.ResponseWriter, r *http.Request) {
	idstring := r.URL.Query().Get("id")

	idInt64, err := strconv.ParseInt(idstring, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	p, err := product.GetProduct(context.Background(), idInt64)
	if err != nil {
		// log.Fatal(err)
		fmt.Fprintf(w, err.Error())
	}

	val, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(val))
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var p dictionary.Product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	product.DeleteProduct(context.Background(), p.ID)
	fmt.Fprintf(w, "success")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var p dictionary.Product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	product.UpdateProduct(context.Background(), p)

	fmt.Fprintf(w, "success")
}
