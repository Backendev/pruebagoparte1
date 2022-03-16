package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	url := "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/2019-12-01"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Request Error %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	fmt.Printf("Res\n\v", res.Body)
	cuerpoRespuesta, err := ioutil.ReadAll(res.Body)
	fmt.Printf("Res\n\v", string(cuerpoRespuesta))

}
