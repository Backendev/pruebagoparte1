package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	url := "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/2019-12-01"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Request Error %v", err)
	} else {
		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err == nil {
			fmt.Printf("Res\n\v", res.Body)
			bodyRes, err := ioutil.ReadAll(res.Body)
			// contentType := res.Header.Get("Content-Type")
			if err == nil {
				res_b := string(bodyRes)
				list_res := strings.Split(res_b, "},{")
				for _, i := range list_res {
					clean_text := strings.Replace(i, "{", "", -1)
					clean_text = strings.Replace(clean_text, "}", "", -1)
					clean_text = strings.Replace(clean_text, "[", "", -1)
					clean_text = strings.Replace(clean_text, "]", "", -1)
					fmt.Println("____")
					fmt.Println(clean_text)
				}

			}

		}
	}
}
