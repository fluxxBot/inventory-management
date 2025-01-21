package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.jfrog.info/kanishkg/inventory-management/constants"
	"git.jfrog.info/kanishkg/inventory-management/item"
	"io"
	"net/http"
)

func AddBook(book item.Book) {
	url := constants.HOST + ":" + constants.PORT + "/items/books"
	res, err := http.Post(url, "application/json", prepareBookBody(book))
	if err != nil {
		fmt.Println("Error in adding book")
	}
	defer res.Body.Close()
	fmt.Println("Added a book")
}

func prepareBookBody(book item.Book) io.Reader {
	json, _ := json.Marshal(book)
	requestBody := bytes.NewBuffer(json)
	return requestBody
}
