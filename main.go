package main

import (
	"gin-gorm-oj/router"
)

func main() {
	r := router.Router()

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
