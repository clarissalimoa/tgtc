package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	productResolver *Resolver
	Schema          graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithProductResolver(pr *Resolver) *SchemaWrapper {
	s.productResolver = pr

	return s
}

func (s *SchemaWrapper) Init() error {
	// add gql schema as needed
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "ProductGetter",
			Description: "All query related to getting product data",
			Fields: graphql.Fields{
				"ProductDetail": &graphql.Field{
					Type:        ProductType,
					Description: "Get product by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: s.productResolver.GetProduct(),
				},
				// "AllProductsList": &graphql.Field{
				// 	Type:        graphql.NewList(ProductType), //harus dibuat jadi list
				// 	Description: "Get all products",
				// 	Resolve:     s.productResolver.GetAllProducts(),
				// },
			},
		}),
		// uncomment this and add objects for mutation
		// Mutation: graphql.NewObject(graphql.ObjectConfig{
		// 	Name:        "ProductGetter",
		// 	Description: "All query related to getting product data",
		// 	Fields: graphql.Fields{
		// 		"AddProduct": &graphql.Field{
		// 			Type:        ProductType,
		// 			Description: "Add new Product",
		// 			Resolve:     s.productResolver.AddProduct(),
		// 		},
		// 	},
		// }),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
