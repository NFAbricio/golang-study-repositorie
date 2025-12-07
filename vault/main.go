package main

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"

)

func main() {
	ctx := context.Background()

	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
		vault.WithRequestTimeout(30*time.Second),		
	)
	if err != nil {
		panic(err)
	}

	//with token authentication (secure)
	if err := client.SetToken("my-token"); err != nil { // replace "my-token" with a functional token
		panic(err)
	}

	// Key/Value Version 2 secrets engine -> API method for managing secrets (create, read, update, delete)
	_, err = client.Secrets.KvV2Write(ctx, "foo", schema.KvV2WriteRequest{
		Data: map[string]any{
			"password1": "9k9BsVZnbw6rVvnx",
			"password2": "jA61usDY8tk5Wyg2",
		}},
		vault.WithMountPath("secret"),
	) 

	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret written succesfully")

	// read the secret
	secret, err := client.Secrets.KvV2Read(ctx, "foo", vault.WithMountPath("secret"))
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("secret retrived: ", secret.Data.Data)
}