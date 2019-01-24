package graph

import (
	"context"

	"github.com/widnyana/grepe/models"
)

type Resolver struct{}

func (r *Resolver) Customer() CustomerResolver {
	return &customerResolver{r}
}
func (r *Resolver) Film() FilmResolver {
	return &filmResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Store() StoreResolver {
	return &storeResolver{r}
}

type customerResolver struct{ *Resolver }

func (r *customerResolver) ID(ctx context.Context, obj *models.Customer) (string, error) {
	panic("not implemented")
}
func (r *customerResolver) Address(ctx context.Context, obj *models.Customer) (string, error) {
	panic("not implemented")
}
func (r *customerResolver) City(ctx context.Context, obj *models.Customer) (string, error) {
	panic("not implemented")
}
func (r *customerResolver) Country(ctx context.Context, obj *models.Customer) (string, error) {
	panic("not implemented")
}
func (r *customerResolver) District(ctx context.Context, obj *models.Customer) (string, error) {
	panic("not implemented")
}
func (r *customerResolver) FirstName(ctx context.Context, obj *models.Customer) (string, error) {
	panic("not implemented")
}
func (r *customerResolver) LastName(ctx context.Context, obj *models.Customer) (string, error) {
	panic("not implemented")
}
func (r *customerResolver) Phone(ctx context.Context, obj *models.Customer) (string, error) {
	panic("not implemented")
}
func (r *customerResolver) Rentals(ctx context.Context, obj *models.Customer) ([]models.Rental, error) {
	panic("not implemented")
}

type filmResolver struct{ *Resolver }

func (r *filmResolver) ID(ctx context.Context, obj *models.Film) (string, error) {
	panic("not implemented")
}
func (r *filmResolver) Actors(ctx context.Context, obj *models.Film) ([]models.Actor, error) {
	panic("not implemented")
}
func (r *filmResolver) Category(ctx context.Context, obj *models.Film) (string, error) {
	panic("not implemented")
}
func (r *filmResolver) Description(ctx context.Context, obj *models.Film) (string, error) {
	panic("not implemented")
}
func (r *filmResolver) Length(ctx context.Context, obj *models.Film) (int, error) {
	panic("not implemented")
}
func (r *filmResolver) Rating(ctx context.Context, obj *models.Film) (string, error) {
	panic("not implemented")
}
func (r *filmResolver) RentalDuration(ctx context.Context, obj *models.Film) (int, error) {
	panic("not implemented")
}
func (r *filmResolver) ReplacementCost(ctx context.Context, obj *models.Film) (float64, error) {
	panic("not implemented")
}
func (r *filmResolver) SpecialFeatures(ctx context.Context, obj *models.Film) (string, error) {
	panic("not implemented")
}
func (r *filmResolver) Title(ctx context.Context, obj *models.Film) (string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateCustomer(ctx context.Context, input models.NewCustomer) (models.Customer, error) {
	panic("not implemented")
}
func (r *mutationResolver) RentFilm(ctx context.Context, input models.RentFilm) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) NewFilm(ctx context.Context, input models.NewFilm) (models.Film, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Film(ctx context.Context) ([]models.Film, error) {
	panic("not implemented")
}
func (r *queryResolver) Store(ctx context.Context) ([]models.Store, error) {
	panic("not implemented")
}
func (r *queryResolver) Customer(ctx context.Context) ([]models.Customer, error) {
	panic("not implemented")
}

type storeResolver struct{ *Resolver }

func (r *storeResolver) ID(ctx context.Context, obj *models.Store) (string, error) {
	panic("not implemented")
}
func (r *storeResolver) Address(ctx context.Context, obj *models.Store) (string, error) {
	panic("not implemented")
}
func (r *storeResolver) City(ctx context.Context, obj *models.Store) (string, error) {
	panic("not implemented")
}
func (r *storeResolver) Country(ctx context.Context, obj *models.Store) (string, error) {
	panic("not implemented")
}
func (r *storeResolver) ManagerStaffID(ctx context.Context, obj *models.Store) (int, error) {
	panic("not implemented")
}
func (r *storeResolver) Inventory(ctx context.Context, obj *models.Store) ([]models.Inventory, error) {
	panic("not implemented")
}
