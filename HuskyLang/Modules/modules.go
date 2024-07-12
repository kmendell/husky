package Modules

import (
	"fmt"
	"net/http"
)

func HuskyHTTPModule() {
	fmt.Println("Http Server Started on port 3333")
	fmt.Println("Ctrl+C to Quit Server...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}
