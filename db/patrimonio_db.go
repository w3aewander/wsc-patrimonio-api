package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type patrimonio struct {
	id string `json:"id"`
	marca string `json:"marca"`
	modelo string `json:"modelo"`
	local string `json:"local"`
}


func (p *patrimonio) GetAll(ctx context.Context){
	return 
}

func (p *patrimonio) Show(ctx context.Context, id string) ([]patrimonio.Patrimonio, error){

	return nil, nil
}

func (p *patrimonio) Update(ctx context.Context, id string) (bool, error) {

	return false, nil
}

func (p *patrimonio) Delete (ctx context.Context, id string) (bool, error) {
	return false, nil
}

func (p *patrimonio) Add (ctx context.Context) (bool, error){
	return false, nil
}

func (*Patrimonio) dbConnect() {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client()
		.ApplyURI("mongodb+srv://wsc_patrimonio_app:<password>@cluster0.azqyj9u.mongodb.net/?retryWrites=true&w=majority")
		.SetServerAPIOptions(serverAPIOptions)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	CheckError(err)

	// get collection as ref
	collection := client.Database("wsc_patrimonio_app").Collection("patrimonio")

}
