package main

import (
	"context"
	"fmt"
	"log"

	h "github.com/Sona-28/netxd_customer"
	"github.com/Sona-28/netxd_customer_controllers/constants"

	"google.golang.org/grpc"
)

func Create(client h.CustomerServiceClient) {
	customer := &h.CustomerData{
		CustomerId: 102,
		Firstname:  "Sona",
		Lastname:   "Sivasundari",
		BankId:     1001,
		Balance:    7000,
	}

	res, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatal("failed ", err)
	}
	fmt.Println("Response of Add: ", res)
}

func Read(client h.CustomerServiceClient) {
	custres, err := client.GetCustomer(context.Background(), &h.CustomerID{
		CustomerId: 101,
	})
	if err != nil {
		log.Fatal("failed ", err)
	}
	fmt.Println("Response of Get: ", custres)
}

func Update(client h.CustomerServiceClient) {
	res, err := client.UpdateCustomer(context.Background(), &h.UpdateCustomerRequest{
		Id:       102,
		Topic:    "last_name",
		Newvalue: "Styles",
	})
	if err != nil {
		log.Fatal("failed ", err)
	}
	fmt.Println("Response of Update: ", res)
}

func Delete(client h.CustomerServiceClient) {

	res, err := client.DeleteCustomer(context.Background(), &h.CustomerID{
		CustomerId: 102,
	})
	if err != nil {
		log.Fatal("failed ", err)
	}
	fmt.Println("Response of Delete: ", res)
}

func main() {
	con, err := grpc.Dial(constants.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed ", err)
	}
	defer con.Close()
	client := h.NewCustomerServiceClient(con)
	fmt.Println("1.Create 2.Update 3.Read 4.Delete")
	var ch int
	fmt.Scan(&ch)
	if err!=nil{
		panic(err)
	}
	fmt.Println(ch)
	switch{
	case ch==1:
		fmt.Println("Create")
		Create(client)
	case ch==2:
		Update(client)
	case ch==3:
		Read(client)
	case ch==4:
		Delete(client)
	}

}
