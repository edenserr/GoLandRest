package product

type Service interface {
	GetProductbyId(param *getProductByIDRequest) (*Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetProductbyId(param *getProductByIDRequest) (*Product, error) {
	return s.repo.GetProductById(param.ProductID)
}
