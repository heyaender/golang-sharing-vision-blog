package main

import (
	"sv-article/pkg/databases"
	"sv-article/routes"
)

func main() {
	databases.MySQLConnect()
	routes.LaunchApp()
}
