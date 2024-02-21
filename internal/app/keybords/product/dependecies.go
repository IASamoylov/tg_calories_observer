package product

//go:generate mockgen -source=dependecies.go -destination=dependecies_mocks.go -package=product . inlineButton

type inlineButton interface {
	Text() string
	Callback() string
}
