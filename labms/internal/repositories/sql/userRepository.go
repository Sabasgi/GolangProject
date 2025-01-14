package repositories

import (
	"repogin/internal/db"
	"repogin/internal/models"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type UserRepo interface {
	Create(models.UserModel) error
	Update(models.UserModel) error
	GetAll() ([]models.UserModel, error)
	GetOne(u models.UserModel) (models.UserModel, error)
}

// type UserRepoStruct struct {
// 	mongoRepo *db.MongoRepo
// }

// func NewUserRepo(mr *db.MongoRepo) *UserRepoStruct {
// 	return &UserRepoStruct{
// 		mongoRepo: mr,
// 	}
// }
// func (ur *UserRepoStruct) Create(user models.UserModel) error {
// 	user.UserId = strings.Replace(uuid.New().String(), "-", "", -1)
// 	Collection := ur.mongoRepo.Client.Database(ur.mongoRepo.DBInfo.Database).Collection(os.Getenv("coll_users"))
// 	_, err := Collection.InsertOne(context.Background(), user)
// 	if err != nil {
// 		fmt.Println("ERROR : Create - User", err)
// 		return err
// 	}
// 	return nil
// }
// func (ur *UserRepoStruct) Update(user models.UserModel) error {
// 	Collection := ur.mongoRepo.Client.Database(ur.mongoRepo.DBInfo.Database).Collection(os.Getenv("coll_users"))

// 	filter := bson.M{"id": user.UserId}
// 	opt := options.Update().SetUpsert(true)
// 	update := bson.M{
// 		"$set": bson.M{
// 			"name":     user.Username,
// 			"age":      user.Age,
// 			"emailId":  user.EmailId,
// 			"username": user.Username,
// 			"password": user.Password,
// 		},
// 	}
// 	res, RError := Collection.UpdateOne(context.TODO(), filter, update, opt)
// 	if RError != nil {
// 		fmt.Println("ERROR : create user", RError)
// 		return RError
// 	}
// 	if res.ModifiedCount > 0 {
// 		fmt.Println("Modified USER ", user.UserId)
// 		return nil
// 	}
// 	return errors.New("ERROR_UPDATION_FAILED")
// }

// func (ur *UserRepoStruct) GetAll() ([]models.UserModel, error) {
// 	Collection := ur.mongoRepo.Client.Database(ur.mongoRepo.DBInfo.Database).Collection(os.Getenv("coll_users"))
// 	filter := bson.M{}
// 	var users []models.UserModel
// 	Cursor, curerror := Collection.Find(context.Background(), filter)
// 	if curerror != nil {
// 		fmt.Println("ERROR GetAll", curerror)
// 		return []models.UserModel{}, curerror
// 	}
// 	err := Cursor.All(context.TODO(), &users)
// 	if err != nil {
// 		fmt.Println("ERROR : GetAll", err)
// 		return []models.UserModel{}, err
// 	}
// 	if len(users) > 0 {
// 		fmt.Println("Count of users ", len(users))
// 		return users, nil
// 	}
// 	return []models.UserModel{}, errors.New("ERROR_NO_PRODUCTS_FOUND")
// }
// func (ur *UserRepoStruct) GetOne(u models.UserModel) (models.UserModel, error) {
// 	Collection := ur.mongoRepo.Client.Database(ur.mongoRepo.DBInfo.Database).Collection(os.Getenv("coll_users"))
// 	filter := bson.M{
// 		"id": u.UserId,
// 	}
// 	var user models.UserModel
// 	err := Collection.FindOne(context.Background(), filter).Decode(&user)
// 	if err != nil {
// 		fmt.Println("ERROR : GetOne", err)
// 		return models.UserModel{}, err
// 	}
// 	return user, nil
// }

// SQL Repo
type UserRepoSQL struct {
	sqlrepo *db.SQLRepo
}

func NewUserSQLRepo(sr *db.SQLRepo) *UserRepoSQL {
	return &UserRepoSQL{
		sqlrepo: sr,
	}
}
func (ur *UserRepoSQL) Create(user models.UserModel) error {
	user.UserId = strings.Replace(uuid.New().String(), "-", "", -1)
	query := "INSERT INTO " + os.Getenv("coll_users") + " (userId,name,age,emailId,password,username) VALUES (?,?,?,?,?,?);"
	res, rerr := ur.sqlrepo.Session.Exec(query, user.UserId, user.Name, user.Age, user.EmailId, user.Password, user.Username)
	if rerr != nil {
		fmt.Println("ERROR : UserRepoSQL Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("user Inserted / updated successfully ! with sql id - ", inc)
		return nil
	}
	return errors.New("ERROR_USER_NOT_INSERTED_UPDATED")
	// return nil
}
func (ur *UserRepoSQL) Update(user models.UserModel) error {
	query := "UPDATE " + os.Getenv("coll_users") + " SET name = ? , emailId = ? , username = ? , password = ? , age = ? WHERE  id = ? AND userId = ? ;"
	res, rerror := ur.sqlrepo.Session.Exec(query, user.Name, user.EmailId, user.Username, user.Password, user.Age, user.Id, user.UserId)
	if rerror != nil {
		fmt.Println("ERROR Update users ", user.UserId)
		return rerror
	}
	cnt, _ := res.RowsAffected()
	if cnt > 0 {
		return nil
	}
	return errors.New("ERROR_USER_UPDATION_FAILED")
}

func (ur *UserRepoSQL) GetAll() ([]models.UserModel, error) {
	query := "SELECT * FROM " + os.Getenv("coll_users") + ";"
	var users []models.UserModel
	cnt, cerr := ur.sqlrepo.Session.SelectBySql(query).Load(&users)
	if cerr != nil {
		fmt.Println("ERROR : SQL - GetAll ", cerr)
		return []models.UserModel{}, cerr
	}
	if cnt > 0 {
		return users, nil
	}
	return []models.UserModel{}, errors.New("ERROR_NO_USERS_FOUND")
}
func (ur *UserRepoSQL) GetOne(u models.UserModel) (models.UserModel, error) {
	// Collection := ur.mongoRepo.Client.Database(ur.mongoRepo.DBInfo.Database).Collection(os.Getenv("coll_users"))
	// filter := bson.M{
	// 	"id": u.UserId,
	// }
	var user models.UserModel
	// err := Collection.FindOne(context.Background(), filter).Decode(&user)
	// if err != nil {
	// 	fmt.Println("ERROR : GetOne", err)
	// 	return models.UserModel{}, err
	// }
	return user, nil
}
