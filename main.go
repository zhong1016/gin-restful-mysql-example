package main

import (
	r "todolist/api"
)

func main() {

	r.Route().Run(":8989")
}
