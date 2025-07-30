package product

type ProductUsecase struct {
	repo ProductRepositoryInterface
}

func NewProductUsecase(repo ProductRepositoryInterface) *ProductUsecase {
	return &ProductUsecase{repo}
}

func (u *ProductUsecase) GetProduct(id int) (Product, error) {
	return u.repo.GetProduct(id)
}
