package movies

type Service struct {
	MovieStorage
	ListStorage
}

func NewService(
	movieStorage MovieStorage,
	listStorage ListStorage,
) *Service {
	return &Service{
		MovieStorage: movieStorage,
		ListStorage:  listStorage,
	}
}
