package newhouse

import "github.com/graphql-go/graphql"

// queryType 用户类型
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "newhouse",
	Description: "房源列表字段",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "ID值",
		},
		"info": &graphql.Field{
			Type:        graphql.String,
			Description: "数据",
		},
		"data": &graphql.Field{
			Type:        graphql.String,
			Description: "data",
		},
		"request_id": &graphql.Field{
			Type:        graphql.String,
			Description: "request_id",
		},
		"uniqid": &graphql.Field{
			Type:        graphql.String,
			Description: "uniqid",
		},
		"errno": &graphql.Field{
			Type:        graphql.String,
			Description: "uniqid",
		},
		"errer": &graphql.Field{
			Type:        graphql.String,
			Description: "errer",
		},
	},
})


