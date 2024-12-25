package services

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, Config.MongodbDatabase, options.Client().ApplyURI(Config.MongodbUri))
	if err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")
}

var redisDefaultClient *redis.Client
var redisDefaultOnce sync.Once

var redisCache *cache.Cache
var redisCacheOnce sync.Once

func GetRedisDefaultClient() *redis.Client {
	redisDefaultOnce.Do(func() {
		redisDefaultClient = redis.NewClient(&redis.Options{
			Addr: Config.RedisDefaultAddr,
		})
	})

	return redisDefaultClient
}

func GetRedisCache() *cache.Cache {
	redisCacheOnce.Do(func() {
		redisCache = cache.New(&cache.Options{
			Redis:      GetRedisDefaultClient(),
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		})
	})

	return redisCache
}

func CheckRedisConnection() {
	redisClient := GetRedisDefaultClient()
	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to Redis!")
}

func Set(ctx context.Context, key string, value interface{}) error {
	err := redisCache.Set(&cache.Item{
		Ctx: ctx,
		Key:   key,
		Value: value,
	})
	return err
}

func Get(ctx context.Context, key string) (interface{}, error) {
	item := cache.Item{};
	err := redisCache.Get(ctx, key, &item);
	if err != nil {
		return nil, err
	}
	return item.Value, nil
}
