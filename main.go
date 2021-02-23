package main

import (
	"net/http"

	"mrigankapaul/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
