package handlers

import (
	"fmt"
	"net/http"
	"repogin/internal/models"

	"github.com/gin-gonic/gin"
)

// type ProductHandler struct {
// 	service services.ProdServiceStruct //implemented interface here  // ite will act as a service to access repository struct.
// }

// func NewProductHandler(P services.ProdServiceStruct) *ProductHandler {
// 	return &ProductHandler{
// 		service: P,
// 	}
// }

func (ph *MainHandlers) CreateProductRoute(c *gin.Context) {
	var productDetails models.ProductModel
	// Pr := repositories.NewProductRepo("products")
	Error := c.Bind(&productDetails)
	if Error != nil {
		fmt.Println("ERROR : CreateProductRoute", Error)
		// c.JSON()
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := ph.Ps.CreateProductService(productDetails)
	if Err != nil {
		fmt.Println("Err : CreateProduct", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Product created successfully"})
	return
}
func (ph *MainHandlers) UpdateProductRoute(c *gin.Context) {
	var productDetails models.ProductModel
	// Pr := repositories.NewProductRepo("products")
	Error := c.Bind(&productDetails)
	if Error != nil {
		fmt.Println("ERROR : UpdateProductService", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := ph.Ps.UpdateProductService(productDetails)
	if Err != nil {
		fmt.Println("ERROR : UpdateProductService", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})

		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Product updated successfully"})
	return
}
func (ph *MainHandlers) GetAllProductsRoute(c *gin.Context) {
	products, err := ph.Ps.GetAllProductsService()
	if err != nil {
		fmt.Println("ERROR : GetAllProductsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
	return
}
func (ph *MainHandlers) GetOneProductRoute(c *gin.Context) {
	var prd models.ProductModel
	binderror := c.Bind(&prd)
	if binderror != nil {
		fmt.Println("ERROR : GetOneProductRoute", binderror)
	}
	product, perr := ph.Ps.ProdService.GetOne(prd)
	if perr != nil {
		fmt.Println("ERROR : GetAllProductsRoute", perr)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": perr.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
	return
}
func (ph *MainHandlers) DeleteOneRoute(c *gin.Context) {
	var prd models.ProductModel
	binderror := c.Bind(&prd)
	if binderror != nil {
		fmt.Println("ERROR : GetOneProductRoute", binderror)
	}
	perr := ph.Ps.ProdService.DeleteOne(prd)
	if perr != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": perr.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"error": perr.Error()})
	return
}
