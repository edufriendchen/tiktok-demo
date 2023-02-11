package main

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/user"
	userservice "github.com/edufriendchen/tiktok-demo/kitex_gen/user/userservice"
)

func main() {
	client, err := userservice.NewClient("user", client.WithHostPorts("0.0.0.0:9000"))
	if err != nil {
		log.Fatal(err)
	}
	req := &user.CreateUserRequest{Username: "admin_test", Password: "123456"}
	resp, err := client.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
