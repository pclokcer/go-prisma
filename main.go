package main

import (
	"context"
	"encoding/json"
	"fmt"
	relations "go-prisma/generatedPrisma"
)

var client *relations.PrismaClient

func main() {

	client = relations.NewClient()
	if err := client.Connect(); err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			fmt.Println("could not disconnect: %w", err)
		}
	}()

	usermodel, _ := Insert()
	userData, _ := Get(usermodel)

	result, _ := json.MarshalIndent(userData, "", "  ")
	fmt.Printf("post: %s\n", result)
}

func Insert() (*relations.UsersModel, error) {
	return client.Users.CreateOne(
		relations.Users.Name.Set("erdem"),
		relations.Users.Surname.Set("un"),
		relations.Users.Phone.Set("905111111111"),
	).Exec(context.Background())
}

func Get(usermodel *relations.UsersModel) (*relations.UsersModel, error) {
	return client.Users.FindFirst(
		relations.Users.ID.Equals(usermodel.ID),
	).Exec(context.Background())
}
