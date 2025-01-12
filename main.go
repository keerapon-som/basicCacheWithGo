package main

import (
	"cachelearn/cache"
	"cachelearn/service"
	"cachelearn/utils"
	"fmt"
)

func main() {
	redisClient := utils.GetRedisClient()
	c := cache.NewRedisCache(redisClient)
	// c := cache.NewMemoryCache()
	// Set a struct value

	c.Set("people_one_naja", service.Userinfo{Username: "John Doe Za 0011", Email: "johnxxx.doe@example.com", Role: "admin"}, 30)
	c.Set("ThisIsPeopleTwoZa", service.Userinfo{Username: "John Doe", Email: "john.doe@example.com", Role: "admin"}, 30)

	// Get the struct value
	var peopleOne service.Userinfo
	var peopleTwo service.Userinfo
	c.Get("people_one_naja", &peopleOne)
	c.Get("ThisIsPeopleTwoZa", &peopleTwo)

	fmt.Println("people_one_naja:", peopleOne)
	fmt.Println("--------------------")
	fmt.Println("ThisIsPeopleTwoZa:", peopleTwo)
}
