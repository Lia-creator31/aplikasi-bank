package main

import (
	"create-api/database"
	"create-api/router"
)

func main() {
	//print('dqw')
	//fmt.Print("Hello, ")
	database.StartDB()
	r := router.StartApp()
	r.Run(":8000")
}
