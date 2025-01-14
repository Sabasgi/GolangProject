package main

import (
	"repogin/internal/db"
	"repogin/internal/services"
	"repogin/logs"

	"fmt"
	"os"
	"repogin/internal/handlers"
	mongo "repogin/internal/repositories/mongo"
	sql "repogin/internal/repositories/sql"

	"repogin/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello in main")
	Error := godotenv.Load(".env.dev")
	if Error != nil {
		fmt.Println("ENV File Reading Error : ", Error)
		// return Error
	}
	// Reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Choose Database\n1-SQL DB\n2-MONGO DB\nEnter your choice:")
	// choice, dbError := Reader.ReadString('\n')
	// if dbError != nil {
	// 	fmt.Println("ERROR : ", dbError)
	// }
	// fmt.Println("YOUr chhoice ", choice)
	// dbchoice, _ := strconv.Atoi(strings.TrimSpace(choice))
	DataSources := createDataSource()
	inject(DataSources, 1)

}

func createDataSource() models.DataSource {
	var dataSource models.DataSource
	dataSource.MongoDSN = os.Getenv("mongodsn")
	dataSource.SQLDSN = os.Getenv("sqldsn")
	dataSource.MongoDBName = os.Getenv("mongodb")
	dataSource.SQLDBName = os.Getenv("sqldb")
	logs.Init()
	// dataSource.Collections["Products"] = os.Getenv("coll_products")
	dataSource.Port = os.Getenv("port")
	return dataSource

}

func inject(dataSource models.DataSource, dbchoice int) {
	var allServices handlers.AllServices
	switch dbchoice {
	case 1:
		fmt.Println("Case 1 SQL")
		sqlRepo := db.NewSQLRepo(
			models.DBInfo{
				DSN:      dataSource.SQLDSN,
				Name:     "SQL",
				Database: dataSource.SQLDBName,
			})
		//repos to inject into services
		prodrepo := sql.NewProductSQLRepo(sqlRepo)
		Countryrepo := sql.NewCountryRepo(sqlRepo)
		Staterepo := sql.NewStateRepo(sqlRepo)
		Cityrepo := sql.NewCityRepo(sqlRepo)
		Labrepo := sql.NewLabRepo(sqlRepo)
		Branchrepo := sql.NewBranchRepo(sqlRepo)
		Urepo := sql.NewUserrRepo(sqlRepo)
		Rolerepo := sql.NewRoleRepo(sqlRepo)
		Doctorrepo := sql.NewDoctorRepo(sqlRepo)
		Hospitalrepo := sql.NewHospitalRepo(sqlRepo)
		HospBranchrepo := sql.NewHospitalBranchRepo(sqlRepo)
		HospDocrepo := sql.NewHospitalDoctorRepo(sqlRepo)
		DeptRepo := sql.NewDepartmentRepo(sqlRepo)
		LabServRepo := sql.NewServicesRepo(sqlRepo)
		PatRepo := sql.NewPatientRepo(sqlRepo)

		//services to inject into handlers
		prodService := services.NewProdService(prodrepo)
		Countryservice := services.NewCountryService(Countryrepo)
		Stateservice := services.NewStateService(Staterepo)
		Cityservice := services.NewCityService(Cityrepo)
		Labservice := services.NewLabService(Labrepo)
		Branchservice := services.NewBranchService(Branchrepo) //lab branches
		Roleservice := services.NewRoleService(Rolerepo)
		Userservice := services.NewUserrService(Urepo)
		Docservice := services.NewDoctorService(Doctorrepo)
		Hospitalservice := services.NewHospitalService(Hospitalrepo)
		HospBrnchservice := services.NewHospBranchService(HospBranchrepo)
		HospDocservice := services.NewHospitalDoctorService(HospDocrepo)
		Deptservice := services.NewDepartmentService(DeptRepo)
		LabServ := services.NewServiceService(LabServRepo)
		PatServ := services.NewPatientService(PatRepo)

		// handlers
		allServices.Prodservice = prodService
		allServices.Countryservice = Countryservice
		allServices.Stateservice = Stateservice
		allServices.Cityservice = Cityservice
		allServices.Labservice = Labservice
		allServices.Branchservice = Branchservice
		allServices.Roleservice = Roleservice
		allServices.Uservice = Userservice
		allServices.Doctorservice = Docservice
		allServices.Hospitalservice = Hospitalservice
		allServices.HospBranchservice = HospBrnchservice
		allServices.HospDocservice = HospDocservice
		allServices.DepartmentService = Deptservice
		allServices.DepartmentService = Deptservice
		allServices.LabServiceService = LabServ
		allServices.PatientService = PatServ
	case 2:
		fmt.Println("Case 2 MONGO")
		mongoRepo := db.NewMongoDBRepo(models.DBInfo{
			DSN:      dataSource.MongoDSN,
			Database: dataSource.MongoDBName,
			Name:     "MONGO",
		})
		prodRepo := mongo.NewProductRepo(mongoRepo)
		// userRepo := mongo.NewUserRepo(mongoRepo)
		prodService := services.NewProdService(prodRepo)
		// userService := services.NewUserService(userRepo)
		allServices.Prodservice = prodService
		// allServices.Userservice = userService

	}
	// e := echo.New()
	// handlers.NewHandlers(allServices, e)
	routes := gin.Default()
	handlers.NewHandlers(allServices, routes)
	Err := routes.Run(":8768")
	if Err != nil {
		fmt.Println("ERROR in server starting - ", Err)
	}
	// e.Logger.Fatal(e.Start(":" + dataSource.Port))
}
