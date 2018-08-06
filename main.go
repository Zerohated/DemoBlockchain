package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"DemoBlockchain/controller"
	"DemoBlockchain/model"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	// bcServer handles incoming concurrent model.Blocks
	bcServer chan []model.Block
)

func run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("PORT")
	log.Println("Listening on ", os.Getenv("PORT"))
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", controller.HandleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", controller.HandleWriteBlock).Methods("POST")
	return muxRouter
}

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

		model.Mutex.Lock()
		model.Blockchain = append(model.Blockchain, genesisBlock)
		model.Mutex.Unlock()
	}()
	log.Fatal(run())

}
