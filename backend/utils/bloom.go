package utils

import (
	"github.com/bits-and-blooms/bloom/v3"
	"go.uber.org/zap"
	"log"
	"sparks/dal/mysql"
	"sparks/model"
)

var userBloomFilter *bloom.BloomFilter

func InitUserBloomFilter() {
	userBloomFilter = bloom.NewWithEstimates(100000, 0.01) // 假设预期元素数量为 100000，误判率为 0.01
}

func AddToUserBloom(data string) {
	userBloomFilter.Add([]byte(data))
}

func TestUserBloom(data string) bool {
	return userBloomFilter.Test([]byte(data))
}

func LoadUsernamesToBloomFilter() {
	var usernames []string
	err := mysql.DB.Model(&model.User{}).Pluck("username", &usernames).Error
	if err != nil {
		log.Fatal("Failed to retrieve usernames from database:", err)
	}

	for _, username := range usernames {
		AddToUserBloom(username)
	}

	zap.L().Info("Loaded %d usernames to the bloom filter.\n", zap.Int("size", len(usernames)))
}
