package service

import (
	"fmt"

	"github.com/radityaqb/tgtc/backend/database"
)

func SampleFunction() {
	fmt.Printf("My Service!")

	// // you can connect and
	// // get current database connection
	// db := database.GetDB()

	// // construct query
	// query := `
	// SELECT something FROM table_something WHERE id = $1
	// `
	// // actual query process
	// row = db.QueryRow(query, paramID)

	// // read query result, and assign to variable(s)
	// err = row.Scan(&ID, &name)
}

func GetAllProducts() {

	// // you can connect and
	// // get current database connection
	db := database.GetDB()

	// query rows from a table
	var (
		ename string
		sal   int
	)

	rows, err := db.Query("SELECT product_name, product_price, product_image FROM products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&ename, &sal)
		if err != nil {
			panic(err)
		}
		fmt.Println("\n", ename, sal)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// return rows.Scan()
}
