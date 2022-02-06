package movies

type Service struct {
	MovieStorage
}

func NewService(stg MovieStorage) *Service {
	return &Service{
		MovieStorage: stg,
	}
}
