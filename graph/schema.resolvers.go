package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/afrizal423/go-supermarket-gpql/app/customer"
	"github.com/afrizal423/go-supermarket-gpql/graph/generated"
	"github.com/afrizal423/go-supermarket-gpql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Customer(ctx context.Context, limit *int, offset *int) ([]*model.Customer, error) {
	// fmt.Println(customer.GetAllData())
	var result []*model.Customer
	// var dt []customer.Customers
	dt, _ := customer.GetAllData(limit, offset)
	for _, link := range dt {
		// fmt.Println(link)
		result = append(result,
			&model.Customer{ID: link.CustomerID, CustomerName: link.CustomerNames,
				Segment: link.Segment, Age: link.Age, Country: link.Country, City: link.City,
				State: link.State, PostalCode: link.PostalCode, Region: link.Region})
	}
	// panic(fmt.Errorf("not implemented"))
	return result, nil
}

func (r *queryResolver) CariCustomerByID(ctx context.Context, id string) ([]*model.Customer, error) {
	var result []*model.Customer
	dt, _ := customer.GetDataByID(id)
	for _, link := range dt {
		// fmt.Println(link)
		result = append(result,
			&model.Customer{ID: link.CustomerID, CustomerName: link.CustomerNames,
				Segment: link.Segment, Age: link.Age, Country: link.Country, City: link.City,
				State: link.State, PostalCode: link.PostalCode, Region: link.Region})
	}
	// panic(fmt.Errorf("not implemented"))
	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
