package main

import (
	"database/sql"
	"fmt"
	_ "github.com/trinodb/trino-go-client/trino"
)

func main(){
	connectionString := "http://test:@localhost:8080?catalog=tpch"
	db, err := sql.Open("trino", connectionString)
	if err != nil {
		fmt.Println("Error connecting to db: ", err)
		return
	}
	result, err := db.Query("show schemas")
	if err != nil {
		fmt.Println("Error querying to db: ", err)
		return
	}
	for result.Next(){
		var schemaName string
		err = result.Scan(&schemaName)
		if err != nil {
			fmt.Println("Error scanning result: ", err)
			return
		}
		fmt.Println(schemaName)
	}
}