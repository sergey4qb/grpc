package training_grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	pb "training_grpc/server/server"
)

type Server struct {
	pb.UnimplementedUserServer
}

func (s Server) CreateUser(ctx context.Context, data *pb.UserData) (*pb.Id, error) {
	id := uuid.New()
	data.Id.Uuid = id.ID()
	content, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("./server/server/filesystem/userfile.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return data.Id, nil
}

func (s Server) GetUserByID(ctx context.Context, uuid *pb.Id) (*pb.UserData, error) {
	var user pb.UserData
	//var id pb.Id
	jsonFile, err := os.Open("./server/server/filesystem/userfile.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	var byteValue, _ = ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.Id.Uuid, uuid.Uuid)

	if user.Id.Uuid == uuid.Uuid{
		return &user, nil
	}
	return &pb.UserData{},nil
}

func (s Server) UpdateUserByID(ctx context.Context, userdata *pb.UpdateUserData) (*pb.UserData, error) {
	var user pb.UserData
	//var id pb.Id
	jsonFile, err := os.Open("./server/server/filesystem/userfile.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	var byteValue, _ = ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.Id.Uuid, userdata.Id.Uuid)

	if user.Id.Uuid == userdata.Id.Uuid{
		user.Name = userdata.Userdata.Name
		user.Surname = userdata.Userdata.Surname
		user.CurrentLivePlace = userdata.Userdata.CurrentLivePlace
		user.CurrentPosition = userdata.Userdata.CurrentPosition
		content, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("./server/server/filesystem/userfile.json", content, 0644)
		if err != nil {
			log.Fatal(err)
		}
		return &user, nil
	}
	return &pb.UserData{},nil
}

func (s Server) DeleteUserByID(ctx context.Context, uuid *pb.Id) (*pb.Empty, error) {

	var user pb.UserData
	//var id pb.Id
	jsonFile, err := os.Open("./server/server/filesystem/userfile.json")
	if err != nil {
		return &pb.Empty{}, err
	}
	defer jsonFile.Close()
	var byteValue, _ = ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &user)
	if err != nil {
		return &pb.Empty{}, err
	}
	fmt.Println(user.Id.Uuid, uuid.Uuid)

	if user.Id.Uuid == uuid.Uuid{
		err := jsonFile.Close()
		if err != nil {
			return &pb.Empty{}, err
		}
		e := os.Remove("./server/server/filesystem/userfile.json")
		if e != nil {
			return &pb.Empty{}, err
		}
	}
	return &pb.Empty{}, err
}

func (s Server) mustEmbedUnimplementedUserServer() {
	panic("implement me")
}
