package product

//go:generate mockgen -source=dependecies.go -destination=dependecies_mocks.go -package=product . command

type command interface {
	Alias() string
}
