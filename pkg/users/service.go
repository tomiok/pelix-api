package users

type UserService struct {
	UserStorage
}

func NewService(storage UserStorage) *UserService {
	return &UserService{
		UserStorage: storage,
	}
}
