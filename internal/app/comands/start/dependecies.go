package start

//go:generate mockgen -source=dependecies.go -destination=dependecies_mocks.go -package=start . helpCommand,keyboardButton

type helpCommand interface {
	Alias() string
}

type keyboardButton interface {
	Text() string
}
