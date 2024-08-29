package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/shania-kawa/Billing_Project/api"

	"google.golang.org/grpc"
)

func TestProcessPayment(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)

	}
	defer conn.Close()
	client := api.NewBilligServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &api.PaymentRequest{
		CardNumber: "4111111111111111",
		CardExpiry: "12/25",
		CardCvc:    "123",
		Amount:     100.00,
		Currency:   "USD",
	}
	res, err := client.ProcessPayment(ctx, req)
	if err != nil {
		t.Fatalf("could not procced payment :%v\n", err)

	}
	if !res.Success {
		t.Fatalf("excpected success, got failed: %v", res.Message)

	}
	log.Printf("transaction ID :%s", res.TransactionId)
	log.Printf("message: %s", res.Message)
}
