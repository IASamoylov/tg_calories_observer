package help

//go:generate mockgen -source=dependecies.go -destination=dependecies_mocks.go -package=help . сommand

type сommand interface {
	Alias() string
	Description() string
}
