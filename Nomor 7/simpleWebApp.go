package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(viper.GetString("server.port"), nil)
}
