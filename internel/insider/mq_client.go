package insider

var redisClient = &RedisClient{nil, "", "", 0}

func ConsumeOne(callback func([]byte) error) error {
	return nil
}
