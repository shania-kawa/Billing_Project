package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/shania-kawa/Billing-Project/api"

	"github.com/shania-kawa/Billing-Project/pkg/db"
)

type BillingService struct {
	api.UnimplementedBilligServiceServer
}

func (s *BillingService) ProcessPyament(ctx context.Context, req *api.PaymentRequest) (*api.PaymentResponse, error) {
	if len(req.CardNumber) != 16 {
		return &api.PaymentResponse{
			Success:       false,
			TransactionId: "",
			Message:       "Invalid Card number ",
		}, fmt.Errorf("invalid card number")
	}
	TransactionID := fmt.Sprintf("txn_%d", rand.Intn(1000000))
	paymentSuccess := true

	conn, err := db.DB.Acquire(ctx)
	if err != nil {
		log.Fatalf("unable to acquire database connection: %v", err)

	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `
	INSERT INTO transactions (user_id, transaction_id,amount,currency,status)
	VALUES ($1, $2, $3, $4, $5)`,
		req.UserID, req.transactionID, req.Amount, req.Currency, "SUCCESS")
	if err != nil {
		log.Printf("failed to store transaction: %v", err)
		return &api.PaymentResponse{
			Success:       false,
			TransactionID: TransactionID,
			Message:       " Payment successful but failed to store transaction",
		}, err

	}
	return &api.PaymentResponse{
		Success:       paymentSuccess,
		TransactionID: TransactionID,
		Message:       "payment successful",
	}, nil
}
