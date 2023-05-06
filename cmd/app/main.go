package main

import "database/sql"

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/prodcuts")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
