package main

import (
	"log"
	"time"

	"DemoBlockchain/controller"
	"DemoBlockchain/model"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

var (
	// bcServer handles incoming concurrent model.Blocks
	bcServer chan []model.Block
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := model.Block{}
		genesisBlock = model.Block{0, t.String(), "genesisBlock is here", model.CalculateHash(genesisBlock), "", model.Difficulty, ""}
		spew.Dump(genesisBlock)

		model.Blockchain = append(model.Blockchain, genesisBlock)
	}()
	// httpAddr := os.Getenv("PORT")
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())
	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", controller.HandleGetBlockchain)
	app.Handle("POST", "/", controller.HandleWriteBlock)
	// Server Configuration
	// serverConfig := &http.Server{
	// 	Addr:           ":" + httpAddr,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("./configs/iris.tml")))
}
