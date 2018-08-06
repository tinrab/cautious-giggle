package main

import (
	"github.com/tinrab/cautious-giggle"
	"github.com/tinrab/cautious-giggle/config"
	"log"
)

func main() {
	cfg := config.Config{
		StoreRepository: config.StoreRepositoryConfig{
			Address: "postgres://giggle:123456@giggle/giggle?sslmode=disable",
		},
	}
	a := cautious_giggle.NewApp(cfg)
	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
