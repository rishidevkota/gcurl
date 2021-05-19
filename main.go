package main

import "gcurl/cmd"

func main() {
	// requestStr := `
	// {
	// 	"method": "GET",
	// 	"url": "https://newweb.nepalstock.com/api/nots/security/2748",
	// 	"headers":{
	// 		"accept": "application/json, text/plain, */*",
	// 		"user-agent": "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Mobile Safari/537.36"
	// 	}
	// }
	// `
	// var req client.Request
	// err := json.Unmarshal([]byte(requestStr), &req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(req)
	// client.Do(req)
	cmd.Execute()
}
