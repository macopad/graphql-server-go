package newhouse

import (
	"code.macopad.com/graphql-server-go/gplparse"
	//"encoding/json"
	//"fmt"
	"github.com/graphql-go/graphql"
	"log"
)

type responseData struct {
	Errno       int    `json:"errno"`
	Error       string `json:"error"`
	Request_id  string `json:"request_id"`
	Data        string `json:"data"`
}


// List 房源列表查询
func List() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(queryType),
		Description: "房源列表查询",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: 1,
			},
			"query": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			log.Printf("Args %v \n", params.Args)
			log.Printf("VariableValues %v \n", params.Info.VariableValues)
			log.Printf("FieldASTs %v \n", gplparse.SelectionFieldNames(params.Info.FieldASTs))
			return getNewhouseData("110000"), nil
		},
	}
}

// ByID 根据ID用户查询
func ByID() *graphql.Field {
	return &graphql.Field{
		Type:        queryType,
		Description: "根据城市ID查询列表信息",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// 查找对应的id，SQL...
			ID := params.Args["id"]
			data := getNewhouseData(ID)
			return data, nil  //这里data返回必须是一个map
		},
	}
}
