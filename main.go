package main

import (
	"context"
	"fmt"
	"log"
	"github.com/Sona-28/netxd_customer_controllers/constants"
	h "github.com/Sona-28/netxd_customer"

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

	Create(client)

	// Update(client)

	// Read(client)

	// Delete(client)

}
