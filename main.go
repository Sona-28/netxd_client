package main

import (
	"context"
	"fmt"
	"log"

	h "github.com/Sona-28/netxd_customer"
	tc "github.com/Sona-28/netxd_transaction"

	"github.com/Sona-28/netxd_customer_controllers/constants"

	"google.golang.org/grpc"
)

func Create(client h.CustomerServiceClient) {
	customer := &h.CustomerData{
		CustomerId: 101,
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
		CustomerId: 102,
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
	fmt.Println("Client running on ", constants.Port)
	con, err := grpc.Dial(constants.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed ", err)
	}
	defer con.Close()
	client := h.NewCustomerServiceClient(con)
	tclient := tc.NewTransactionServiceClient(con)
	fmt.Println("Enter a choice: \n 1.Create 2.Update 3.Read 4.Delete")
	var ch int
	fmt.Scan(&ch)
	if err!=nil{
		panic(err)
	}
	switch{
	case ch==1:
		Create(client)
	case ch==2:
		Update(client)
	case ch==3:
		Read(client)
	case ch==4:
		Delete(client)
	}
	res,err := tclient.TransferMoney(context.Background(), &tc.TransactionData{
		From:   101,
		To:     102,
		Amount: 1000,
	})
	if err!=nil{
		log.Fatal("failed,", err)
	}
	fmt.Println(res)

}
