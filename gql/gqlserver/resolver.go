package gqlserver

import (
	"github.com/graphql-go/graphql"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) GetProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		// id, _ := p.Args["product_id"].(int) //dipaksa jadi int

		// update to use Usecase from previous session
		return dictionary.Product{}, nil
	}
}

// func (r *Resolver) GetAllProducts() graphql.FieldResolveFn {
// 	return func(p graphql.ResolveParams) (interface{}, error) {
// 		products, _ = service.GetAllProducts()
// 		return products, nil
// 	}
// }

// func (r *Resolver) AddProduct() graphql.FieldResolveFn {
// 	return func(p graphql.ResolveParams) (interface{}, error) {
// 		// id, _ := p.Args["product_id"].(int) //dipaksa jadi int

// 		// update to use Usecase from previous session
// 		return dictionary.Product{}, nil
// 	}
// }
