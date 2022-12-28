package main

import (
	"context"
	"encoding/json"
	"go-learn/common/config"
	model "go-learn/common/model"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "n00b",
		Name:     "Gare",
		Password: "Tttest",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}
	// garage1 := model.Garage{
	// 	Id:   "q001",
	// 	Name: "Mobil garasi bang",
	// 	Coordinate: &model.GarageCoordinate{
	// 		Latitude:  45.123123123,
	// 		Longitude: 54.1231313123,
	// 	},
	// }

	user := serviceUser()

	user.Register(context.Background(), &user1)
	user.Register(context.Background(), &user1)

	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))
}
