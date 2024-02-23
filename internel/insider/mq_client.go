package insider

import (
	"encoding/json"

	"github.com/lovelydett/my_ieee/internel/components"
	"github.com/lovelydett/my_ieee/internel/utils"
)

var redisClient *components.RedisClient = nil

func initRedisClient() {
	if redisClient != nil {
		return
	}
	config := utils.GetDeployConfig("config/deploy_config.yaml")
	redisConfig := config["redis"].(map[string]interface{})
	address := redisConfig["address"].(string)
	password := redisConfig["password"].(string)
	db := int(redisConfig["db"].(int64))

	client, err := components.GetRedisClient(address, password, db)
	if err != nil {
		panic(err)
	}

	redisClient = client
}

func popOneFromMq(key string) string {
	// Key is "cmd", type is list, rpop one
	initRedisClient()
	value := redisClient.Client.RPop(key)

	if value.Err() != nil {
		panic(value.Err())
	}

	return value.String()
}

func RunMqClient(handler func(map[string]interface{})) {
	for {
		cmdJson := popOneFromMq("cmd")
		// decode JSON
		jsonMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(cmdJson), &jsonMap)
		if err != nil {
			panic(err)
		}
		// do the actual work
		handler(jsonMap)
	}
}
