package movies

type MovieStorage interface {
	Lookup(title string) ([]MovieSearchRes, error)
}
