package handlers

import (
	"repogin/internal/middleware"
	"repogin/internal/services"

	"github.com/gin-gonic/gin"
)

// struct
type AllServices struct {
	Prodservice     *services.ProdServiceStruct
	Countryservice  *services.CountryServiceStruct
	Stateservice    *services.StateServiceStruct
	Cityservice     *services.CityServiceStruct
	Labservice      *services.LabServiceStruct
	Roleservice     *services.RoleServiceStruct
	Doctorservice   *services.DoctorServiceStruct
	Branchservice   *services.BranchServiceStruct // lab branches relationship
	Uservice        *services.UserrServiceStruct
	Hospitalservice *services.HospitalServiceStruct
	//relationships
	HospBranchservice *services.HospBranchServiceStruct
	HospDocservice    *services.HospitalDoctorService
	VisitTestService  *services.VisitTestService
	Visitservice      *services.VisitService
	DepartmentService *services.DepartmentService
	LabServiceService *services.ServiceService
	PatientService    *services.PatientService

	// Testservice       *services.TestService
}

// handler to handle all the services structs
type MainHandlers struct {
	Ps                *services.ProdServiceStruct
	Countryservice    *services.CountryServiceStruct
	Stateservice      *services.StateServiceStruct
	Cityservice       *services.CityServiceStruct
	Labservice        *services.LabServiceStruct
	Branchservice     *services.BranchServiceStruct
	Roleservice       *services.RoleServiceStruct
	Doctorservice     *services.DoctorServiceStruct
	Uservice          *services.UserrServiceStruct
	Hospitalservice   *services.HospitalServiceStruct
	HospBranchservice *services.HospBranchServiceStruct
	HospDocservice    *services.HospitalDoctorService
	VisitTestService  *services.VisitTestService
	Visitservice      *services.VisitService
	DepartmentService *services.DepartmentService
	LabServiceService *services.ServiceService
	PatientService    *services.PatientService

	// Testservice       *services.TestService

	//you can write all the structs here
}

func NewHandlers(s AllServices, e *gin.Engine) {
	o := e.Group("/o")
	c := e.Group("/c")
	c.Use(middleware.TokenValidationMiddleware()) // Middleware to validate token
	// r := e.Group("/r")
	// adminRoutes := e.Group("/admin")
	// adminRoutes.Use(middleware.RoleBasedMiddleware("admin"))

	// clientRoutes := e.Group("/client")
	// clientRoutes.Use(middleware.RoleBasedMiddleware("client"))

	h := &MainHandlers{
		Ps:                s.Prodservice,
		Countryservice:    s.Countryservice,
		Stateservice:      s.Stateservice,
		Cityservice:       s.Cityservice,
		Labservice:        s.Labservice,
		Branchservice:     s.Branchservice,
		Uservice:          s.Uservice,
		Roleservice:       s.Roleservice,
		Doctorservice:     s.Doctorservice,
		Hospitalservice:   s.Hospitalservice,
		HospBranchservice: s.HospBranchservice,
		HospDocservice:    s.HospDocservice,
		VisitTestService:  s.VisitTestService,
		Visitservice:      s.Visitservice,
		DepartmentService: s.DepartmentService,
		LabServiceService: s.LabServiceService,
		PatientService:    s.PatientService,
		// Testservice:       s.Testservice,
	}
	// c.Use(middleware.BaicAuth())
	//product routes registered here
	c.POST("/product/create", h.CreateProductRoute)
	o.PUT("/product/update", h.UpdateProductRoute)
	o.GET("/product/getall", h.GetAllProductsRoute)
	o.POST("/product/getone", h.GetOneProductRoute)
	o.POST("/product/deleteone", h.DeleteOneRoute)
	//users routes registered here
	// adminRoutes.POST("/user/create", h.CreateUserRoute)
	c.POST("/user/create", middleware.RoleAuthorization([]string{"admin"}), h.CreateUserRoute) /// usage of rolebased authorization
	c.GET("/user/getall", middleware.RoleAuthorization([]string{"admin"}), h.GetAllUsersRoute)
	c.GET("/labs/users/getall", middleware.RoleAuthorization([]string{"admin"}), h.GetAllLabsAllUsersRoute)
	o.POST("/user/create", h.CreateUserRoute)
	o.PUT("/user/update", h.UpdateUserRoute)
	// o.GET("/user/getall", h.GetAllUsersRoute)
	o.GET("/user/getone", h.GetOneUserRoute)
	o.POST("/user/deleteone", h.DeleteUserRoute)
	//role
	c.POST("/role/create", middleware.RoleAuthorization([]string{"admin", "superadmin"}), h.CreateRoleRoute)
	o.PUT("/role/update", h.UpdateRoleRoute)
	o.GET("/role/getall", h.GetAllRolesRoute)
	o.GET("/role/getone", h.GetOneRoleRoute)
	o.POST("/role/deleteone", h.DeleteRoleRoute)
	// //login
	o.POST("/user/login", h.LoginRoute)
	o.POST("/role/menus", h.GetMenusByRole)
	//country
	o.POST("/country/create", h.CreateCountryRoute)
	o.PUT("/country/update", h.UpdateCountryRoute)
	o.GET("/country/getall", h.GetAllCountriesRoute)
	o.POST("/country/getone", h.GetOneCountryRoute)
	o.POST("/country/deleteone", h.DeleteCountryRoute)
	//State
	o.POST("/state/create", h.CreateStateRoute)
	o.POST("/states/create", h.CreateStatesRoute)
	o.PUT("/state/update", h.UpdateStateRoute)
	o.GET("/state/getall", h.GetAllStatesRoute)
	o.GET("/country/states/:id", h.GetAllStatesOfCountryRoute)
	o.POST("/state/getone", h.GetOneStateRoute)
	o.POST("/state/deleteone", h.DeleteStateRoute)
	//City
	o.POST("/city/create", h.CreateCityRoute)
	o.POST("/cities/create", h.CreateCitiesRoute)
	o.PUT("/city/update", h.UpdateCityRoute)
	o.GET("/city/getall", h.GetAllCitiesRoute)
	o.POST("/city/getone", h.GetOneCityRoute)
	o.POST("/city/deleteone", h.DeleteCityRoute)
	o.GET("/state/cities/:id", h.GetAllCitiesOfStateRoute)

	//Lab
	o.POST("/lab/create", h.CreateLabRoute)
	o.PUT("/lab/update", h.UpdateLabRoute)
	c.GET("/lab/getall", h.GetAllLabsRoute)
	c.GET("/labs/branches/getall", middleware.RoleAuthorization([]string{"admin"}), h.GetAllLabsAllBranchesRoute)
	c.POST("/lab/getone", h.GetOneLabRoute)
	o.POST("/lab/deleteone", h.DeleteLabRoute)
	//Branches of lab
	o.POST("/branch/create", h.CreateBranchRoute)
	o.POST("/branches/create", h.CreateBranchesRoute)
	o.PUT("/branch/update", h.UpdateBranchRoute)
	c.GET("/branch/getall", h.GetAllBranchesRoute)
	o.POST("/branch/getone", h.GetOneBranchRoute)
	o.POST("/branch/deleteone", h.DeleteBranchRoute)
	c.GET("/branches/depts/getall", middleware.RoleAuthorization([]string{"admin"}), h.GetAllBranchessAllDeptRoute)

	//Branches of department
	// o.POST("/dept/create", h.CreateDepartmentRoute)
	c.POST("/dept/create", h.CreateDepartmentRoute)
	o.PUT("/dept/update", h.UpdateDepartmentRoute)
	o.GET("/dept/getall", h.GetAllDepartmentsRoute)
	o.POST("/dept/getone", h.GetAllDepartmentsRoute)
	o.POST("/dept/deleteone", h.DeleteDepartmentRoute)
	//services of branches oflabs
	o.POST("/lab/bulk/services/create", h.CreateBulkServicesRoute)
	o.POST("/lab/service/create", h.CreateServiceRoute)
	o.PUT("/lab/service/update", h.UpdateServiceRoute)
	o.GET("/lab/service/getall", h.GetAllServicesRoute)
	o.POST("/lab/service/getone", h.GetOneServiceRoute)
	o.POST("/lab/service/deleteone", h.DeleteServiceRoute)
	//patient
	o.POST("/patient/bulk/create", h.CreateBulkPatientsRoute)
	o.POST("/patient/create", h.CreatePatientRoute)
	o.PUT("/patient/update", h.UpdatePatientRoute)
	o.GET("/patient/getall", h.GetAllPatientsRoute)
	o.POST("/patient/getone", h.GetOnePatintRoute)
	o.POST("/patient/deleteone", h.DeletePatientRoute)
	/*// Doctor
	o.POST("/doctor/create", h.CreateDoctorRoute)
	o.POST("/doctors/create", h.CreateDoctorsRoute)
	o.PUT("/doctor/update", h.UpdateDoctorRoute)
	o.GET("/doctor/getall", h.GetAllDoctorsRoute)
	o.POST("/doctor/getone", h.GetOneDoctorRoute)
	o.POST("/doctor/deleteone", h.DeleteDoctorRoute)
	// Hospital
	o.GET("/hospital/getall", h.GetAllHsopitalRoute)
	o.POST("/hospital/create", h.CreateHospitalRoute)
	o.POST("/hospitals/create", h.CreateHospitalsRoute)
	o.POST("/hospital/getone", h.GetOneHospitalRoute)
	// o.PUT("/hospital/update", h.UpdateBranchRou)
	// o.POST("/hospital/deleteone", h.DeleteDoctorRoute)
	// Hospital Branch
	o.POST("/hosp/branch/create", h.CreateHospitalBranchRoute)
	o.POST("/hosp/branches/create", h.CreateHospitalBranchesRoute)  //cretae in bullk
	o.GET("/hosp/branch/getall/:id", h.GetAllHsopitalBranchesRoute) // get all the branches of particular hospital
	o.POST("/hosp/branch/getone", h.GetOneHospitalRoute)
	// o.PUT("/hospital/update", h.UpdateBranchRou)
	// o.POST("/hospital/deleteone", h.DeleteDoctorRoute)
	// Hospital doctor
	// o.POST("/hosp/doctor/create", h.CreateBranchesRoute)
	o.POST("/hosp/doctors/create", h.CreateHospitalDoctorsRoute)   //cretae in bulk
	o.GET("/hosp/doctor/getall/:id", h.GetAllHospitalDoctorsRoute) // get all the doctors of particular hospital
	o.POST("/hosp/doctor/getone", h.GetOneHospitalDoctorRoute)     // get hospital_doctor by hd_id
	// o.PUT("/hospital/update", h.UpdateBranchRou)
	// o.POST("/hospital/deleteone", h.DeleteDoctorRoute)
	//visits
	o.GET("/visit/getall", h.GetAllVisitsRoute)
	o.POST("/visit/create", h.CreateVisitRoute)
	//visit test
	o.GET("/visittest/getall", h.GetAllVisitTestsRoute)
	o.POST("/visittest/create", h.CreateVisitTestsRoute)
	//test
	o.GET("/test/getall", h.GetAllTestsRoute)
	o.POST("/test/create", h.CreateTestRoute)
	*/
}
