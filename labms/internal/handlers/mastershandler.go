package handlers

import (
	"fmt"
	"log"
	"net/http"
	models "repogin/internal/models/masters"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ph *MainHandlers) CreateCountryRoute(c *gin.Context) {
	var country models.Country
	err := c.Bind(&country)
	if err != nil {
		fmt.Println("ERROR : CreateCountryRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Countryservice.CreateCountryService(country)
	if err != nil {
		fmt.Println("ERROR : CreateCountryRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Country created successfully"})
}

func (ph *MainHandlers) UpdateCountryRoute(c *gin.Context) {
	var country models.Country
	err := c.Bind(&country)
	if err != nil {
		fmt.Println("ERROR : UpdateCountryRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Countryservice.UpdateCountryService(country)
	if err != nil {
		fmt.Println("ERROR : UpdateCountryRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Country updated successfully"})
}

func (ph *MainHandlers) GetAllCountriesRoute(c *gin.Context) {
	countries, err := ph.Countryservice.GetAllCountriesService()
	if err != nil {
		fmt.Println("ERROR : GetAllCountriesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, countries)
}

func (ph *MainHandlers) GetOneCountryRoute(c *gin.Context) {
	var country models.Country
	err := c.Bind(&country)
	if err != nil {
		fmt.Println("ERROR : GetOneCountryRoute", err)
	}
	countryDetails, err := ph.Countryservice.GetOneCountryService(country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, countryDetails)
}

func (ph *MainHandlers) DeleteCountryRoute(c *gin.Context) {
	var country models.Country
	err := c.Bind(&country)
	if err != nil {
		fmt.Println("ERROR : DeleteCountryRoute", err)
	}
	err = ph.Countryservice.DeleteOneCountryService(country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Country deleted successfully"})
}

// ......................................................
func (ph *MainHandlers) CreateStateRoute(c *gin.Context) {
	var state models.State
	err := c.Bind(&state)
	if err != nil {
		fmt.Println("ERROR : CreateStateRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Stateservice.CreateStateService(state)
	if err != nil {
		fmt.Println("ERROR : CreateStateRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "State created successfully"})
}
func (ph *MainHandlers) CreateStatesRoute(c *gin.Context) {
	var cities []models.State
	err := c.Bind(&cities)
	if err != nil {
		fmt.Println("ERROR : CreateStatesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Stateservice.CreateStatesService(cities)
	if err != nil {
		fmt.Println("ERROR : CreateStatesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "States created successfully"})
}
func (ph *MainHandlers) UpdateStateRoute(c *gin.Context) {
	var state models.State
	err := c.Bind(&state)
	if err != nil {
		fmt.Println("ERROR : UpdateStateRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Stateservice.UpdateStateService(state)
	if err != nil {
		fmt.Println("ERROR : UpdateStateRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "State updated successfully"})
}

func (ph *MainHandlers) GetAllStatesRoute(c *gin.Context) {
	states, err := ph.Stateservice.GetAllStatesService()
	if err != nil {
		fmt.Println("ERROR : GetAllStatesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, states)
}
func (ph *MainHandlers) GetAllStatesOfCountryRoute(c *gin.Context) {
	id := c.Param("id")
	countryid, e := strconv.Atoi(id)
	if e != nil {
		fmt.Println("ERROR : GetAllStatesRoute , id not valid", e)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": e.Error()})
		return
	}

	states, err := ph.Stateservice.GetAllStatesOfCountryService(countryid)
	if err != nil {
		fmt.Println("ERROR : GetAllStatesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, states)
}
func (ph *MainHandlers) GetOneStateRoute(c *gin.Context) {
	var state models.State
	err := c.Bind(&state)
	if err != nil {
		fmt.Println("ERROR : GetOneStateRoute", err)
	}
	stateDetails, err := ph.Stateservice.GetOneStateService(state)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stateDetails)
}

func (ph *MainHandlers) DeleteStateRoute(c *gin.Context) {
	var state models.State
	err := c.Bind(&state)
	if err != nil {
		fmt.Println("ERROR : DeleteStateRoute", err)
	}
	err = ph.Stateservice.DeleteOneStateService(state)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "State deleted successfully"})
}
func (ph *MainHandlers) CreateCityRoute(c *gin.Context) {
	var city models.City
	err := c.Bind(&city)
	if err != nil {
		fmt.Println("ERROR : CreateCityRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Cityservice.CreateCityService(city)
	if err != nil {
		fmt.Println("ERROR : CreateCityRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "City created successfully"})
}
func (ph *MainHandlers) CreateCitiesRoute(c *gin.Context) {
	var cities []models.City
	err := c.Bind(&cities)
	if err != nil {
		fmt.Println("ERROR : CreateCitiesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Cityservice.CreateCitiesService(cities)
	if err != nil {
		fmt.Println("ERROR : CreateCitiesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "City created successfully"})
}
func (ph *MainHandlers) UpdateCityRoute(c *gin.Context) {
	var city models.City
	err := c.Bind(&city)
	if err != nil {
		fmt.Println("ERROR : UpdateCityRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Cityservice.UpdateCityService(city)
	if err != nil {
		fmt.Println("ERROR : UpdateCityRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "City updated successfully"})
}

func (ph *MainHandlers) GetAllCitiesRoute(c *gin.Context) {
	cities, err := ph.Cityservice.GetAllCitiesService()
	if err != nil {
		fmt.Println("ERROR : GetAllCitiesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cities)
}
func (ph *MainHandlers) GetAllCitiesOfStateRoute(c *gin.Context) {
	id := c.Param("id")
	stateid, e := strconv.Atoi(id)
	if e != nil {
		fmt.Println("ERROR : GetAllStatesRoute , id not valid", e)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": e.Error()})
		return
	}

	cts, err := ph.Cityservice.GetAllCitiesOfStateService(stateid)
	if err != nil {
		fmt.Println("ERROR : GetAllStatesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cts)
}
func (ph *MainHandlers) GetOneCityRoute(c *gin.Context) {
	var city models.City
	err := c.Bind(&city)
	if err != nil {
		fmt.Println("ERROR : GetOneCityRoute", err)
	}
	cityDetails, err := ph.Cityservice.GetOneCityService(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cityDetails)
}

func (ph *MainHandlers) DeleteCityRoute(c *gin.Context) {
	var city models.City
	err := c.Bind(&city)
	if err != nil {
		fmt.Println("ERROR : DeleteCityRoute", err)
	}
	err = ph.Cityservice.DeleteOneCityService(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "City deleted successfully"})
}

// ................................................................................................................
func (mh *MainHandlers) CreateLabRoute(c *gin.Context) {
	var labDetails models.Lab
	Error := c.Bind(&labDetails)
	if Error != nil {
		fmt.Println("ERROR : CreateLabRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Labservice.CreateLabService(labDetails)
	if Err != nil {
		fmt.Println("ERROR : CreateLab", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Lab created successfully"})
	return
}
func (ph *MainHandlers) CreateLabsRoute(c *gin.Context) {
	var labs []models.Lab
	err := c.Bind(&labs)
	if err != nil {
		fmt.Println("ERROR : CreateLabsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Labservice.CreateLabsService(labs)
	if err != nil {
		fmt.Println("ERROR : CreateLabsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "labs created successfully"})
}

func (mh *MainHandlers) UpdateLabRoute(c *gin.Context) {
	var labDetails models.Lab
	Error := c.Bind(&labDetails)
	if Error != nil {
		fmt.Println("ERROR : UpdateLabRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Labservice.UpdateLabService(labDetails)
	if Err != nil {
		fmt.Println("ERROR : UpdateLab", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Lab updated successfully"})
	return
}

func (mh *MainHandlers) GetAllLabsRoute(c *gin.Context) {
	role := c.GetString("role")
	labId := c.GetString("labId")
	labs, err := mh.Labservice.GetAllLabsService(role, labId)
	if err != nil {
		fmt.Println("ERROR : GetAllLabsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, labs)
	return
}

func (mh *MainHandlers) GetOneLabRoute(c *gin.Context) {
	var lab models.Lab
	bindError := c.Bind(&lab)
	if bindError != nil {
		fmt.Println("ERROR : GetOneLabRoute", bindError)
	}
	labDetails, err := mh.Labservice.GetOneLabService(lab)
	if err != nil {
		fmt.Println("ERROR : GetOneLabRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, labDetails)
	return
}

func (mh *MainHandlers) DeleteLabRoute(c *gin.Context) {
	var lab models.Lab
	bindError := c.Bind(&lab)
	if bindError != nil {
		fmt.Println("ERROR : DeleteLabRoute", bindError)
	}
	err := mh.Labservice.DeleteOneLabService(lab)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Lab deleted successfully"})
	return
}

// ............................................................................
// CreateBranchesRoute handles the branch creation request
func (ph *MainHandlers) CreateBranchesRoute(c *gin.Context) {
	var branches []models.Branch
	err := c.Bind(&branches)
	if err != nil {
		fmt.Println("ERROR : CreateBranchesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Branchservice.CreateBranchesService(branches)
	if err != nil {
		fmt.Println("ERROR : CreateBranchesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "branches created successfully"})
}

// ............................................................................................................................................

func (mh *MainHandlers) CreateRoleRoute(c *gin.Context) {
	var roleDetails models.Role
	Error := c.Bind(&roleDetails)
	if Error != nil {
		fmt.Println("ERROR : CreateRoleRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Roleservice.CreateRoleService(roleDetails)
	if Err != nil {
		fmt.Println("ERROR : CreateRole", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Role created successfully"})
	return
}

func (mh *MainHandlers) UpdateRoleRoute(c *gin.Context) {
	var roleDetails models.Role
	Error := c.Bind(&roleDetails)
	if Error != nil {
		fmt.Println("ERROR : UpdateRoleRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Roleservice.UpdateRoleService(roleDetails)
	if Err != nil {
		fmt.Println("ERROR : UpdateRole", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Role updated successfully"})
	return
}

func (mh *MainHandlers) GetAllRolesRoute(c *gin.Context) {
	roles, err := mh.Roleservice.GetAllRolesService()
	if err != nil {
		fmt.Println("ERROR : GetAllRolesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
	return
}

// func (mh *MainHandlers) GetSuperadminAllRolesRoute(c *gin.Context) {

//		roles, err := mh.Roleservice.GetAllRolesService()
//		if err != nil {
//			fmt.Println("ERROR : GetAllRolesRoute", err)
//			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, roles)
//		return
//	}
func (mh *MainHandlers) GetOneRoleRoute(c *gin.Context) {
	var role models.Role
	bindError := c.Bind(&role)
	if bindError != nil {
		fmt.Println("ERROR : GetOneRoleRoute", bindError)
	}
	roleDetails, err := mh.Roleservice.GetOneRoleService(role)
	if err != nil {
		fmt.Println("ERROR : GetOneRoleRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roleDetails)
	return
}

func (mh *MainHandlers) DeleteRoleRoute(c *gin.Context) {
	var role models.Role
	bindError := c.Bind(&role)
	if bindError != nil {
		fmt.Println("ERROR : DeleteRoleRoute", bindError)
	}
	err := mh.Roleservice.DeleteOneRoleService(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Role deleted successfully"})
	return
}
func (mh *MainHandlers) GetMenusByRole(c *gin.Context) {
	var temp models.Role
	bindError := c.Bind(&temp)
	if bindError != nil {
		fmt.Println("ERROR : DeleteRoleRoute", bindError)
	}
	R, e := mh.Roleservice.RoleService.GetOne(temp)
	if e != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": e.Error()})
		return
	}
	log.Println("Role found ", R)
	M, err := mh.Roleservice.GetMenusByRole(R)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, M)
	return
}

// ..................................................................................................................................

func (mh *MainHandlers) CreateUserRoute(c *gin.Context) {
	var userDetails models.Userr
	Error := c.Bind(&userDetails)
	if Error != nil {
		fmt.Println("ERROR : CreateUserRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Uservice.CreateUserService(userDetails)
	if Err != nil {
		fmt.Println("ERROR : CreateUser", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
	return
}

func (mh *MainHandlers) UpdateUserRoute(c *gin.Context) {
	var userDetails models.Userr
	Error := c.Bind(&userDetails)
	if Error != nil {
		fmt.Println("ERROR : UpdateUserRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Uservice.UpdateUserService(userDetails)
	if Err != nil {
		fmt.Println("ERROR : UpdateUser", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "User updated successfully"})
	return
}

func (mh *MainHandlers) GetAllUsersRoute(c *gin.Context) {
	userRole := c.GetString("role") // Assume role is extracted from token in earlier middleware
	labId := c.GetString("labId")

	users, err := mh.Uservice.GetAllUsersService(userRole, labId)
	if err != nil {
		fmt.Println("ERROR : GetAllUsersRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
	return
}
func (mh *MainHandlers) GetAllLabsAllUsersRoute(c *gin.Context) {
	r := c.GetString("role")
	id := c.GetString("labId")
	labs, er := mh.Labservice.GetAllLabsService(r, id)
	if er != nil {
		fmt.Println("ERROR : GetAllLabsAllUsersRoute", er)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": er.Error()})
		return
	}
	users, err := mh.Uservice.GetAllLabsAllUsersService(labs)
	if err != nil {
		fmt.Println("ERROR : GetAllUsersRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
	return
}
func (mh *MainHandlers) GetOneUserRoute(c *gin.Context) {
	var user models.Userr
	bindError := c.Bind(&user)
	if bindError != nil {
		fmt.Println("ERROR : GetOneUserRoute", bindError)
	}
	userDetails, err := mh.Uservice.GetOneUserService(user)
	if err != nil {
		fmt.Println("ERROR : GetOneUserRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userDetails)
	return
}

/*
func (mh *MainHandlers) GetUserLoginRoute(c *gin.Context) {
	var user models.Userr
	bindError := c.Bind(&user)
	if bindError != nil {
		fmt.Println("ERROR : GetOneUserRoute", bindError)
	}
	userDetails, err := mh.Uservice.GetUserLoginService(user)
	if err != nil {
		fmt.Println("ERROR : GetOneUserRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	t, terr := GenerateToken(userDetails)
	if terr != nil {
		logs.Error("ERROR : LoginRoute", terr)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": terr.Error()})
		return
	}
	c.Request.Header.Set("Authorization", t)
	c.JSON(http.StatusOK, map[string]string{"mesage": "SuccessFull Login"})
	c.JSON(http.StatusOK, userDetails)
	return
}*/

func (mh *MainHandlers) DeleteUserRoute(c *gin.Context) {
	var user models.Userr
	bindError := c.Bind(&user)
	if bindError != nil {
		fmt.Println("ERROR : DeleteUserRoute", bindError)
	}
	err := mh.Uservice.DeleteOneUserService(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
	return
}

// ............................................................................................................
func (mh *MainHandlers) CreateDoctorRoute(c *gin.Context) {
	var doctorDetails models.Doctor
	Error := c.Bind(&doctorDetails)
	if Error != nil {
		fmt.Println("ERROR : CreateDoctorRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Doctorservice.CreateDoctorService(doctorDetails)
	if Err != nil {
		fmt.Println("ERROR : CreateDoctor", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Doctor created successfully"})
	return
}

// CreateDoctorsRoute handles the doctor creation request
func (ph *MainHandlers) CreateDoctorsRoute(c *gin.Context) {
	var doctors []models.Doctor
	err := c.Bind(&doctors)
	if err != nil {
		fmt.Println("ERROR : CreateDoctorsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Doctorservice.CreateDoctorsService(doctors)
	if err != nil {
		fmt.Println("ERROR : CreateDoctorsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "doctors created successfully"})
}

func (mh *MainHandlers) UpdateDoctorRoute(c *gin.Context) {
	var doctorDetails models.Doctor
	Error := c.Bind(&doctorDetails)
	if Error != nil {
		fmt.Println("ERROR : UpdateDoctorRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Doctorservice.UpdateDoctorService(doctorDetails)
	if Err != nil {
		fmt.Println("ERROR : UpdateDoctor", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Doctor updated successfully"})
	return
}

func (mh *MainHandlers) GetAllDoctorsRoute(c *gin.Context) {
	doctors, err := mh.Doctorservice.GetAllDoctorsService()
	if err != nil {
		fmt.Println("ERROR : GetAllDoctorsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doctors)
	return
}

func (mh *MainHandlers) GetOneDoctorRoute(c *gin.Context) {
	var doctor models.Doctor
	bindError := c.Bind(&doctor)
	if bindError != nil {
		fmt.Println("ERROR : GetOneDoctorRoute", bindError)
	}
	doctorDetails, err := mh.Doctorservice.GetOneDoctorService(doctor)
	if err != nil {
		fmt.Println("ERROR : GetOneDoctorRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doctorDetails)
	return
}

func (mh *MainHandlers) DeleteDoctorRoute(c *gin.Context) {
	var doctor models.Doctor
	bindError := c.Bind(&doctor)
	if bindError != nil {
		fmt.Println("ERROR : DeleteDoctorRoute", bindError)
	}
	err := mh.Doctorservice.DeleteOneDoctorService(doctor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Doctor deleted successfully"})
	return
}

// .......................................................................................................

func (mh *MainHandlers) CreateBranchRoute(c *gin.Context) {
	var branchDetails models.Branch
	Error := c.Bind(&branchDetails)
	if Error != nil {
		fmt.Println("ERROR : CreateBranchRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Branchservice.CreateBranchService(branchDetails)
	if Err != nil {
		fmt.Println("ERROR : CreateBranch", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Branch created successfully"})
	return
}

func (mh *MainHandlers) UpdateBranchRoute(c *gin.Context) {
	var branchDetails models.Branch
	Error := c.Bind(&branchDetails)
	if Error != nil {
		fmt.Println("ERROR : UpdateBranchRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Branchservice.UpdateBranchService(branchDetails)
	if Err != nil {
		fmt.Println("ERROR : UpdateBranch", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Branch updated successfully"})
	return
}

func (mh *MainHandlers) GetAllBranchesRoute(c *gin.Context) {
	branches, err := mh.Branchservice.GetAllBranchesService()
	if err != nil {
		fmt.Println("ERROR : GetAllBranchesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, branches)
	return
}

func (mh *MainHandlers) GetOneBranchRoute(c *gin.Context) {
	var branch models.Branch
	bindError := c.Bind(&branch)
	if bindError != nil {
		fmt.Println("ERROR : GetOneBranchRoute", bindError)
	}
	branchDetails, err := mh.Branchservice.GetOneBranchService(branch)
	if err != nil {
		fmt.Println("ERROR : GetOneBranchRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, branchDetails)
	return
}

func (mh *MainHandlers) DeleteBranchRoute(c *gin.Context) {
	var branch models.Branch
	bindError := c.Bind(&branch)
	if bindError != nil {
		fmt.Println("ERROR : DeleteBranchRoute", bindError)
	}
	err := mh.Branchservice.DeleteOneBranchService(branch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Branch deleted successfully"})
	return
}
func (mh *MainHandlers) GetAllLabsAllBranchesRoute(c *gin.Context) {
	r := c.GetString("role")
	id := c.GetString("labId")
	labs, er := mh.Labservice.GetAllLabsService(r, id)
	if er != nil {
		fmt.Println("ERROR : GetAllLabsAllUsersRoute", er)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": er.Error()})
		return
	}
	users, err := mh.Branchservice.GetAllLabsAllBranchesService(labs)
	if err != nil {
		fmt.Println("ERROR : GetAllUsersRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
	return

}

// ................................................................................
func (mh *MainHandlers) CreateHospitalRoute(c *gin.Context) {
	var hDetails models.Hospital
	Error := c.Bind(&hDetails)
	if Error != nil {
		fmt.Println("ERROR : CreateHospitalRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Hospitalservice.CreateHospitalService(hDetails)
	if Err != nil {
		fmt.Println("ERROR : CreateHospitalRoute", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Hospital created successfully"})
	return
}

// CreateHospitalsRoute handles the hospital creation request
func (ph *MainHandlers) CreateHospitalsRoute(c *gin.Context) {
	var hospitals []models.Hospital
	err := c.Bind(&hospitals)
	if err != nil {
		fmt.Println("ERROR : CreateHospitalsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.Hospitalservice.CreateHospitalsService(hospitals)
	if err != nil {
		fmt.Println("ERROR : CreateHospitalsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "hospitals created successfully"})
}
func (mh *MainHandlers) GetAllHsopitalRoute(c *gin.Context) {
	Hsopitals, err := mh.Hospitalservice.GetAllHospitalsService()
	if err != nil {
		fmt.Println("ERROR : GetAllHsopitalRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Hsopitals)
	return
}

func (mh *MainHandlers) GetOneHospitalRoute(c *gin.Context) {
	var h models.Hospital
	bindError := c.Bind(&h)
	if bindError != nil {
		fmt.Println("ERROR : GetOneBranchRoute", bindError)
	}
	branchDetails, err := mh.Hospitalservice.GetOneHospitalService(h)
	if err != nil {
		fmt.Println("ERROR : GetOneBranchRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, branchDetails)
	return
}

// .......................................................
func (mh *MainHandlers) CreateHospitalBranchRoute(c *gin.Context) {
	var hDetails models.HospitalBranch
	Error := c.Bind(&hDetails)
	if Error != nil {
		fmt.Println("ERROR : CreateHospitalRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.HospBranchservice.CreateHospBranchService(hDetails)
	if Err != nil {
		fmt.Println("ERROR : CreateHospitalRoute", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Hospitalbranch created successfully"})
	return
}

// CreateBranchesRoute handles the hospital branch creation request
func (ph *MainHandlers) CreateHospitalBranchesRoute(c *gin.Context) {
	var branches []models.HospitalBranch
	err := c.Bind(&branches)
	if err != nil {
		fmt.Println("ERROR: CreateBranchesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.HospBranchservice.CreateBranchesService(branches)
	if err != nil {
		fmt.Println("ERROR: CreateBranchesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "hospital branches created successfully"})
}
func (mh *MainHandlers) GetAllHsopitalBranchesRoute(c *gin.Context) {
	strId := c.Param("id")
	if strId != "" {
		id, e := strconv.Atoi(strId)
		if e != nil {
			fmt.Println("ERROR : GetAllHsopitalBranchesRoute id not found", e)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": e.Error()})
			return
		}
		Hsopitals, err := mh.HospBranchservice.GetAllHospBranchsService(id)
		if err != nil {
			fmt.Println("ERROR : GetAllHsopitalRoute", err)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, Hsopitals)
		return
	}
	fmt.Println("ERROR : GetAllHsopitalBranchesRoute id not found")
	c.JSON(http.StatusInternalServerError, map[string]string{"error": "hospital Id not found "})
	return

}

func (mh *MainHandlers) GetOneHospitalBranchRoute(c *gin.Context) {
	var h models.HospitalBranch
	bindError := c.Bind(&h)
	if bindError != nil {
		fmt.Println("ERROR : GetOneHospitalBranchRoute", bindError)
	}
	branchDetails, err := mh.HospBranchservice.GetOneHospBranchService(h)
	if err != nil {
		fmt.Println("ERROR : GetOneHospitalBranchRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, branchDetails)
	return
}
func (mh *MainHandlers) UpdateHospitalBranchRoute(c *gin.Context) {
	var branchDetails models.HospitalBranch
	Error := c.Bind(&branchDetails)
	if Error != nil {
		fmt.Println("ERROR : UpdateBranchRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.HospBranchservice.UpdateHospBranchService(branchDetails)
	// mh.HospBranchservice.UpdateHospBranchService()(branchDetails)
	if Err != nil {
		fmt.Println("ERROR : UpdateBranch", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Branch updated successfully"})
	return
}

// ........................................................................
func (ph *MainHandlers) CreateHospitalDoctorsRoute(c *gin.Context) {
	var hospitalDoctors []models.HospitalDoctor
	err := c.Bind(&hospitalDoctors)
	if err != nil {
		fmt.Println("ERROR: CreateHospitalDoctorsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.HospDocservice.CreateHospitalDoctorService(hospitalDoctors)
	if err != nil {
		fmt.Println("ERROR: CreateHospitalDoctorsRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "hospital-doctor relationships created successfully"})
}

// GetAllHospitalDoctorsRoute handles the route to get all -doctors of a hospital relationships
func (ph *MainHandlers) GetAllHospitalDoctorsRoute(c *gin.Context) {
	strId := c.Param("id")
	if strId != "" {
		id, e := strconv.Atoi(strId)
		if e != nil {
			fmt.Println("ERROR : GetAllHsopitalBranchesRoute id not found", e)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": e.Error()})
			return
		}
		hospitalDoctors, err := ph.HospDocservice.GetAllHospitalDoctorsService(id)
		if err != nil {
			fmt.Println("ERROR: GetAllHospitalDoctorsRoute", err)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, hospitalDoctors)
		return
	}
	fmt.Println("ERROR : GetAllHsopitalDOcsRoute id not found")
	c.JSON(http.StatusInternalServerError, map[string]string{"error": "hospital Id not found "})
	return
}

// GetOneHospitalDoctorRoute handles the route to get a single hospital-doctor relationship by its ID
func (ph *MainHandlers) GetOneHospitalDoctorRoute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("ERROR: GetOneHospitalDoctorRoute", err)
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
		return
	}

	hospitalDoctor, err := ph.HospDocservice.GetOneHospitalDoctorService(id)
	if err != nil {
		fmt.Println("ERROR: GetOneHospitalDoctorRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hospitalDoctor)
}
