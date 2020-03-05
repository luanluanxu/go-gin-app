package main

import "github.com/luanluanxu/go-gin-app/internal/my_app"

func main() {
	my_app.Run()
	select {}
}
