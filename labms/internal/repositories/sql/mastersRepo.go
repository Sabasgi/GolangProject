package repositories

import (
	"errors"
	"fmt"
	"log"
	"os"
	"repogin/internal/db"
	kfka "repogin/internal/queues/kafka"
	"strconv"
	"strings"
	"time"

	"crypto/rand"
	"math/big"

	models "repogin/internal/models/masters"

	"github.com/gocraft/dbr"
)

//Masters
// Country , State , City , User , Role , Lab , Branch

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numeric = "0123456789"

func generateRandomString(length int) (string, error) {
	result := make([]byte, length)

	// Generate first two alphabets
	for i := 0; i < 2; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", err
		}
		result[i] = alphabet[num.Int64()]
	}

	// Generate next four characters (alphanumeric)
	for i := 2; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(numeric)))) // 26 letters + 10 digits = 36 characters
		if err != nil {
			return "", err
		}

		// if num.Int64() < 26 {
		// If number is less than 26, it's an uppercase letter
		result[i] = numeric[num.Int64()]
		// } else {
		// If number is 26 or above, it's a digit (convert 26 -> '0', 27 -> '1', ..., 35 -> '9')
		// result[i] = byte('0' + (num.Int64() - 26))
		// }
	}

	return string(result), nil
}

type ContryRepo interface {
	Create(models.Country) error
	Modify(models.Country) error
	GetAll() ([]models.Country, error)
	GetOne(models.Country) (models.Country, error)
	DeleteOne(models.Country) error
}
type CountrySQLRepo struct {
	CRepo *db.SQLRepo
}

func NewCountryRepo(sr *db.SQLRepo) *CountrySQLRepo {
	return &CountrySQLRepo{
		CRepo: sr,
	}
}
func (cr *CountrySQLRepo) Create(c models.Country) error {
	codeLength, _ := strconv.Atoi(os.Getenv("codelength")) // Fetch code length from environment variable
	code, err := generateRandomString(codeLength)          // Generate random string
	if err != nil {
		fmt.Println("ERROR: HospitalSQLRepo Create - Code generation error", err)
		return errors.New("ERROR_CODE_GENERATION_ERR")
	}
	c.Code = code // Set generated hospital code
	// c.Id = "1"
	query := "INSERT INTO Country (country_name, country_code,gst_percentage) VALUES (?, ?,?)"
	res, rerr := cr.CRepo.Session.Exec(query, c.Name, c.Code, c.GST_Percentage)
	if rerr != nil {
		fmt.Println("ERROR : ProductSQLRepo Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("product Inserted / updated successfully ! with sql id - ", inc)
		return nil
	}
	return errors.New("ERROR_PRODUCT_NOT_INSERTED_UPDATED")
}
func (cr *CountrySQLRepo) Modify(c models.Country) error {
	query := "UPDATE Country SET country_name = ?,gst_percentage = ?,  country_code = ? WHERE country_id = ?"
	res, rerr := cr.CRepo.Session.Exec(query, c.Name, c.GST_Percentage, c.Code, c.ID)
	if rerr != nil {
		fmt.Println("ERROR : CountrySQLRepo Modify ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Country updated successfully!")
		return nil
	}
	return errors.New("ERROR_COUNTRY_NOT_UPDATED")
}
func (cr *CountrySQLRepo) GetAll() ([]models.Country, error) {
	var countries []models.Country
	query := "SELECT country_id, country_name, country_code FROM Country"
	rowsCount, rerr := cr.CRepo.Session.SelectBySql(query).Load(&countries)
	if rerr != nil {
		fmt.Println("ERROR : CountrySQLRepo GetAll ", rerr)
		return nil, rerr
	}

	if rowsCount > 0 && len(countries) > 0 {
		fmt.Println("Suuccesss : CountrySQLRepo Len(countries) ", len(countries) > 0)
		return countries, nil
	}
	fmt.Println("ERROR : CountrySQLRepo GetAll ", "not found")
	return countries, errors.New("Not Found")
}
func (cr *CountrySQLRepo) GetOne(c models.Country) (models.Country, error) {
	query := "SELECT country_id, country_name, country_code FROM Country WHERE country_id = ?"
	var country models.Country
	err := cr.CRepo.Session.SelectBySql(query).LoadOne(&country)
	if err != nil {
		fmt.Println("ERROR : CountrySQLRepo GetOne ", err)
		return country, err
	}

	return country, nil
}
func (cr *CountrySQLRepo) DeleteOne(c models.Country) error {
	query := "DELETE FROM Country WHERE country_id = ?"
	res, rerr := cr.CRepo.Session.Exec(query, c.ID)
	if rerr != nil {
		fmt.Println("ERROR : CountrySQLRepo DeleteOne ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Country deleted successfully!")
		return nil
	}
	return errors.New("ERROR_COUNTRY_NOT_DELETED")
}

// ..............................................................................................................................................................
type CityRepo interface {
	Create(models.City) error
	Modify(models.City) error
	GetAll() ([]models.City, error)
	GetOne(models.City) (models.City, error)
	DeleteOne(models.City) error
	BulkCreate(cities []models.City) error
	GetAllCitiessOfState(stateid int) ([]models.City, error)
}
type CitySQLRepo struct {
	CRepo *db.SQLRepo
}

func NewCityRepo(sr *db.SQLRepo) *CitySQLRepo {
	return &CitySQLRepo{
		CRepo: sr,
	}
}

// Create inserts a new city into the City table
func (cr *CitySQLRepo) Create(c models.City) error {
	query := "INSERT INTO City (city_name, state_id) VALUES (?, ?)"
	res, rerr := cr.CRepo.Session.Exec(query, c.CityName, c.StateID)
	if rerr != nil {
		fmt.Println("ERROR : CitySQLRepo Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("City inserted successfully! with sql id - ", inc)
		return nil
	}
	return errors.New("ERROR_CITY_NOT_INSERTED")
}

// Create inserts a array of city into the City table
func (cr *CitySQLRepo) BulkCreate(cities []models.City) error {
	if len(cities) > 0 {
		query := "INSERT INTO City (city_name, state_id) VALUES (?, ?)"
		for _, c := range cities {
			res, rerr := cr.CRepo.Session.Exec(query, c.CityName, c.StateID)
			if rerr != nil {
				fmt.Println("ERROR : CitySQLRepo Create ", rerr)
				return rerr
			}
			inc, _ := res.LastInsertId()
			rac, _ := res.RowsAffected()
			if inc > 0 || rac > 0 {
				fmt.Println("City inserted successfully! with sql id - ", inc)
			} else {
				fmt.Println("City failed to insert - ", c.CityName, c.StateID)
				//LOGS Implementation remaining
			}
		}
	}
	return errors.New("ERROR_CITIES_LENGTH_IS_ZERO")
}

// Modify updates an existing city in the City table
func (cr *CitySQLRepo) Modify(c models.City) error {
	query := "UPDATE City SET city_name = ?, state_id = ? WHERE city_id = ?"
	res, rerr := cr.CRepo.Session.Exec(query, c.CityName, c.StateID, c.CityID)
	if rerr != nil {
		fmt.Println("ERROR : CitySQLRepo Modify ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("City updated successfully!")
		return nil
	}
	return errors.New("ERROR_CITY_NOT_UPDATED")
}

// GetAll retrieves all cities from the City table
func (cr *CitySQLRepo) GetAll() ([]models.City, error) {
	var cities []models.City
	query := "SELECT city_id, city_name, state_id FROM City"
	rowsCount, rerr := cr.CRepo.Session.SelectBySql(query).Load(&cities)
	if rerr != nil {
		fmt.Println("ERROR : CitySQLRepo GetAll ", rerr)
		return nil, rerr
	}

	if rowsCount > 0 && len(cities) > 0 {
		fmt.Println("SUCCESS : CitySQLRepo Len(cities) ", len(cities) > 0)
		return cities, nil
	}
	fmt.Println("ERROR : CitySQLRepo GetAll ", "not found")
	return cities, errors.New("not Found")
}

// GetAll retrieves all cities from the city table
func (cr *CitySQLRepo) GetAllCitiessOfState(stateid int) ([]models.City, error) {
	var cts []models.City
	query := "SELECT city_id, city_name, state_id FROM City WHERE state_id = ?"
	rowsCount, rerr := cr.CRepo.Session.SelectBySql(query, stateid).Load(&cts)
	if rerr != nil {
		fmt.Println("ERROR : GetAllCitiessOfState GetAll ", rerr)
		return nil, rerr
	}

	if rowsCount > 0 && len(cts) > 0 {
		fmt.Println("SUCCESS : GetAllCitiessOfState Len(cities) ", len(cts) > 0)
		return cts, nil
	}
	fmt.Println("ERROR : GetAllCitiessOfState GetAll ", "not found")
	return cts, errors.New("Not Found")
}

// GetOne retrieves a specific city by its ID
func (cr *CitySQLRepo) GetOne(c models.City) (models.City, error) {
	query := "SELECT city_id, city_name, state_id FROM City WHERE city_id = ?"
	var city models.City
	err := cr.CRepo.Session.SelectBySql(query, c.CityID).LoadOne(&city)
	if err != nil {
		fmt.Println("ERROR : CitySQLRepo GetOne ", err)
		return city, err
	}

	return city, nil
}

// DeleteOne deletes a city by its ID from the City table
func (cr *CitySQLRepo) DeleteOne(c models.City) error {
	query := "DELETE FROM City WHERE city_id = ?"
	res, rerr := cr.CRepo.Session.Exec(query, c.CityID)
	if rerr != nil {
		fmt.Println("ERROR : CitySQLRepo DeleteOne ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("City deleted successfully!")
		return nil
	}
	return errors.New("ERROR_CITY_NOT_DELETED")
}

// ......................................................................................................................................................................
type StateRepo interface {
	Create(models.State) error
	Modify(models.State) error
	GetAll() ([]models.State, error)
	BulkCreate([]models.State) error
	GetOne(models.State) (models.State, error)
	DeleteOne(models.State) error
	GetAllStatesOfCountry(int) ([]models.State, error)
}
type StateSQLRepo struct {
	CRepo *db.SQLRepo
}

func NewStateRepo(sr *db.SQLRepo) *StateSQLRepo {
	return &StateSQLRepo{
		CRepo: sr,
	}
}

// Create inserts a new state into the State table
func (sr *StateSQLRepo) Create(s models.State) error {
	query := "INSERT INTO State (state_name, country_id) VALUES (?, ?)"
	res, rerr := sr.CRepo.Session.Exec(query, s.StateName, s.CountryID)
	if rerr != nil {
		fmt.Println("ERROR : StateSQLRepo Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("State inserted successfully! with sql id - ", inc)
		return nil
	}
	return errors.New("ERROR_STATE_NOT_INSERTED")
}

// BulkCreate inserts an array of states into the State table
func (sr *StateSQLRepo) BulkCreate(states []models.State) error {
	if len(states) > 0 {
		query := "INSERT INTO State (state_name, country_id) VALUES (?, ?)"
		for _, s := range states {
			res, rerr := sr.CRepo.Session.Exec(query, s.StateName, s.CountryID)
			if rerr != nil {
				fmt.Println("ERROR : StateSQLRepo Create ", rerr)
				return rerr
			}
			inc, _ := res.LastInsertId()
			rac, _ := res.RowsAffected()
			if inc > 0 || rac > 0 {
				fmt.Println("State inserted successfully! with sql id - ", inc)
			} else {
				fmt.Println("State failed to insert - ", s.StateName, s.CountryID)
				//LOGS Implementation remaining
			}
		}
		return nil // Successful insertion, so we return nil
	}
	return errors.New("ERROR_STATES_LENGTH_IS_ZERO")
}

// Modify updates an existing state in the State table
func (sr *StateSQLRepo) Modify(s models.State) error {
	query := "UPDATE State SET state_name = ?, country_id = ? WHERE state_id = ?"
	res, rerr := sr.CRepo.Session.Exec(query, s.StateName, s.CountryID, s.StateID)
	if rerr != nil {
		fmt.Println("ERROR : StateSQLRepo Modify ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("State updated successfully!")
		return nil
	}
	return errors.New("ERROR_STATE_NOT_UPDATED")
}

// GetAll retrieves all states from the State table
func (sr *StateSQLRepo) GetAll() ([]models.State, error) {
	var states []models.State
	query := "SELECT state_id, state_name, country_id FROM State"
	rowsCount, rerr := sr.CRepo.Session.SelectBySql(query).Load(&states)
	if rerr != nil {
		fmt.Println("ERROR : StateSQLRepo GetAll ", rerr)
		return nil, rerr
	}

	if rowsCount > 0 && len(states) > 0 {
		fmt.Println("SUCCESS : StateSQLRepo Len(states) ", len(states) > 0)
		return states, nil
	}
	fmt.Println("ERROR : StateSQLRepo GetAll ", "not found")
	return states, errors.New("Not Found")
}

// GetAll retrieves all states from the State table
func (sr *StateSQLRepo) GetAllStatesOfCountry(country_id int) ([]models.State, error) {
	var states []models.State
	query := "SELECT state_id, state_name, country_id FROM State WHERE country_id = ?"
	rowsCount, rerr := sr.CRepo.Session.SelectBySql(query, country_id).Load(&states)
	if rerr != nil {
		fmt.Println("ERROR : StateSQLRepo GetAll ", rerr)
		return nil, rerr
	}

	if rowsCount > 0 && len(states) > 0 {
		fmt.Println("SUCCESS : StateSQLRepo Len(states) ", len(states) > 0)
		return states, nil
	}
	fmt.Println("ERROR : StateSQLRepo GetAll ", "not found")
	return states, errors.New("Not Found")
}

// GetOne retrieves a specific state by its ID
func (sr *StateSQLRepo) GetOne(s models.State) (models.State, error) {
	query := "SELECT state_id, state_name, country_id FROM State WHERE state_id = ?"
	var state models.State
	err := sr.CRepo.Session.SelectBySql(query, s.StateID).LoadOne(&state)
	if err != nil {
		fmt.Println("ERROR : StateSQLRepo GetOne ", err)
		return state, err
	}

	return state, nil
}

// DeleteOne deletes a state by its ID from the State table
func (sr *StateSQLRepo) DeleteOne(s models.State) error {
	query := "DELETE FROM State WHERE state_id = ?"
	res, rerr := sr.CRepo.Session.Exec(query, s.StateID)
	if rerr != nil {
		fmt.Println("ERROR : StateSQLRepo DeleteOne ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("State deleted successfully!")
		return nil
	}
	return errors.New("ERROR_STATE_NOT_DELETED")
}

// .........................................................................................................................................................
type RoleRepo interface {
	Create(models.Role) error
	Modify(models.Role) error
	GetAll() ([]models.Role, error)
	GetOne(models.Role) (models.Role, error)
	DeleteOne(models.Role) error
	FetchMenusByRole(int) ([]models.Menu, error)
	// AddMenusForRole([]models.Menu) error
}
type RoleSQLRepo struct {
	CRepo *db.SQLRepo
}

func NewRoleRepo(sr *db.SQLRepo) *RoleSQLRepo {
	return &RoleSQLRepo{
		CRepo: sr,
	}
}

// Create inserts a new role into the Role table
func (rr *RoleSQLRepo) Create(r models.Role) error {
	query := "INSERT INTO Role (role_name, description) VALUES (?, ?)"
	res, rerr := rr.CRepo.Session.Exec(query, r.RoleName, r.Description)
	if rerr != nil {
		fmt.Println("ERROR : RoleSQLRepo Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("Role inserted successfully! with sql id - ", inc)
		return nil
	}
	return errors.New("ERROR_ROLE_NOT_INSERTED")
}

// Create inserts a new menus into the mennu table for role
func (rr *RoleSQLRepo) AddMenusForRole(r models.Role) error {
	r.RoleName = strings.ToLower(r.RoleName)
	// `insert into menu ("menu_id", "label", "to_url", "icon", "parent_menu_id") values('1','Home','','',NULL);`
	query := "INSERT INTO Role (role_name, description) VALUES (?, ?)"
	res, rerr := rr.CRepo.Session.Exec(query, r.RoleName, r.Description)
	if rerr != nil {
		fmt.Println("ERROR : RoleSQLRepo Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("Role inserted successfully! with sql id - ", inc)
		return nil
	}
	return errors.New("ERROR_ROLE_NOT_INSERTED")
}

// Modify updates an existing role in the Role table
func (rr *RoleSQLRepo) Modify(r models.Role) error {
	query := "UPDATE Role SET role_name = ?, description = ? WHERE role_id = ?"
	res, rerr := rr.CRepo.Session.Exec(query, r.RoleName, r.Description, r.RoleID)
	if rerr != nil {
		fmt.Println("ERROR : RoleSQLRepo Modify ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Role updated successfully!")
		return nil
	}
	return errors.New("ERROR_ROLE_NOT_UPDATED")
}

// GetAll retrieves all roles from the Role table
func (rr *RoleSQLRepo) GetAll() ([]models.Role, error) {
	var roles []models.Role
	query := "SELECT role_id, role_name, description FROM Role where role_id>1"
	rowsCount, rerr := rr.CRepo.Session.SelectBySql(query).Load(&roles)
	if rerr != nil {
		fmt.Println("ERROR : RoleSQLRepo GetAll ", rerr)
		return nil, rerr
	}

	if rowsCount > 0 && len(roles) > 0 {
		fmt.Println("SUCCESS : RoleSQLRepo Len(roles) ", len(roles))
		return roles, nil
	}
	fmt.Println("ERROR : RoleSQLRepo GetAll ", "not found")
	return roles, errors.New("Not Found")
}

// GetOne retrieves a specific role by its ID
func (rr *RoleSQLRepo) GetOne(r models.Role) (models.Role, error) {
	var query string
	var st *dbr.SelectStmt

	if r.RoleID != 0 {
		query = "SELECT role_id, role_name, description FROM Role WHERE role_id = ?"
		st = rr.CRepo.Session.SelectBySql(query, r.RoleID)
	} else {
		if r.RoleName != "" {
			query = "SELECT role_id, role_name, description FROM Role WHERE role_name = ?"
			st = rr.CRepo.Session.SelectBySql(query, r.RoleName)
		}
	}
	var role models.Role
	err := st.LoadOne(&role)
	if err != nil {
		fmt.Println("ERROR : RoleSQLRepo GetOne ", err)
		return role, err
	}

	return role, nil
}

// DeleteOne deletes a role by its ID from the Role table
func (rr *RoleSQLRepo) DeleteOne(r models.Role) error {
	query := "DELETE FROM Role WHERE role_id = ?"
	res, rerr := rr.CRepo.Session.Exec(query, r.RoleID)
	if rerr != nil {
		fmt.Println("ERROR : RoleSQLRepo DeleteOne ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Role deleted successfully!")
		return nil
	}
	return errors.New("ERROR_ROLE_NOT_DELETED")
}

// FetchMenusByRole fetches all menus a role has permission to access
func (rr *RoleSQLRepo) FetchMenusByRole(roleID int) ([]models.Menu, error) {
	query := `
	SELECT m.menu_id, m.label, m.to_url, m.icon, m.parent_menu_id,p.allowed
	FROM menu m
	INNER JOIN permission p ON m.menu_id = p.menu_id
	WHERE p.role_id = ? AND p.allowed = TRUE`
	var menus []models.Menu
	res, rerr := rr.CRepo.Session.SelectBySql(query, roleID).Load(&menus)
	if rerr != nil {
		log.Println("ERRPR: FetchMenusByRole", rerr)
	}
	if len(menus) > 0 {
		log.Println("Count of records ", res)
		return menus, nil
	}
	log.Println("Count of records ", res)
	return nil, errors.New("ERROR_NO_MENU_FOUND")
}

// ..................................................................................................................................................
type LabRepo interface {
	Create(models.Lab) error
	BulkCreate([]models.Lab) error
	Modify(models.Lab) error
	GetAll(role string, labId int) ([]models.Lab, error)
	// GetAllLabsDeptsBranchesServices() ([]models.Lab, error)
	GetOne(models.Lab) (models.Lab, error)
	DeleteOne(models.Lab) error
}
type LabSQLRepo struct {
	CRepo *db.SQLRepo
}

func NewLabRepo(sr *db.SQLRepo) *LabSQLRepo {
	return &LabSQLRepo{
		CRepo: sr,
	}
}

// Create a new lab
func (lr *LabSQLRepo) Create(l models.Lab) error {
	ct := time.Now()
	layout := "2006-01-02 15:04:05"
	l.CreatedOn = ct.Format(layout)
	l.CreatedBy = ""
	codeLength, _ := strconv.Atoi(os.Getenv("codelength")) // Fetch code length from environment variable
	code, err := generateRandomString(codeLength)          // Generate random string
	if err != nil {
		fmt.Println("ERROR: HospitalSQLRepo Create - Code generation error", err)
		return errors.New("ERROR_CODE_GENERATION_ERR")
	}
	l.LabCode = code // Set generated hospital code

	query := "INSERT INTO Lab (lab_name, lab_code,created_on,created_by) VALUES (?, ?,?,?)"
	res, rerr := lr.CRepo.Session.Exec(query, l.LabName, l.LabCode, l.CreatedOn, l.CreatedBy)
	if rerr != nil {
		fmt.Println("ERROR : LabSQLRepo Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("Lab inserted successfully! with sql id - ", inc)
		return nil
	}
	return errors.New("ERROR_LAB_NOT_INSERTED")
}

// BulkCreate inserts an array of labs into the Lab table
func (lr *LabSQLRepo) BulkCreate(labs []models.Lab) error {
	if len(labs) > 0 {
		query := "INSERT INTO Lab (lab_name, lab_code) VALUES (?, ?)"
		for _, lab := range labs {
			res, rerr := lr.CRepo.Session.Exec(query, lab.LabName, lab.LabCode)
			if rerr != nil {
				fmt.Println("ERROR : LabSQLRepo Create ", rerr)
				return rerr
			}
			inc, _ := res.LastInsertId()
			rac, _ := res.RowsAffected()
			if inc > 0 || rac > 0 {
				fmt.Println("Lab inserted successfully! with sql id - ", inc)
			} else {
				fmt.Println("Lab failed to insert - ", lab.LabName, lab.LabCode)
				//LOGS Implementation remaining
			}
		}
		return nil // Successful insertion, so return nil
	}
	return errors.New("ERROR_LABS_LENGTH_IS_ZERO")
}

// Modify an existing lab
func (lr *LabSQLRepo) Modify(l models.Lab) error {
	query := "UPDATE Lab SET lab_name = ?, lab_code = ? WHERE lab_id = ?"
	res, rerr := lr.CRepo.Session.Exec(query, l.LabName, l.LabCode, l.LabID)
	if rerr != nil {
		fmt.Println("ERROR : LabSQLRepo Modify ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Lab updated successfully!")
		return nil
	}
	return errors.New("ERROR_LAB_NOT_UPDATED")
}

// Get all labs
func (lr *LabSQLRepo) GetAll(role string, labId int) ([]models.Lab, error) {
	var labs []models.Lab
	var st *dbr.SelectStmt

	var q string
	if role == "superadmin" {
		q = "SELECT lab_id, lab_name,lab_code,created_on,created_by FROM Lab where lab_id>0"
		st = lr.CRepo.Session.SelectBySql(q)
	} else {
		q = "SELECT lab_id, lab_name,lab_code,created_on,created_by FROM Lab where lab_id = ?"
		st = lr.CRepo.Session.SelectBySql(q, labId)
	}
	rowsCount, rerr := st.Load(&labs)
	if rerr != nil {
		fmt.Println("ERROR : LabSQLRepo GetAll ", rerr)
		return nil, rerr
	}
	if rowsCount > 0 && len(labs) > 0 {
		fmt.Println("SUCCESS : LabSQLRepo Len(labs) ", len(labs))
		return labs, nil
	}
	fmt.Println("ERROR : LabSQLRepo GetAll ", "not found")
	return labs, errors.New("ERRORCODE_LABS_NOT_FOUND")
}

// Get all labs
// func (lr *LabSQLRepo) GetAllLabsDeptsBranchesServices() ([]models.Lab, error) {
// 	var labs []models.Lab
// 	// var st *dbr.SelectStmt

//	}
//
// Get a single lab by its ID
func (lr *LabSQLRepo) GetOne(l models.Lab) (models.Lab, error) {
	var query string
	var st *dbr.SelectStmt
	if l.LabCode == "" {
		query = "SELECT lab_id, lab_name,lab_code FROM Lab WHERE lab_code = ?"
		st = lr.CRepo.Session.SelectBySql(query, l.LabCode)
	} else {
		query = "SELECT lab_id, lab_name,lab_code FROM Lab WHERE lab_id = ?"
		st = lr.CRepo.Session.SelectBySql(query, l.LabID)
	}
	var lab models.Lab
	err := st.LoadOne(&lab)
	if err != nil {
		fmt.Println("ERROR : LabSQLRepo GetOne ", err)
		return lab, err
	}

	return lab, nil
}

// Delete a lab by its ID
func (lr *LabSQLRepo) DeleteOne(l models.Lab) error {
	query := "DELETE FROM Lab WHERE lab_id = ?"
	res, rerr := lr.CRepo.Session.Exec(query, l.LabID)
	if rerr != nil {
		fmt.Println("ERROR : LabSQLRepo DeleteOne ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Lab deleted successfully!")
		return nil
	}
	return errors.New("ERROR_LAB_NOT_DELETED")
}

// ...........................................................................................................................................................
// BRANCHES OF LAB
type BranchRepo interface {
	Create(models.Branch) error
	Modify(models.Branch) error
	BulkCreate([]models.Branch) error
	GetAll(role string, labId int) ([]models.Branch, error)
	GetOne(models.Branch) (models.Branch, error)
	DeleteOne(models.Branch) error
	GetAllLabsAllBranches([]models.Lab) ([]models.LabsBranches, error)
	GetAllBranchesAllDepts([]models.Branch) ([]models.BranchDepts, error)
}

type BranchSQLRepo struct {
	CRepo *db.SQLRepo
}

func NewBranchRepo(sr *db.SQLRepo) *BranchSQLRepo {
	return &BranchSQLRepo{
		CRepo: sr,
	}
}

// Create a new branch
func (br *BranchSQLRepo) Create(b models.Branch) error {
	codeLength, _ := strconv.Atoi(os.Getenv("codelength")) // Fetch code length from environment variable
	code, err := generateRandomString(codeLength)          // Generate random string
	if err != nil {
		fmt.Println("ERROR: HospitalSQLRepo Create - Code generation error", err)
		return errors.New("ERROR_CODE_GENERATION_ERR")
	}
	b.BranchCode = code // Set generated hospital code

	query := "INSERT INTO Branch (branch_name, lab_id, address, branch_code, city_id) VALUES (?, ?, ?, ?, ?)"
	res, err := br.CRepo.Session.Exec(query, b.BranchName, b.LabID, b.Address, b.BranchCode, b.CityID)
	if err != nil {
		fmt.Println("ERROR : BranchSQLRepo Create ", err)
		return err
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("Branch inserted successfully with ID - ", inc)
		return nil
	}
	return errors.New("ERROR_BRANCH_NOT_INSERTED")
}

// BulkCreate inserts multiple branches into the Branch table
func (br *BranchSQLRepo) BulkCreate(branches []models.Branch) error {
	if len(branches) == 0 {
		return errors.New("ERROR_BRANCHES_LENGTH_IS_ZERO")
	}
	query := "INSERT INTO Branch (branch_name, lab_id, address, branch_code, city_id) VALUES (?, ?, ?, ?, ?)"
	for _, branch := range branches {
		codeLength, _ := strconv.Atoi(os.Getenv("codelength")) // Fetch code length from environment variable
		code, err := generateRandomString(codeLength)          // Generate random string
		if err != nil {
			fmt.Println("ERROR: HospitalSQLRepo Create - Code generation error", err)
			return errors.New("ERROR_CODE_GENERATION_ERR")
		}
		branch.BranchCode = code // Set generated hospital code

		res, err := br.CRepo.Session.Exec(query, branch.BranchName, branch.LabID, branch.Address, branch.BranchCode, branch.CityID)
		if err != nil {
			fmt.Println("ERROR : BranchSQLRepo BulkCreate ", err)
			return err
		}
		inc, _ := res.LastInsertId()
		rac, _ := res.RowsAffected()
		if inc > 0 || rac > 0 {
			fmt.Println("Branch inserted successfully with ID - ", inc)
		} else {
			fmt.Println("Branch failed to insert - ", branch.BranchName, branch.LabID)
		}
	}
	return nil
}

// Modify updates an existing branch
func (br *BranchSQLRepo) Modify(b models.Branch) error {
	query := "UPDATE Branch SET branch_name = ?, lab_id = ?, address = ?, branch_code = ?, city_id = ? WHERE branch_id = ?"
	res, err := br.CRepo.Session.Exec(query, b.BranchName, b.LabID, b.Address, b.BranchCode, b.CityID, b.BranchID)
	if err != nil {
		fmt.Println("ERROR : BranchSQLRepo Modify ", err)
		return err
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Branch updated successfully!")
		return nil
	}
	return errors.New("ERROR_BRANCH_NOT_UPDATED")
}

// GetAll retrieves all branches
func (br *BranchSQLRepo) GetAll(role string, labId int) ([]models.Branch, error) {
	var branches []models.Branch
	var st *dbr.SelectStmt
	var q string
	if role == "superadmin" {
		q = `SELECT branch_id, branch_name, l.lab_id,l.lab_name, address, branch_code, city_id, created_at FROM Branch AS b INNER JOIN lab AS l ON l.lab_id = b.lab_id WHERE l.lab_id > 1 ;`
		st = br.CRepo.Session.SelectBySql(q)
	} else {
		q = `SELECT branch_id, branch_name, l.lab_id,l.lab_name, address, branch_code, city_id, created_at FROM Branch AS b INNER JOIN lab AS l ON l.lab_id = b.lab_id WHERE b.lab_id = ? ;`
		st = br.CRepo.Session.SelectBySql(q, labId)

	}
	rowsCount, err := st.Load(&branches)
	if err != nil {
		fmt.Println("ERROR : BranchSQLRepo GetAll ", err)
		return nil, err
	}
	if rowsCount > 0 {
		fmt.Println("SUCCESS : BranchSQLRepo GetAll count:", len(branches))
		return branches, nil
	}
	return branches, errors.New("Not Found")
}

// GetOne retrieves a single branch by ID
func (br *BranchSQLRepo) GetOne(b models.Branch) (models.Branch, error) {
	query := "SELECT branch_id, branch_name, lab_id, address, branch_code, city_id, created_at FROM Branch WHERE branch_id = ?"
	var branch models.Branch
	err := br.CRepo.Session.SelectBySql(query, b.BranchID).LoadOne(&branch)
	if err != nil {
		fmt.Println("ERROR : BranchSQLRepo GetOne ", err)
		return branch, err
	}
	return branch, nil
}

// DeleteOne deletes a branch by ID
func (br *BranchSQLRepo) DeleteOne(b models.Branch) error {
	query := "DELETE FROM Branch WHERE branch_id = ?"
	res, err := br.CRepo.Session.Exec(query, b.BranchID)
	if err != nil {
		fmt.Println("ERROR : BranchSQLRepo DeleteOne ", err)
		return err
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Branch deleted successfully!")
		return nil
	}
	return errors.New("ERROR_BRANCH_NOT_DELETED")
}

// Get all users
func (br *BranchSQLRepo) GetAllLabsAllBranches(labs []models.Lab) ([]models.LabsBranches, error) {
	var labsBranches []models.LabsBranches
	// if r == "superadmin" {

	// }
	for _, lab := range labs {
		var users []models.Branch
		var lu models.LabsBranches
		lu.LabID = lab.LabID
		lu.LabCode = lab.LabCode
		lu.CreatedOn = lab.CreatedOn
		lu.LabName = lab.LabName
		// lu.ValidityDate = lab.ValidityDate
		q := "SELECT branch_id, branch_name, lab_id, address, branch_code, city_id, created_at FROM Branch WHERE lab_id = ? ; "
		// q := "SELECT user_id, NAME,username, email, role, phone_number, created_at ,PASSWORD FROM USER u WHERE lab_id = ? ;"
		rowsCount, rerr := br.CRepo.Session.SelectBySql(q, lab.LabID).Load(&users)
		if rerr != nil {
			fmt.Println("ERROR : GetAllLabsAllBranches GetAll ", rerr)
			// return nil, rerr
		}
		if rowsCount > 0 && len(users) > 0 {
			fmt.Println("SUCCESS : GetAllLabsAllBranches Len(users) ", len(users))
			lu.Branches = append(lu.Branches, users...)
			// return users, nil
		}
		labsBranches = append(labsBranches, lu)
	}
	if len(labsBranches) > 0 {
		fmt.Println("SUCCESS : GetAllLabsAllBranches GetAll Found", len(labsBranches))
		return labsBranches, nil

	}
	fmt.Println("ERROR : GetAllLabsAllBranches GetAll ", "not found")
	return labsBranches, errors.New("Not Found")
}

// Get all users
func (br *BranchSQLRepo) GetAllBranchesAllDepts(brnches []models.Branch) ([]models.BranchDepts, error) {
	var depts []models.Department
	var brnchDepts []models.BranchDepts
	// if r == "superadmin" {

	// }
	for _, b := range brnches {
		var lu models.BranchDepts
		lu.BranchCode = b.BranchCode
		lu.Address = b.Address
		lu.BranchName = b.BranchName
		lu.BranchID = b.BranchID
		lu.LabID = b.LabID
		q := "SELECT department_id, branch_id, department_name, description,lab_id,branch_id FROM department WHERE branch_id = ? AND lab_id = ?"
		rowsCount, rerr := br.CRepo.Session.SelectBySql(q, b.BranchID, b.LabID).Load(&depts)
		if rerr != nil {
			fmt.Println("ERROR : GetAllLabsAllBranches GetAll ", rerr)
			return brnchDepts, rerr
		}
		if rowsCount > 0 && len(depts) > 0 {
			fmt.Println("SUCCESS : GetAllLabsAllBranches Len(depts) ", len(depts), " for branch - ", b.BranchID)
			lu.Departments = append(lu.Departments, depts...)
		}
		brnchDepts = append(brnchDepts, lu)
	}
	if len(brnchDepts) > 0 {
		fmt.Println("SUCCESS : GetAllLabsAllBranches GetAll Found", len(brnchDepts))
		return brnchDepts, nil
	}
	fmt.Println("ERROR : GetAllLabsAllBranches GetAll ", "not found")
	return brnchDepts, errors.New("Not Found")
}

// ...........................................................................................................................................................................
type UserrRepo interface {
	Create(models.Userr) error
	Modify(models.Userr) error
	GetAll(r string, id int) ([]models.ResponseUser, error)
	GetAllLabsAllUsers([]models.Lab) ([]models.LabsUsers, error)
	GetOne(models.Userr) (models.Userr, error)
	GetUser(models.Userr) (models.Userr, error)
	DeleteOne(models.Userr) error
}
type UserSQLRepo struct {
	CRepo        *db.SQLRepo
	UserProducer *kfka.KafkaProducer
}

func NewUserrRepo(sr *db.SQLRepo, up *kfka.KafkaProducer) *UserSQLRepo {
	return &UserSQLRepo{
		CRepo:        sr,
		UserProducer: up,
	}
}

// Create a new user
func (ur *UserSQLRepo) Create(u models.Userr) error {
	ct := time.Now()
	layout := "2006-01-02 15:04:05"
	u.CreatedAt = ct.Format(layout)
	u.Role = strings.ToLower(u.Role)
	query := "INSERT INTO User (name, email,username, password, role, phone_number, created_at,lab_id) VALUES (?, ?, ?, ?, ?, ?,?,?)"
	res, rerr := ur.CRepo.Session.Exec(query, u.Name, u.Email, u.Username, u.Password, u.Role, u.PhoneNumber, u.CreatedAt, u.LabID)
	if rerr != nil {
		fmt.Println("ERROR : UserSQLRepo Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		c := kfka.CustomMessage{
			Message: "Successfuly User created",
			Status:  "success",
			Data:    u,
		}
		fmt.Println("User inserted successfully! with sql id - ", inc)
		ur.UserProducer.ProduceMessage(c)
		return nil
	}
	return errors.New("ERROR_USER_NOT_INSERTED")
}

// Modify an existing user // cannot modify username
func (ur *UserSQLRepo) Modify(u models.Userr) error {
	query := "UPDATE User SET name = ?, email = ?, password = ?, role = ?, phone_number = ? WHERE user_id = ?"
	res, rerr := ur.CRepo.Session.Exec(query, u.Name, u.Email, u.Password, u.Role, u.PhoneNumber, u.UserID)
	if rerr != nil {
		fmt.Println("ERROR : UserSQLRepo Modify ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("User updated successfully!")
		return nil
	}
	return errors.New("ERROR_USER_NOT_UPDATED")
}

// Get all users
func (ur *UserSQLRepo) GetAll(r string, id int) ([]models.ResponseUser, error) {
	var users []models.ResponseUser
	var query string
	// if r == "superadmin" {
	// query = `SELECT user_id, NAME,username, email, role, phone_number, created_at,l.lab_id,l.lab_name,l.lab_code
	// FROM USER u
	// INNER JOIN lab l ON l.lab_id = u.lab_id;`
	// } else {
	query = `SELECT user_id, NAME,username, email, role, phone_number, created_at,l.lab_id,l.lab_name,l.lab_code
					FROM USER u
					INNER JOIN lab l ON l.lab_id = u.lab_id AND u.lab_id = ?;`
	// }
	rowsCount, rerr := ur.CRepo.Session.SelectBySql(query, id).Load(&users)
	if rerr != nil {
		fmt.Println("ERROR : UserSQLRepo GetAll ", rerr)
		return nil, rerr
	}

	if rowsCount > 0 && len(users) > 0 {
		fmt.Println("SUCCESS : UserSQLRepo Len(users) ", len(users))
		return users, nil
	}
	fmt.Println("ERROR : UserSQLRepo GetAll ", "not found")
	return users, errors.New("Not Found")
}

// Get all users
func (ur *UserSQLRepo) GetAllLabsAllUsers(labs []models.Lab) ([]models.LabsUsers, error) {
	var users []models.Userr
	var labsUsers []models.LabsUsers
	// if r == "superadmin" {

	// }
	for _, lab := range labs {
		var lu models.LabsUsers
		lu.LabID = lab.LabID
		lu.LabCode = lab.LabCode
		lu.CreatedOn = lab.CreatedOn
		lu.LabName = lab.LabName
		// lu.ValidityDate = lab.ValidityDate
		q := "SELECT user_id, NAME,username, email, role, phone_number, created_at ,PASSWORD FROM USER u WHERE lab_id = ? ;"
		rowsCount, rerr := ur.CRepo.Session.SelectBySql(q, lab.LabID).Load(&users)
		if rerr != nil {
			fmt.Println("ERROR : UserSQLRepo GetAll ", rerr)
			// return nil, rerr
		}
		if rowsCount > 0 && len(users) > 0 {
			fmt.Println("SUCCESS : UserSQLRepo Len(users) ", len(users))
			lu.Users = append(lu.Users, users...)
			// return users, nil
		}
		labsUsers = append(labsUsers, lu)
	}
	if len(labsUsers) > 0 {
		fmt.Println("SUCCESS : UserSQLRepo GetAll Found", len(labsUsers))
		return labsUsers, nil

	}
	fmt.Println("ERROR : UserSQLRepo GetAll ", "not found")
	return labsUsers, errors.New("Not Found")
}

// Get a single user by its ID
func (ur *UserSQLRepo) GetOne(u models.Userr) (models.Userr, error) {
	query := "SELECT user_id, name,username, email, role, phone_number, address, created_at FROM User WHERE user_id = ?"
	var user models.Userr
	err := ur.CRepo.Session.SelectBySql(query, u.UserID).LoadOne(&user)
	if err != nil {
		fmt.Println("ERROR : UserSQLRepo GetOne ", err)
		return user, err
	}

	return user, nil
}

// get user by username and password
func (ur *UserSQLRepo) GetUser(u models.Userr) (models.Userr, error) {
	query := "SELECT user_id, name, username,role,lab_id FROM User WHERE username = ? AND password = ?"
	var user models.Userr
	err := ur.CRepo.Session.SelectBySql(query, u.Username, u.Password).LoadOne(&user)
	if err != nil {
		log.Println("ERROR : UserSQLRepo GetOne ", err)
		return user, err
	}
	log.Println("User found ", user)
	return user, nil
}

// Delete a user by its ID
func (ur *UserSQLRepo) DeleteOne(u models.Userr) error {
	query := "DELETE FROM User WHERE user_id = ?"
	res, rerr := ur.CRepo.Session.Exec(query, u.UserID)
	if rerr != nil {
		fmt.Println("ERROR : UserSQLRepo DeleteOne ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("User deleted successfully!")
		return nil
	}
	return errors.New("ERROR_USER_NOT_DELETED")
}

// ...........................................................................................................................
type DoctorRepo interface {
	Create(models.Doctor) error
	BulkCreate([]models.Doctor) error
	Modify(models.Doctor) error
	GetAll() ([]models.Doctor, error)
	GetOne(models.Doctor) (models.Doctor, error)
	DeleteOne(models.Doctor) error
}
type DoctorSQLRepo struct {
	CRepo *db.SQLRepo
}

func NewDoctorRepo(sr *db.SQLRepo) *DoctorSQLRepo {
	return &DoctorSQLRepo{
		CRepo: sr,
	}
}

// Create a new doctor
func (dr *DoctorSQLRepo) Create(d models.Doctor) error {
	query := "INSERT INTO Doctor (user_id, specialization, years_of_experience, license_number, clinic_hours) VALUES (?, ?, ?, ?, ?)"
	res, rerr := dr.CRepo.Session.Exec(query, d.UserID, d.Specialization, d.YearsOfExperience, d.LicenseNumber, d.ClinicHours)
	if rerr != nil {
		fmt.Println("ERROR : DoctorSQLRepo Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("Doctor inserted successfully! with sql id - ", inc)
		return nil
	}
	return errors.New("ERROR_DOCTOR_NOT_INSERTED")
}

// BulkCreate inserts an array of doctors into the Doctor table
func (dr *DoctorSQLRepo) BulkCreate(doctors []models.Doctor) error {
	if len(doctors) > 0 {
		query := "INSERT INTO Doctor (user_id, specialization, years_of_experience, license_number, clinic_hours) VALUES (?, ?, ?, ?,  ?)"
		for _, doctor := range doctors {
			// Execute the insert query
			res, rerr := dr.CRepo.Session.Exec(query, doctor.UserID, doctor.Specialization, doctor.YearsOfExperience, doctor.LicenseNumber, doctor.ClinicHours)
			if rerr != nil {
				fmt.Println("ERROR: DoctorSQLRepo Create ", rerr)
				return rerr
			}
			inc, _ := res.LastInsertId()
			rac, _ := res.RowsAffected()
			if inc > 0 || rac > 0 {
				fmt.Println("Doctor inserted successfully! with sql id - ", inc)
			} else {
				fmt.Println("Doctor failed to insert - ", doctor.UserID, doctor.Specialization)
				// LOGS Implementation remaining
			}
		}
		return nil // Successful insertion, so return nil
	}
	return errors.New("ERROR_DOCTORS_LENGTH_IS_ZERO")
}

// Modify an existing doctor
func (dr *DoctorSQLRepo) Modify(d models.Doctor) error {
	query := "UPDATE Doctor SET user_id = ?, specialization = ?, years_of_experience = ?, license_number = ?, clinic_hours  = ? WHERE doctor_id = ?"
	res, rerr := dr.CRepo.Session.Exec(query, d.UserID, d.Specialization, d.YearsOfExperience, d.LicenseNumber, d.ClinicHours, d.DoctorID)
	if rerr != nil {
		fmt.Println("ERROR : DoctorSQLRepo Modify ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Doctor updated successfully!")
		return nil
	}
	return errors.New("ERROR_DOCTOR_NOT_UPDATED")
}

// Get all doctors
func (dr *DoctorSQLRepo) GetAll() ([]models.Doctor, error) {
	var doctors []models.Doctor
	query := "SELECT doctor_id, user_id, specialization, years_of_experience, license_number, clinic_hours, assigned_lab_id FROM Doctor"
	rowsCount, rerr := dr.CRepo.Session.SelectBySql(query).Load(&doctors)
	if rerr != nil {
		fmt.Println("ERROR : DoctorSQLRepo GetAll ", rerr)
		return nil, rerr
	}

	if rowsCount > 0 && len(doctors) > 0 {
		fmt.Println("SUCCESS : DoctorSQLRepo Len(doctors) ", len(doctors))
		return doctors, nil
	}
	fmt.Println("ERROR : DoctorSQLRepo GetAll ", "not found")
	return doctors, errors.New("Not Found")
}

// Get a single doctor by ID
func (dr *DoctorSQLRepo) GetOne(d models.Doctor) (models.Doctor, error) {
	query := "SELECT doctor_id, user_id, specialization, years_of_experience, license_number, clinic_hours, assigned_lab_id FROM Doctor WHERE doctor_id = ?"
	var doctor models.Doctor
	err := dr.CRepo.Session.SelectBySql(query, d.DoctorID).LoadOne(&doctor)
	if err != nil {
		fmt.Println("ERROR : DoctorSQLRepo GetOne ", err)
		return doctor, err
	}

	return doctor, nil
}

// Delete a doctor by ID
func (dr *DoctorSQLRepo) DeleteOne(d models.Doctor) error {
	query := "DELETE FROM Doctor WHERE doctor_id = ?"
	res, rerr := dr.CRepo.Session.Exec(query, d.DoctorID)
	if rerr != nil {
		fmt.Println("ERROR : DoctorSQLRepo DeleteOne ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Doctor deleted successfully!")
		return nil
	}
	return errors.New("ERROR_DOCTOR_NOT_DELETED")
}

// ........................................................................
type HospitalRepo interface {
	Create(models.Hospital) error
	BulkCreate([]models.Hospital) error
	// Modify(models.Hospital) error
	GetAll() ([]models.Hospital, error)
	GetOne(models.Hospital) (models.Hospital, error)
	// DeleteOne(models.Hospital) error
}
type HospitalSQLRepo struct {
	HRepo *db.SQLRepo
}

func NewHospitalRepo(sr *db.SQLRepo) *HospitalSQLRepo {
	return &HospitalSQLRepo{
		HRepo: sr,
	}
}

// Create a new Hospital
func (dr *HospitalSQLRepo) Create(hospital models.Hospital) error {
	l, _ := strconv.Atoi(os.Getenv("codelength"))
	c, e := generateRandomString(l)
	if e != nil {
		fmt.Println("Hospital inserted successfully! with sql id - ")
		return errors.New("ERRORCDE_CODE_GENERATION_ERR")
	}
	hospital.HospitalCode = c
	query := `INSERT INTO hospital (hospital_name, city_id, state_id, address, phone_number,hospital_code) VALUES (?, ?, ?, ?, ?,?)`
	res, rerr := dr.HRepo.Session.Exec(query, hospital.HospitalName, hospital.CityID, hospital.StateID, hospital.Address, hospital.PhoneNumber, hospital.HospitalCode)
	if rerr != nil {
		fmt.Println("ERROR : HospitalSQLRepo Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("Hospital inserted successfully! with sql id - ", inc)
		return nil
	}
	return errors.New("ERROR_Hospital_NOT_INSERTED")
}

// BulkCreate inserts an array of hospitals into the Hospital table
func (hr *HospitalSQLRepo) BulkCreate(hospitals []models.Hospital) error {
	if len(hospitals) > 0 {
		query := "INSERT INTO Hospital (hospital_name, city_id, state_id, address, phone_number, hospital_code) VALUES (?, ?, ?, ?, ?, ?)"
		for _, hospital := range hospitals {
			// Generate hospital code using random string generator
			codeLength, _ := strconv.Atoi(os.Getenv("codelength")) // Fetch code length from environment variable
			code, err := generateRandomString(codeLength)          // Generate random string
			if err != nil {
				fmt.Println("ERROR: HospitalSQLRepo Create - Code generation error", err)
				return errors.New("ERROR_CODE_GENERATION_ERR")
			}
			hospital.HospitalCode = code // Set generated hospital code

			// Execute the insert query
			res, rerr := hr.HRepo.Session.Exec(query, hospital.HospitalName, hospital.CityID, hospital.StateID, hospital.Address, hospital.PhoneNumber, hospital.HospitalCode)
			if rerr != nil {
				fmt.Println("ERROR: HospitalSQLRepo Create ", rerr)
				return rerr
			}
			inc, _ := res.LastInsertId()
			rac, _ := res.RowsAffected()
			if inc > 0 || rac > 0 {
				fmt.Println("Hospital inserted successfully! with sql id - ", inc)
			} else {
				fmt.Println("Hospital failed to insert - ", hospital.HospitalName, hospital.CityID)
				// LOGS Implementation remaining
			}
		}
		return nil // Successful insertion, so return nil
	}
	return errors.New("ERROR_HOSPITALS_LENGTH_IS_ZERO")
}

// Get all hospitals
func (dr *HospitalSQLRepo) GetAll() ([]models.Hospital, error) {
	var hs []models.Hospital
	query := "SELECT * from hospital"
	rowsCount, rerr := dr.HRepo.Session.SelectBySql(query).Load(&hs)
	if rerr != nil {
		fmt.Println("ERROR : HospitalSQLRepo GetAll ", rerr)
		return nil, rerr
	}
	if rowsCount > 0 && len(hs) > 0 {
		fmt.Println("SUCCESS : HospitalSQLRepo Len(doctors) ", len(hs))
		return hs, nil
	}
	fmt.Println("ERROR : HspitalSQLRepo GetAll ", "not found")
	return hs, errors.New("Not Found")
}

// Get a single HOSPITAL by ID or code
func (dr *HospitalSQLRepo) GetOne(d models.Hospital) (models.Hospital, error) {
	var query string
	var st *dbr.SelectStmt
	if d.HospitalID != 0 {
		query = "SELECT * FROM Hospital WHERE hospital_id = ?"
		st = dr.HRepo.Session.SelectBySql(query, d.HospitalID)
	} else {
		if d.HospitalCode != "" {
			query = "SELECT * FROM Hospital WHERE hospital_code = ?"
			st = dr.HRepo.Session.SelectBySql(query, d.HospitalCode)
		}
	}
	var h models.Hospital
	err := st.LoadOne(&h)
	if err != nil {
		fmt.Println("ERROR : DoctorSQLRepo GetOne ", err)
		return h, err
	}
	return h, nil
}

/*

// Modify an existing doctor
func (dr *HospitalSQLRepo) Modify(d models.Hospital) error {
	query := "UPDATE Hospital SET user_id = ?, specialization = ?, years_of_experience = ?, license_number = ?, clinic_hours = ?, assigned_lab_id = ? WHERE doctor_id = ?"
	res, rerr := dr.HRepo.Session.Exec(query, d.UserID, d.Specialization, d.YearsOfExperience, d.LicenseNumber, d.ClinicHours, d.AssignedLabID, d.DoctorID)
	if rerr != nil {
		fmt.Println("ERROR : HospitalSQLRepo Modify ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Hospital updated successfully!")
		return nil
	}
	return errors.New("ERROR_Hospital_NOT_UPDATED")
}
// Get all doctors
func (dr *HospitalSQLRepo) GetAll() ([]models.Doctor, error) {
	var doctors []models.Doctor
	query := "SELECT doctor_id, user_id, specialization, years_of_experience, license_number, clinic_hours, assigned_lab_id FROM Doctor"
	rowsCount, rerr := dr.CRepo.Session.SelectBySql(query).Load(&doctors)
	if rerr != nil {
		fmt.Println("ERROR : DoctorSQLRepo GetAll ", rerr)
		return nil, rerr
	}

	if rowsCount > 0 && len(doctors) > 0 {
		fmt.Println("SUCCESS : DoctorSQLRepo Len(doctors) ", len(doctors))
		return doctors, nil
	}
	fmt.Println("ERROR : DoctorSQLRepo GetAll ", "not found")
	return doctors, errors.New("Not Found")
}

// Get a single doctor by ID
func (dr *DoctorSQLRepo) GetOne(d models.Doctor) (models.Doctor, error) {
	query := "SELECT doctor_id, user_id, specialization, years_of_experience, license_number, clinic_hours, assigned_lab_id FROM Doctor WHERE doctor_id = ?"
	var doctor models.Doctor
	err := dr.CRepo.Session.SelectBySql(query, d.DoctorID).LoadOne(&doctor)
	if err != nil {
		fmt.Println("ERROR : DoctorSQLRepo GetOne ", err)
		return doctor, err
	}

	return doctor, nil
}

// Delete a doctor by ID
func (dr *DoctorSQLRepo) DeleteOne(d models.Doctor) error {
	query := "DELETE FROM Doctor WHERE doctor_id = ?"
	res, rerr := dr.CRepo.Session.Exec(query, d.DoctorID)
	if rerr != nil {
		fmt.Println("ERROR : DoctorSQLRepo DeleteOne ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Doctor deleted successfully!")
		return nil
	}
	return errors.New("ERROR_DOCTOR_NOT_DELETED")
}
*/
// ........................................................................
type HospitalBranchRepo interface {
	Create(models.HospitalBranch) error
	Modify(models.HospitalBranch) error
	GetAll(int) ([]models.HospitalBranch, error)
	GetOne(models.HospitalBranch) (models.HospitalBranch, error)
	BulkCreate([]models.HospitalBranch) error
	// DeleteOne(models.HospitalBranch) error
}
type HospitalBranchSQLRepo struct {
	HBRepo *db.SQLRepo
}

func NewHospitalBranchRepo(sr *db.SQLRepo) *HospitalBranchSQLRepo {
	return &HospitalBranchSQLRepo{
		HBRepo: sr,
	}
}

// Create a new hospitalbranch
func (br *HospitalBranchSQLRepo) Create(b models.HospitalBranch) error {
	l, _ := strconv.Atoi(os.Getenv("codelength"))
	c, e := generateRandomString(l)
	if e != nil {
		fmt.Println("Hospital inserted successfully! with sql id - ")
		return errors.New("ERRORCDE_CODE_GENERATION_ERR")
	}
	b.BranchCode = c
	query := `INSERT INTO hospital_branch (hospital_id, branch_name, address, phone_number, branchCode)
              VALUES (?, ?, ?, ?, ?);`
	res, err := br.HBRepo.Session.Exec(query, b.HospitalID, b.BranchName, b.Address, b.PhoneNumber, b.BranchCode)

	if err != nil {
		fmt.Println("ERROR : HospitalBranchSQLRepo Create ", err)
		return err
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("HospitalBranchSQLRepo inserted successfully! with sql id - ", inc)
		return nil
	}
	return errors.New("ERROR_HospitalBranchSQLRepo_NOT_INSERTED")
}

// BulkCreate inserts an array of hospital branches into the HospitalBranch table
func (hbr *HospitalBranchSQLRepo) BulkCreate(branches []models.HospitalBranch) error {
	if len(branches) > 0 {
		query := "INSERT INTO HospitalBranch (hospital_id, branch_name, address, phone_number, branch_code) VALUES (?, ?, ?, ?, ?)"
		for _, branch := range branches {
			// Generate Branch Code
			l, _ := strconv.Atoi(os.Getenv("codelength")) // Assuming you have a length in environment variable
			c, e := generateRandomString(l)               // Implement this function to generate a random string
			if e != nil {
				fmt.Println("ERROR: Branch code generation failed", e)
				return errors.New("ERROR_BRANCH_CODE_GENERATION")
			}
			branch.BranchCode = c

			// Execute the insert query
			res, rerr := hbr.HBRepo.Session.Exec(query, branch.HospitalID, branch.BranchName, branch.Address, branch.PhoneNumber, branch.BranchCode)
			if rerr != nil {
				fmt.Println("ERROR: HospitalBranchSQLRepo Create ", rerr)
				return rerr
			}
			inc, _ := res.LastInsertId()
			rac, _ := res.RowsAffected()
			if inc > 0 || rac > 0 {
				fmt.Println("Branch inserted successfully! with sql id - ", inc)
			} else {
				fmt.Println("Branch failed to insert - ", branch.BranchName)
				// LOGS Implementation remaining
			}
		}
		return nil // Successful insertion, so return nil
	}
	return errors.New("ERROR_BRANCHES_LENGTH_IS_ZERO")
}

// Modify an existing branch
func (br *HospitalBranchSQLRepo) Modify(branch models.HospitalBranch) error {
	query := `UPDATE hospital_branch
              SET  branch_name = ?, address = ?, phone_number = ?
              WHERE branch_id = ?;`
	res, rerr := br.HBRepo.Session.Exec(query, branch.BranchName, branch.Address, branch.PhoneNumber, branch.BranchID)
	if rerr != nil {
		fmt.Println("ERROR : HospitalBranchSQLRepo Modify ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("HospitalBranchSQLRepo updated successfully!")
		return nil
	}
	return errors.New("ERROR_HospitalBranchSQLRepo_NOT_UPDATED")
}

// Get all  branches of hospital
func (br *HospitalBranchSQLRepo) GetAll(id int) ([]models.HospitalBranch, error) {
	var hbs []models.HospitalBranch
	query := `SELECT branch_id, hospital_id, branch_name, address, phone_number, branchCode
              FROM hospital_branch
              WHERE  hospital_id = ?`
	rowsCount, rerr := br.HBRepo.Session.SelectBySql(query, id).Load(&hbs)
	if rerr != nil {
		fmt.Println("ERROR : HospitalBranchSQLRepo GetAll ", rerr)
		return nil, rerr
	}
	if rowsCount > 0 && len(hbs) > 0 {
		fmt.Println("SUCCESS : BranchSQLRepo Len(branches) ", len(hbs))
		return hbs, nil
	}
	fmt.Println("ERROR : BranchSQLRepo GetAll ", "not found")
	return hbs, errors.New("Not Found")
}

// Get a single branch by its ID
func (br *HospitalBranchSQLRepo) GetOne(b models.HospitalBranch) (models.HospitalBranch, error) {
	q := `SELECT branch_id, hospital_id, branch_name, address, phone_number, branchCode
              FROM hospital_branch
              WHERE branch_id = ?;`
	var branch models.HospitalBranch
	err := br.HBRepo.Session.SelectBySql(q, b.BranchID).LoadOne(&branch)
	if err != nil {
		fmt.Println("ERROR : BranchSQLRepo GetOne ", err)
		return branch, err
	}

	return branch, nil
}

// // Delete a branch by its ID
//
//	func (br *BranchSQLRepo) DeleteOne(b models.Branch) error {
//		query := "DELETE FROM Branch WHERE branch_id = ?"
//		res, rerr := br.CRepo.Session.Exec(query, b.BranchID)
//		if rerr != nil {
//			fmt.Println("ERROR : BranchSQLRepo DeleteOne ", rerr)
//			return rerr
//		}
//		rac, _ := res.RowsAffected()
//		if rac > 0 {
//			fmt.Println("Branch deleted successfully!")
//			return nil
//		}
//		return errors.New("ERROR_BRANCH_NOT_DELETED")
//	}
//
// ..............................................................................................................................................
type HospitalDoctorRepo interface {
	// Create(models.HospitalDoctor) error
	BulkCreate([]models.HospitalDoctor) error
	// Modify(models.HospitalBranch) error
	GetAll(int) ([]models.HospitalDoctor, error)
	GetOne(int) (models.HospitalDoctor, error)
	// BulkCreate([]models.HospitalDoctor) error
	// DeleteOne(models.HospitalDoctor) error
}
type HospitalDoctorStruct struct {
	HBRepo *db.SQLRepo
}

func NewHospitalDoctorRepo(sr *db.SQLRepo) *HospitalDoctorStruct {
	return &HospitalDoctorStruct{
		HBRepo: sr,
	}
}

// BulkCreate inserts multiple hospital-doctor relationships
func (hdr *HospitalDoctorStruct) BulkCreate(hospitalDoctors []models.HospitalDoctor) error {
	if len(hospitalDoctors) > 0 {
		query := "INSERT INTO hospitalDoctor (hospital_id, doctor_id, role, start_date, end_date) VALUES (?, ?, ?, ?, ?)" //role i consultant resiedent
		for _, hd := range hospitalDoctors {
			res, rerr := hdr.HBRepo.Session.Exec(query, hd.HospitalID, hd.DoctorID, hd.Role, hd.StartDate, hd.EndDate)
			if rerr != nil {
				fmt.Println("ERROR: HospitalDoctorSQLRepo Create", rerr)
				return rerr
			}
			inc, _ := res.LastInsertId()
			rac, _ := res.RowsAffected()
			if inc > 0 || rac > 0 {
				fmt.Println("HospitalDoctor mapping inserted successfully! with sql id - ", inc)
			} else {
				fmt.Println("Failed to insert mapping between hospital_id:", hd.HospitalID, "and doctor_id:", hd.DoctorID)
				// LOGS Implementation remaining
			}
		}
		return nil // Successful insertion, return nil
	}
	return errors.New("ERROR_HOSPITAL_DOCTOR_LENGTH_IS_ZERO")
}

// GetAll fetches all doctors of a hospital
func (hdr *HospitalDoctorStruct) GetAll(h_id int) ([]models.HospitalDoctor, error) {
	var hospitalDoctors []models.HospitalDoctor
	query := `SELECT hospital_doctor_id, hospital_id, doctor_id, role, start_date, end_date FROM HospitalDoctor WHERE  hospital_id = ?`
	rowsCount, err := hdr.HBRepo.Session.SelectBySql(query, h_id).Load(&hospitalDoctors)
	if err != nil {
		fmt.Println("ERROR: HospitalDoctorSQLRepo GetAll", err)
		return nil, err
	}

	if rowsCount > 0 && len(hospitalDoctors) > 0 {
		fmt.Println("SUCCESS: HospitalDoctorSQLRepo Len(hospitalDoctors)", len(hospitalDoctors))
		return hospitalDoctors, nil
	}

	fmt.Println("ERROR: HospitalDoctorSQLRepo GetAll", "not found")
	return hospitalDoctors, errors.New("Not Found")
}

// GetOne fetches a single hospital-doctor relationship by its ID
func (hdr *HospitalDoctorStruct) GetOne(id int) (models.HospitalDoctor, error) {
	var hospitalDoctor models.HospitalDoctor
	query := `SELECT hospital_doctor_id, hospital_id, doctor_id, role, start_date, end_date
	          FROM HospitalDoctor WHERE hd_id = ?`
	err := hdr.HBRepo.Session.SelectBySql(query, id).LoadOne(&hospitalDoctor)
	if err != nil {
		fmt.Println("ERROR: HospitalDoctorSQLRepo GetOne", err)
		return hospitalDoctor, err
	}

	return hospitalDoctor, nil
}
