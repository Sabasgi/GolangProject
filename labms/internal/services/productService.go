package services

import (
	"repogin/internal/models"
	repositories "repogin/internal/repositories/sql"
)

// type ProdServiceInterface interface {
// }
type ProdServiceStruct struct {
	ProdService repositories.ProductRepoInterface
}

func NewProdService(repo repositories.ProductRepoInterface) *ProdServiceStruct {
	return &ProdServiceStruct{
		ProdService: repo,
	}
}
func (ph *ProdServiceStruct) CreateProductService(prod models.ProductModel) error {
	Error := ph.ProdService.Create(prod)
	if Error != nil {
		// fmt.Println("ERROR : CreateProduct", Error)
		return Error
	}
	return nil
}
func (ph *ProdServiceStruct) UpdateProductService(prod models.ProductModel) error {
	Error := ph.ProdService.Modify(prod)
	if Error != nil {
		return Error
	}
	return nil
}
func (ph *ProdServiceStruct) GetAllProductsService() ([]models.ProductModel, error) {
	products, perr := ph.ProdService.GetAll()
	if perr != nil {
		return products, perr
	}
	return products, nil
}
func (ph *ProdServiceStruct) GetOneProductService(prd models.ProductModel) (models.ProductModel, error) {

	product, perr := ph.ProdService.GetOne(prd)
	if perr != nil {
		return product, perr
	}
	return product, nil
}
func (ph *ProdServiceStruct) DeleteOneService(prd models.ProductModel) error {

	perr := ph.ProdService.DeleteOne(prd)
	if perr != nil {
		return perr
	}
	return perr
}
