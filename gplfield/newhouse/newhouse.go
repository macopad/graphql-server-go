package newhouse

import (
	"encoding/json"
	//"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
)

// testUserList 固定测试用户数据
var testHouseList = []map[string]interface{}{
	{
		"id":       "guojinyong",
		"house1": "beike user 1 localtion beijing haidian district",
		"house2": "123456",
		"info":     "info-1",
	},
}

func getNewhouseData(cityID interface{}) map[string]interface{} {
	var url = "https://app.api.ke.com/newhouse/shellapp/resblock/detail?project_name=zjxfyybmmma&city_id=110000"
	data, _ := http.Get(url)
	content, _ := ioutil.ReadAll(data.Body)
	data.Body.Close()
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(content) , &result)
	if err != nil {
		log.Println("json uncode error:", err.Error())
	}
	return result
}