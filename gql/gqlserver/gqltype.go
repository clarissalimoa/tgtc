package gqlserver

import "github.com/graphql-go/graphql"

var ProductType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Product",
		Description: "Detail of the product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"imageUrl": &graphql.Field{
				Type: graphql.String,
			},
			"shopName": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
