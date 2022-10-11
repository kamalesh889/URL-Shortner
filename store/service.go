package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type StorageService struct {
	Client *redis.Client
}

var (
	Storeservice = &StorageService{}
	//ctx          = context.Background()
)

const Cacheduration = 6 * time.Hour

// Intialize connection
func IntializeConnection() *StorageService {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})

	conn, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Error in connecting to db")
		panic(err)
	}

	fmt.Println("Connected successfully to redis db ", conn)

	Storeservice.Client = client
	return Storeservice

}

// It will save the original and generated Url
func SaveUrl(generatedUrl, originalUrl string) {
	err := Storeservice.Client.Set(generatedUrl, originalUrl, Cacheduration).Err()
	if err != nil {
		fmt.Println("Error in saving Url ", err)
	}
}

// It will retrive the original url from generated short url
func GetUrl(generatedUrl string) (string, error) {
	result, err := Storeservice.Client.Get(generatedUrl).Result()
	if err != nil {
		fmt.Printf("error in retriving original url %v for generatedUrl %s", err, generatedUrl)
		return "", err
	}

	return result, nil
}
