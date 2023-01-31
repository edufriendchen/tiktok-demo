package main

import (
	"context"
	"log"
	"time"

	userservice "github.com/edufriendchen/tiktok-demo/kitex_gen/user/userservice"

	user "github.com/edufriendchen/tiktok-demo/kitex_gen/user"

	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := userservice.NewClient("register", client.WithHostPorts("0.0.0.0:9000"))
	if err != nil {
		log.Fatal(err)
	}
	req := &user.CheckUserRequest{Username: "admin200", Password: "04291"}
	resp, err := client.CheckUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

	time.Sleep(time.Second)
}
