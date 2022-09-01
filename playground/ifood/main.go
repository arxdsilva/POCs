package main

import (
	"fmt"
	"log"
	"os"
	"time"

	sdk "github.com/arxdsilva/golang-ifood-sdk/container"
)

func main() {
	var clientID, clientSecret string
	clientID = "restauranti"
	clientSecret = os.Getenv("SECRET")
	container := sdk.New(0, time.Minute)
	container.GetHttpAdapter()
	auth := container.GetAuthenticationService(clientID, clientSecret)
	user := "work@insurgencygames.com"
	password := os.Getenv("PASSWORD")
	creds, err := auth.Authenticate(user, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("creds: %+v\n", creds.ExpiresIn)
	container.AuthService.Validate()
	merchant := container.GetMerchantService()
	merchants, err := merchant.ListAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("merchants: %+v\n", merchants)
	catalog := container.GetCatalogService()
	catalogs, err := catalog.ListAll(merchants[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("catalog: %+v\n", catalogs)
}

// func main() {
// url := "https://pos-api.ifood.com.br/oauth/token"
// 	method := "POST"
// 	payload := &bytes.Buffer{}
// 	writer := multipart.NewWriter(payload)
// 	_ = writer.WriteField("client_id", "restauranti")
// 	_ = writer.WriteField("client_secret", "LDA4Fw7H")
// 	_ = writer.WriteField("grant_type", "password")
// 	_ = writer.WriteField("username", "work@insurgencygames.com")
// 	_ = writer.WriteField("password", "JC%RC34a")
// 	err := writer.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, payload)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(body))
// }
