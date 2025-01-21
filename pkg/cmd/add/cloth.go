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

func AddCloth(cloth item.Cloth) {
	url := constants.HOST + ":" + constants.PORT + "/items/cloth"
	res, err := http.Post(url, "application/json", prepareClothBody(cloth))
	if err != nil {
		fmt.Println("Error in adding cloth")
	}
	defer res.Body.Close()
	fmt.Println("Added a cloth")
}

func prepareClothBody(cloth item.Cloth) io.Reader {
	json, _ := json.Marshal(cloth)
	requestBody := bytes.NewBuffer(json)
	return requestBody
}
