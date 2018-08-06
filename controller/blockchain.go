package controller

import (
	"DemoBlockchain/model"
	"encoding/json"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

func HandleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(model.Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
func HandleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m model.Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	model.Mutex.Lock()
	newBlock := model.GenerateBlock(model.Blockchain[len(model.Blockchain)-1], m.Info)
	model.Mutex.Unlock()

	if model.IsBlockValid(newBlock, model.Blockchain[len(model.Blockchain)-1]) {
		newBlockchain := append(model.Blockchain, newBlock)
		model.ReplaceChain(newBlockchain)
		spew.Dump(model.Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}
