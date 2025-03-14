package product

type ServiceProduct interface {
	GetProducts() ([]Product, error)
	CreateProduct(input ProductInput) (Product, error)
	UpdateProduct(input ProductInput, id int) (Product, error)
	DeleteProduct(id int) (Product, error)
	GetOneProduct(id int) (Product, error)
}

type serviceProduct struct {
	repository RepositoryProduct
}

func NewServiceProduct(repository RepositoryProduct) *serviceProduct {
	return &serviceProduct{repository}
}

func (s *serviceProduct) GetProducts() ([]Product, error) {

	result, err := s.repository.FindAll()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *serviceProduct) GetOneProduct(id int) (Product, error) {

	result, err := s.repository.FindById(id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *serviceProduct) CreateProduct(input ProductInput) (Product, error) {
	product := Product{}

	product.Name = input.Name
	product.Sku = input.Sku //should create uuid
	product.Quantity = input.Quantity
	product.Location = input.Location
	product.Status = input.Status

	result, err := s.repository.Save(product)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *serviceProduct) UpdateProduct(input ProductInput, id int) (Product, error) {
	product, err := s.repository.FindById(id)

	product.Name = input.Name
	product.Sku = input.Sku //should update uuid
	product.Quantity = input.Quantity
	product.Location = input.Location
	product.Status = input.Status

	result, err := s.repository.Update(product)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *serviceProduct) DeleteProduct(id int) (Product, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return product, err
	}

	result, err := s.repository.Delete(product)

	if err != nil {
		return result, err
	}
	return result, nil
}
