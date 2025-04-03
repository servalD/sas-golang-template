package main

import (
	"os"
	"fmt"
)

func main(){
	dburl, exists := os.LookupEnv("DB_URL")
	if !exists {
		fmt.Println("DB_URL not set")
		return
	}
	fmt.Println(dburl)
}
