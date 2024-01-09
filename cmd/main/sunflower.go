// @copyright Copyright 2024 Willard Lu
// @email willard.lu@outlook.com
// @language go 1.18.1
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package main

import (
	"github.com/gin-gonic/gin"
)

/*
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "belighT928s"
	dbname   = "youling"
)*/

func main() {
	/*
		var err error
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err = sql.Open("postgres", psqlconn)
		if err != nil {
			log.Fatal(err)
		}*/

	router := gin.Default()
	router.Static("/", "./static")
	router.Run("192.168.0.104:8080")
}
