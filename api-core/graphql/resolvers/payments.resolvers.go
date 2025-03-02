package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"log"
	"math"

	"github.com/polatbilal/gqlgen-ent/api-core/ent"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/jobdetail"
	"github.com/polatbilal/gqlgen-ent/api-core/graphql/model"
	"github.com/polatbilal/gqlgen-ent/api-core/middlewares"
)

// CreateJobPayments is the resolver for the createJobPayments field.
func (r *mutationResolver) CreateJobPayments(ctx context.Context, input model.JobPaymentsInput) (*ent.JobPayments, error) {
	client := middlewares.GetClientFromContext(ctx)

	// İş detayını bul
	jobDetail, err := client.JobDetail.Query().
		Where(jobdetail.YibfNoEQ(*input.YibfNo)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// JobRelations'ı bul
	relations, err := jobDetail.QueryRelations().Only(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("jobDetail", jobDetail.ID)

	var totalPayment float64
	if *input.TotalPayment == 0 && *input.Amount > 0 {
		// Amount'dan %20 çıkar
		totalPayment = math.Round(*input.Amount*0.80*100) / 100
	} else {
		totalPayment = *input.TotalPayment
	}

	payment, err := client.JobPayments.Create().
		SetPaymentNo(input.PaymentNo).
		SetPaymentDate(*input.PaymentDate).
		SetPaymentType(*input.PaymentType).
		SetTotalPayment(totalPayment).
		SetLevelRequest(*input.LevelRequest).
		SetLevelApprove(*input.LevelApprove).
		SetAmount(*input.Amount).
		SetState(*input.State).
		SetYibfNo(*input.YibfNo).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Payment'ı JobRelations'a ekle
	_, err = relations.Update().
		AddPayments(payment).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// UpdateJobPayments is the resolver for the updateJobPayments field.
func (r *mutationResolver) UpdateJobPayments(ctx context.Context, id int, input model.JobPaymentsInput) (*ent.JobPayments, error) {
	client := middlewares.GetClientFromContext(ctx)

	var totalPayment float64
	if *input.TotalPayment == 0 && *input.Amount > 0 {
		// Amount'dan %20 çıkar
		totalPayment = math.Round(*input.Amount*0.80*100) / 100
	} else {
		totalPayment = *input.TotalPayment
	}

	payment, err := client.JobPayments.UpdateOneID(id).
		SetTotalPayment(totalPayment).
		SetLevelApprove(*input.LevelApprove).
		SetAmount(*input.Amount).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

// DeleteJobPayments is the resolver for the deleteJobPayments field.
func (r *mutationResolver) DeleteJobPayments(ctx context.Context, id int) (*ent.JobPayments, error) {
	client := middlewares.GetClientFromContext(ctx)
	payment, err := client.JobPayments.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	err = client.JobPayments.DeleteOne(payment).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

// JobPayments is the resolver for the jobPayments field.
func (r *queryResolver) JobPayments(ctx context.Context, yibfNo int) ([]*ent.JobPayments, error) {
	client := middlewares.GetClientFromContext(ctx)

	// İş detayını bul
	jobDetail, err := client.JobDetail.Query().
		Where(jobdetail.YibfNoEQ(yibfNo)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// JobRelations üzerinden ödemeleri getir
	relations, err := jobDetail.QueryRelations().Only(ctx)
	if err != nil {
		return nil, err
	}

	payments, err := relations.QueryPayments().All(ctx)
	if err != nil {
		return nil, err
	}

	return payments, nil
}
