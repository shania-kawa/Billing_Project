package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/shania-kawa/Billing_Project/api"

	"github.com/shania-kawa/Billing_Project/pkg/db"
)

type BillingService struct {
	api.UnimplementedBilligServiceServer
}

func (s *BillingService) ProcessPayment(ctx context.Context, req *api.PaymentRequest) (*api.PaymentResponse, error) {

	if len(req.CardNumber) != 16 {
		return &api.PaymentResponse{
			Success: false,
			Message: "Invalid Card number ",
		}, fmt.Errorf("invalid card number")
	}
	transaction_id := fmt.Sprintf("txn_%d", rand.Intn(1000000))
	paymentSuccess := true

	conn, err := db.DB.Acquire(ctx)
	if err != nil {
		log.Fatalf("unable to acquire database connection: %v", err)

	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `
	INSERT INTO transactions (transaction_id,amount,currency,status)
	VALUES ($1, $2, $3, $4)`,
		transaction_id, req.Amount, req.Currency, "SUCCESS")
	if err != nil {
		log.Printf("failed to store transaction: %v", err)
		return &api.PaymentResponse{
			Success:       false,
			TransactionId: transaction_id,
			Message:       " Payment successful but failed to store transaction",
		}, err

	}
	return &api.PaymentResponse{
		Success:       paymentSuccess,
		TransactionId: transaction_id,
		Message:       "payment successful",
	}, nil
}

//webhook

func (s *BillingService) HandleWebhook(ctx context.Context, req *api.WebhookRequest) (*api.WebhookReponse, error) {
	conn, err := db.DB.Acquire(ctx)
	if err != nil {
		log.Fatalf("unable to acquire database connection: %v \n", err)

	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `
	UPADRW transactions
	SET status =$1
	WHERE transaction_id =$2`,
		req.Status, req.TransactionId)
	if err != nil {
		log.Fatalf("failed to update transaction status :%v", err)
		return &api.WebhookReponse{
			Success: false,
			Message: "failed to update transactions",
		}, err
	}
	return &api.WebhookReponse{
		Success: true,
		Message: "transcation was successful",
	}, nil
}
