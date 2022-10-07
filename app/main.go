package main

import (
	"fmt"
	"goapi/infrastructure"
	"goapi/utils"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := infrastructure.DbInit()
	if err != nil {
		panic("DB Error")
	}

	err = godotenv.Load(".env")
	if err != nil {
		panic("☹ .envを読めませんでした")
	}
	port := ":" + os.Getenv("PORT")
	httpServeMux := http.NewServeMux()
	Router(httpServeMux)

	fmt.Println("👉Now Listening 8080")
	if err := http.ListenAndServe(port, utils.Log(httpServeMux)); err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}

}
