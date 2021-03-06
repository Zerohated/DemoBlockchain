package controller

import (
	"DemoBlockchain/model"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/kataras/iris"

	"github.com/davecgh/go-spew/spew"
)

var (
	mutex = &sync.Mutex{}
)

func HandleGetBlockchain(ctx iris.Context) {
	bytes, err := json.MarshalIndent(model.Blockchain, "", "  ")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.Writef(string(bytes))
}

func HandleWriteBlock(ctx iris.Context) {
	ctx.Header("Content-Type", "application/json")
	var message model.Message
	// message := model.Message{}
	if err := ctx.ReadJSON(&message); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	fmt.Printf("Received: %#+v\n", message)

	mutex.Lock()
	newBlock := model.GenerateBlock(model.Blockchain[len(model.Blockchain)-1], message.Info)
	mutex.Unlock()

	if model.IsBlockValid(newBlock, model.Blockchain[len(model.Blockchain)-1]) {
		newBlockchain := append(model.Blockchain, newBlock)
		model.ReplaceChain(newBlockchain)
		spew.Dump(model.Blockchain)
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(newBlock)

}
