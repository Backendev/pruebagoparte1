package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp := request(5)
	fmt.Println(resp)
}
func request(len_list int) []map[int]interface{} {
	var maps_i []map[int]interface{}
	client := &http.Client{}
	url := "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/2019-12-01"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Request Error %v", err)
	} else {
		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err == nil {
			bodyRes, err := ioutil.ReadAll(res.Body)
			if err == nil {
				res_b := string(bodyRes)
				list_res := strings.Split(res_b, "},{")
				type mt map[int]interface{}
				if len_list <= len(list_res) {
					for index := 0; index <= len_list; index++ {
						i := list_res[index]
						clean_text := strings.Replace(i, "{", "", -1)
						clean_text = strings.Replace(clean_text, "}", "", -1)
						clean_text = strings.Replace(clean_text, "[", "", -1)
						clean_text = strings.Replace(clean_text, "]", "", -1)
						list_items := strings.Split(clean_text, ",")
						m := make(map[string]string)
						for _, k := range list_items {
							list_fields := strings.Split(k, ":")
							value_field := ""
							for j := 1; j < len(list_fields); j++ {
								if j > 1 {
									value_field += ":" + list_fields[j]
								} else {
									value_field += list_fields[j]
								}

							}
							m[list_fields[0]] = value_field
						}
						map_item := mt{index: m}
						maps_i = append(maps_i, map_item)

					}
					return maps_i

				} else {
					log.Fatalf("The length list exceded to length the request %v", err)
				}

			}
		}

	}
	return maps_i
}
