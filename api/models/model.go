package models

import (
	"github.com/zknow/redisRankDemo/api/database"

	"github.com/go-redis/redis"
)

//AddMember 新增使用者
func AddMember(member string, score float64) error {
	newMember := redis.Z{Member: member, Score: score}
	return database.Conn.ZAdd("Rank", newMember).Err()
}

//RanksForTop100 取得前一百名資訊
func RanksForTop100() []redis.Z {
	return database.Conn.ZRevRangeWithScores("Rank", 0, 99).Val()
}

//GetUser 取得使用者排名/分數
func GetUser(user string) (int64, float64) {
	return database.Conn.ZRevRank("Rank", user).Val(), database.Conn.ZScore("Rank", user).Val()
}

//IsUserExist 檢查存在
func IsUserExist(user string) error {
	return database.Conn.ZRank("Rank", user).Err()
}

//RemoveUser 移除user
func RemoveUser(name string) bool {
	_, err := database.Conn.ZRem("Key", name).Result()
	return err != nil
}

//RemoveAll 移除所有資訊
func RemoveAll() bool {
	_, err := database.Conn.ZRemRangeByRank("Rank", 0, -1).Result()
	return err != nil
}
