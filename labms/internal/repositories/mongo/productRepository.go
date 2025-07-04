package repositories

import (
	"repogin/internal/db"
	"repogin/internal/models"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type ProductRepoInterface interface {
	Create(models.ProductModel) error
	Modify(models.ProductModel) error
	GetAll() ([]models.ProductModel, error)
	GetOne(models.ProductModel) (models.ProductModel, error)
	DeleteOne(models.ProductModel) error
}

// Mongo struct
type ProductRepoStruct struct {
	MongoRepo *db.MongoRepo
}

func NewProductRepo(mr *db.MongoRepo) *ProductRepoStruct {
	return &ProductRepoStruct{
		MongoRepo: mr,
	}
}

func (pr *ProductRepoStruct) Create(prod models.ProductModel) error {
	Collection := pr.MongoRepo.Client.Database(pr.MongoRepo.DBInfo.Database).Collection(os.Getenv("coll_products"))
	// prod.Id = uuid.New().String()
	prod.Id = strings.Replace(uuid.New().String(), "-", "", -1)
	res, err := Collection.InsertOne(context.Background(), prod)
	if err != nil {
		fmt.Println("ERROR : create", err)
		return err
	}
	fmt.Println("ERROR : product added into db -", res.InsertedID)
	return nil
}
func (pr *ProductRepoStruct) Modify(prod models.ProductModel) error {
	Collection := pr.MongoRepo.Client.Database(pr.MongoRepo.DBInfo.Database).Collection(os.Getenv("coll_products"))

	filter := bson.M{"id": prod.Id}
	opt := options.Update().SetUpsert(true)
	update := bson.M{
		"$set": bson.M{
			"name":        prod.Name,
			"description": prod.Description,
			"image":       prod.ImagePaath,
			"price":       prod.Price,
		},
	}
	res, RError := Collection.UpdateOne(context.TODO(), filter, update, opt)
	if RError != nil {
		// fmt.Println("ERROR : create", err)
		return RError
	}
	if res.ModifiedCount > 0 {
		fmt.Println("Modified ", prod.Id)
		return nil
	}
	return errors.New("ERROR_UPDATION_FAILED")
}
func (pr *ProductRepoStruct) GetAll() ([]models.ProductModel, error) {
	fmt.Println("Data ", os.Getenv("coll_products"))
	Collection := pr.MongoRepo.Client.Database(pr.MongoRepo.DBInfo.Database).Collection(os.Getenv("coll_products"))
	filter := bson.M{}
	var products []models.ProductModel
	Cursor, curerror := Collection.Find(context.Background(), filter)
	if curerror != nil {
		fmt.Println("ERROR GetAll", curerror)
		return []models.ProductModel{}, curerror
	}
	err := Cursor.All(context.TODO(), &products)
	if err != nil {
		fmt.Println("ERROR : GetAll", err)
		return []models.ProductModel{}, err
	}
	if len(products) > 0 {
		fmt.Println("Count of products", len(products))
		return products, nil
	}
	return []models.ProductModel{}, errors.New("ERROR_NO_PRODUCTS_FOUND")
}
func (pr *ProductRepoStruct) GetOne(prod models.ProductModel) (models.ProductModel, error) {
	Collection := pr.MongoRepo.Client.Database(pr.MongoRepo.DBInfo.Database).Collection(os.Getenv("coll_products"))
	filter := bson.M{
		"id": prod.Id,
	}
	var product models.ProductModel
	err := Collection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		fmt.Println("ERROR : GetOne", err)
		return models.ProductModel{}, err
	}
	return product, nil
}

func (pr *ProductRepoStruct) DeleteOne(prod models.ProductModel) error {
	// var product models.ProductModel
	// if product.Id != "" {
	// 	query := "Update " + os.Getenv("products") + " SET isDeleted = true  WHERE id = ?;"
	// 	res, reserr := pr.SRepo.Session.Exec(query, product.Id)
	// 	if reserr != nil {
	// 		fmt.Println("ERROR : DeleteOne ", reserr)
	// 		cnt, _ := res.RowsAffected()
	// 		if cnt > 0 {
	// 			fmt.Println("SOFT DELETED id = ", product.Id)
	// 			return nil
	// 		}
	// 		fmt.Println("ERROR_NOT_DELETED = ", product.Id)
	// 		return errors.New("ERROR_NOT_DELETED")
	// 	}
	// }
	//  else {
	return errors.New("ERROR_ID_IS_INVALID")
	// }
}

// // SQL struct
// type ProductSQLRepo struct {
// 	SRepo *db.SQLRepo
// }

// func NewProductSQLRepo(sr *db.SQLRepo) *ProductSQLRepo {
// 	return &ProductSQLRepo{
// 		SRepo: sr,
// 	}
// }

// func (pr *ProductSQLRepo) Create(p models.ProductModel) error {
// 	p.Id = "1"
// 	query := `INSERT INTO products (name,description,price,imagePath,category,inventoryStock) VALUES (?,?,?,?,?,?);`
// 	//  ON DUPLICATE KEY UPDATE
// 	//  name=VALUES(name),
// 	//  price = VALUES (price),
// 	//  description=VALUES(description),
// 	//  imagepath=VALUES(imagepath)
// 	res, rerr := pr.SRepo.Session.Exec(query, p.Name, p.Description, p.Price, p.ImagePaath, p.Category, p.InventoryStock)
// 	if rerr != nil {
// 		fmt.Println("ERROR : ProductSQLRepo Create ", rerr)
// 		return rerr
// 	}
// 	inc, _ := res.LastInsertId()
// 	rac, _ := res.RowsAffected()
// 	if inc > 0 || rac > 0 {
// 		fmt.Println("product Inserted / updated successfully ! with sql id - ", inc)
// 		return nil
// 	}
// 	return errors.New("ERROR_PRODUCT_NOT_INSERTED_UPDATED")
// }
// func (pr *ProductSQLRepo) Modify(p models.ProductModel) error {
// 	// query := `"INSERT INTO products (name,description,price,imagepath) VALUES (?,?,?)
// 	//  ON DUPLICATE KEY UPDATE
// 	//  name=VALUES(name),
// 	//  price = VALUES (price),
// 	//  description=VALUES(description),
// 	//  imagepath=VALUES(imagepath)"`
// 	query := "UPDATE products SET name = ? , price = ? , description = ? , imagePath = ? WHERE Id = ? ;"
// 	res, rerr := pr.SRepo.Session.Exec(query, p.Name, p.Price, p.Description, p.ImagePaath, p.Id)
// 	if rerr != nil {
// 		fmt.Println("ERROR : ProductSQLRepo Create ", rerr)
// 	}
// 	// inc, _ := res.LastInsertId()
// 	rac, _ := res.RowsAffected()

// 	if rac > 0 {
// 		fmt.Println("updated successfully !")
// 		return nil
// 	}
// 	return errors.New("ERROR_NOT_UPDATED")
// }
// func (pr *ProductSQLRepo) GetAll() ([]models.ProductModel, error) {
// 	var products []models.ProductModel
// 	query := "SELECT * FROM products "
// 	count, rerr := pr.SRepo.Session.SelectBySql(query).Load(&products)
// 	if rerr != nil {
// 		fmt.Println("ERROR : ProductSQLRepo Create ", rerr)
// 	}
// 	// inc, _ := res.LastInsertId()
// 	// rac, _ := res.RowsAffected()

// 	if count > 0 {
// 		fmt.Println("Count - ", count)
// 		return products, nil
// 	}
// 	return products, errors.New("ERROR_NO_DATA_FOUND")
// }
// func (pr *ProductSQLRepo) GetOne(prod models.ProductModel) (models.ProductModel, error) {
// 	var product models.ProductModel
// 	var query string
// 	var statement *dbr.SelectStmt
// 	if prod.Id != "" && prod.Name != "" {
// 		query = "SELECT * FROM products WHERE Id = ? AND name = ? ;"
// 		statement = pr.SRepo.Session.SelectBySql(query, prod.Id, prod.Name)
// 	} else {

// 		if prod.Id != "" {
// 			query = "SELECT * FROM products WHERE Id = ? ;"
// 			statement = pr.SRepo.Session.SelectBySql(query, prod.Id)

// 		} else if prod.Name != "" {
// 			query = "SELECT * FROM products WHERE name = ? ;"
// 			statement = pr.SRepo.Session.SelectBySql(query, prod.Name)
// 		}
// 	}

// 	err := statement.LoadOne(&product)
// 	if err != nil {
// 		fmt.Println("ERROR : ProductSQLRepo Create ", err)
// 		return product, err
// 	}
// 	return product, nil
// }
// func (pr *ProductSQLRepo) DeleteOne(product models.ProductModel) error {
// 	if product.Id != "" {
// 		query := "Update " + os.Getenv("products") + " SET isDeleted = true  WHERE id = ?;"
// 		res, reserr := pr.SRepo.Session.Exec(query, product.Id)
// 		if reserr != nil {
// 			fmt.Println("ERROR : DeleteOne ", reserr)
// 			cnt, _ := res.RowsAffected()
// 			if cnt > 0 {
// 				fmt.Println("SOFT DELETED id = ", product.Id)
// 				return nil
// 			}
// 			fmt.Println("ERROR_NOT_DELETED = ", product.Id)
// 			return errors.New("ERROR_NOT_DELETED")
// 		}
// 	}
// 	//  else {
// 	return errors.New("ERROR_ID_IS_INVALID")
// 	// }
// }
