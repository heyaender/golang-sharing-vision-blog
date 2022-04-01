package main

import (
	"sv-article/databases"
	"sv-article/routes"
)

func main() {
	databases.MySQLConnect()
	routes.LaunchApp()
}
