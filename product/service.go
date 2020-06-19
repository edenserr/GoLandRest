package product

type Service interface {
	GetProductbyId(param *getProductByIDRequest) (*Product, error)
	GetProducts(param *getProductRequest) (*ProductList, error)
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

func (s *service) GetProducts(params *getProductRequest) (*ProductList, error) {
	product, err := s.repo.GetProducts(params)
	if err != nil {
		panic(err)
	}
	totalProducts, err := s.repo.GetTotalProducts()
	if err != nil {
		panic(err)
	}
	return &ProductList{Data: product, TotalRecords: totalProducts}, nil
}
