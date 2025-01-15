package services

import (
	models "repogin/internal/models/masters"
	repos "repogin/internal/repositories/sql"
)

type CountryServiceStruct struct {
	CountryService repos.ContryRepo
}

func NewCountryService(repo repos.ContryRepo) *CountryServiceStruct {
	return &CountryServiceStruct{
		CountryService: repo,
	}
}

func (cs *CountryServiceStruct) CreateCountryService(country models.Country) error {
	return cs.CountryService.Create(country)
}

func (cs *CountryServiceStruct) UpdateCountryService(country models.Country) error {
	return cs.CountryService.Modify(country)
}

func (cs *CountryServiceStruct) GetAllCountriesService() ([]models.Country, error) {
	return cs.CountryService.GetAll()
}

func (cs *CountryServiceStruct) GetOneCountryService(country models.Country) (models.Country, error) {
	return cs.CountryService.GetOne(country)
}

func (cs *CountryServiceStruct) DeleteOneCountryService(country models.Country) error {
	return cs.CountryService.DeleteOne(country)
}

// ...............................................................................................................................................................
type StateServiceStruct struct {
	StateService repos.StateRepo
}

func NewStateService(repo repos.StateRepo) *StateServiceStruct {
	return &StateServiceStruct{
		StateService: repo,
	}
}

func (ss *StateServiceStruct) CreateStateService(state models.State) error {
	return ss.StateService.Create(state)
}

func (cs *StateServiceStruct) CreateStatesService(states []models.State) error {
	return cs.StateService.BulkCreate(states)
}
func (ss *StateServiceStruct) UpdateStateService(state models.State) error {
	return ss.StateService.Modify(state)
}

func (ss *StateServiceStruct) GetAllStatesService() ([]models.State, error) {
	return ss.StateService.GetAll()
}
func (ss *StateServiceStruct) GetAllStatesOfCountryService(cid int) ([]models.State, error) {
	return ss.StateService.GetAllStatesOfCountry(cid)
}

func (ss *StateServiceStruct) GetOneStateService(state models.State) (models.State, error) {
	return ss.StateService.GetOne(state)
}

func (ss *StateServiceStruct) DeleteOneStateService(state models.State) error {
	return ss.StateService.DeleteOne(state)
}

// ...................................................................................................................................................................
type CityServiceStruct struct {
	CityService repos.CityRepo
}

func NewCityService(repo repos.CityRepo) *CityServiceStruct {
	return &CityServiceStruct{
		CityService: repo,
	}
}

func (cs *CityServiceStruct) CreateCityService(city models.City) error {
	return cs.CityService.Create(city)
}
func (cs *CityServiceStruct) CreateCitiesService(cities []models.City) error {
	return cs.CityService.BulkCreate(cities)
}
func (cs *CityServiceStruct) UpdateCityService(city models.City) error {
	return cs.CityService.Modify(city)
}

func (cs *CityServiceStruct) GetAllCitiesService() ([]models.City, error) {
	return cs.CityService.GetAll()
}

func (cs *CityServiceStruct) GetOneCityService(city models.City) (models.City, error) {
	return cs.CityService.GetOne(city)
}
func (cs *CityServiceStruct) GetAllCitiesOfStateService(stateid int) ([]models.City, error) {
	return cs.CityService.GetAllCitiessOfState(stateid)
}

func (cs *CityServiceStruct) DeleteOneCityService(city models.City) error {
	return cs.CityService.DeleteOne(city)
}

// .....................................................................................................................................................................
type RoleServiceStruct struct {
	RoleService repos.RoleRepo
}

func NewRoleService(repo repos.RoleRepo) *RoleServiceStruct {
	return &RoleServiceStruct{
		RoleService: repo,
	}
}

func (rs *RoleServiceStruct) CreateRoleService(role models.Role) error {
	return rs.RoleService.Create(role)
}

func (rs *RoleServiceStruct) UpdateRoleService(role models.Role) error {
	return rs.RoleService.Modify(role)
}

func (rs *RoleServiceStruct) GetAllRolesService() ([]models.Role, error) {
	return rs.RoleService.GetAll()
}

func (rs *RoleServiceStruct) GetOneRoleService(role models.Role) (models.Role, error) {
	return rs.RoleService.GetOne(role)
}

func (rs *RoleServiceStruct) DeleteOneRoleService(role models.Role) error {
	return rs.RoleService.DeleteOne(role)
}
func (rs *RoleServiceStruct) GetMenusByRole(role models.Role) ([]models.Menu, error) {
	m, e := rs.RoleService.FetchMenusByRole(role.RoleID)
	return createhierarchy(m), e

}
func createhierarchy(data []models.Menu) []models.Menu {
	menuMap := make(map[int][]models.Menu)
	var ParentNodes []models.Menu
	for _, d := range data {
		if d.ParentMenuID == nil {
			// var Node models.Menu
			// Node.Icon = d.Icon
			// Node.Label = d.Label
			// Node.MenuID = d.MenuID
			// Node.Allowed = d.Allowed
			ParentNodes = append(ParentNodes, d)
			// Node.ToURL = d.ToURL
		} else {
			// if _, ok := Menus[Perm.ParentMenuID]; ok {
			menuMap[*d.ParentMenuID] = append(menuMap[*d.ParentMenuID], d)
			// }
		}
	}
	for i, Node := range ParentNodes {
		ParentNodes[i].Items = append(ParentNodes[i].Items, convertTomenu(menuMap[Node.MenuID])...)
		// ParentNodes[i].MenuID = 0
	}
	// for _, d := range data {
	// 	var menuID int
	// 	// var label, icon, toURL string
	// 	// var parentMenuID sql.NullInt64
	// 	// var allowed bool
	// 	// err := rows.Scan(&menuID, &label, &icon, &toURL, &parentMenuID, &allowed)
	// 	// if err != nil {
	// 	// 	return nil, err
	// 	// }

	// 	menu := models.Menu{
	// 		Label: d.Label,
	// 		Icon:  d.Icon,
	// 		// To:          d.,
	// 		ParentMenuID: nil,
	// 		Items:        []models.Menu{},
	// 	}

	// 	if d.ParentMenuID != nil {
	// 		menu.ParentMenuID = new(int)
	// 		*menu.ParentMenuID = int(*d.ParentMenuID)
	// 	}

	// 	if parentMenu, ok := menuMap[menuID]; ok {
	// 		menu.Items = append(parentMenu.Items, menu)
	// 	} else {
	// 		menuMap[menuID] = menu
	// 	}
	// }

	// Find top-level menus (menus without a parent)
	// var topLevelMenus []models.Menu
	// for _, menu := range menuMap {
	// 	if menu.ParentMenuID == nil {
	// 		topLevelMenus = append(topLevelMenus, menu)
	// 	}
	// }

	return ParentNodes
}
func convertTomenu(all []models.Menu) []models.Menu {
	var menuslice []models.Menu
	for _, t := range all {
		var m models.Menu
		m.Icon = t.Icon
		m.Label = t.Label
		m.MenuID = t.MenuID
		m.ParentMenuID = t.ParentMenuID
		// m.ToURL = t.ToURL
		m.Allowed = t.Allowed
		menuslice = append(menuslice, m)
	}
	return menuslice
}

// ......................................................................................................................
// service of Lab
type LabServiceStruct struct {
	LabService repos.LabRepo
}

func NewLabService(repo repos.LabRepo) *LabServiceStruct {
	return &LabServiceStruct{
		LabService: repo,
	}
}

func (ls *LabServiceStruct) CreateLabService(lab models.Lab) error {
	return ls.LabService.Create(lab)
}
func (cs *LabServiceStruct) CreateLabsService(labs []models.Lab) error {
	return cs.LabService.BulkCreate(labs)
}
func (ls *LabServiceStruct) UpdateLabService(lab models.Lab) error {
	return ls.LabService.Modify(lab)
}

func (ls *LabServiceStruct) GetAllLabsService() ([]models.Lab, error) {
	return ls.LabService.GetAll()
}

func (ls *LabServiceStruct) GetOneLabService(lab models.Lab) (models.Lab, error) {
	return ls.LabService.GetOne(lab)
}

func (ls *LabServiceStruct) DeleteOneLabService(lab models.Lab) error {
	return ls.LabService.DeleteOne(lab)
}

// .................................................................................................................................

type BranchServiceStruct struct {
	BranchService repos.BranchRepo
}

func NewBranchService(repo repos.BranchRepo) *BranchServiceStruct {
	return &BranchServiceStruct{
		BranchService: repo,
	}
}

func (bs *BranchServiceStruct) CreateBranchService(branch models.Branch) error {
	return bs.BranchService.Create(branch)
}

// CreateBranchesService handles the bulk creation of branches
func (bs *BranchServiceStruct) CreateBranchesService(branches []models.Branch) error {
	return bs.BranchService.BulkCreate(branches)
}
func (bs *BranchServiceStruct) UpdateBranchService(branch models.Branch) error {
	return bs.BranchService.Modify(branch)
}

func (bs *BranchServiceStruct) GetAllBranchesService() ([]models.Branch, error) {
	return bs.BranchService.GetAll()
}

func (bs *BranchServiceStruct) GetOneBranchService(branch models.Branch) (models.Branch, error) {
	return bs.BranchService.GetOne(branch)
}

func (bs *BranchServiceStruct) DeleteOneBranchService(branch models.Branch) error {
	return bs.BranchService.DeleteOne(branch)
}

// ...................................................................................................

type UserrServiceStruct struct {
	UserService repos.UserrRepo
}

func NewUserrService(repo repos.UserrRepo) *UserrServiceStruct {
	return &UserrServiceStruct{
		UserService: repo,
	}
}

func (us *UserrServiceStruct) CreateUserService(user models.Userr) error {
	return us.UserService.Create(user)
}

func (us *UserrServiceStruct) UpdateUserService(user models.Userr) error {
	return us.UserService.Modify(user)
}

func (us *UserrServiceStruct) GetAllUsersService() ([]models.Userr, error) {
	return us.UserService.GetAll()
}

func (us *UserrServiceStruct) GetOneUserService(user models.Userr) (models.Userr, error) {
	return us.UserService.GetOne(user)
}
func (us *UserrServiceStruct) GetUserLoginService(user models.Userr) (models.Userr, error) {
	return us.UserService.GetUser(user)
}
func (us *UserrServiceStruct) DeleteOneUserService(user models.Userr) error {
	return us.UserService.DeleteOne(user)
}

// ..........................................................................................
type DoctorServiceStruct struct {
	DoctorService repos.DoctorRepo
}

func NewDoctorService(repo repos.DoctorRepo) *DoctorServiceStruct {
	return &DoctorServiceStruct{
		DoctorService: repo,
	}
}

func (ds *DoctorServiceStruct) CreateDoctorService(doctor models.Doctor) error {
	return ds.DoctorService.Create(doctor)
}
func (ds *DoctorServiceStruct) CreateDoctorsService(doctors []models.Doctor) error {
	return ds.DoctorService.BulkCreate(doctors)
}
func (ds *DoctorServiceStruct) UpdateDoctorService(doctor models.Doctor) error {
	return ds.DoctorService.Modify(doctor)
}

func (ds *DoctorServiceStruct) GetAllDoctorsService() ([]models.Doctor, error) {
	return ds.DoctorService.GetAll()
}

func (ds *DoctorServiceStruct) GetOneDoctorService(doctor models.Doctor) (models.Doctor, error) {
	return ds.DoctorService.GetOne(doctor)
}

func (ds *DoctorServiceStruct) DeleteOneDoctorService(doctor models.Doctor) error {
	return ds.DoctorService.DeleteOne(doctor)
}

// ...............................................................................................................................................
type HospitalServiceStruct struct {
	HospitalService repos.HospitalRepo
}

func NewHospitalService(repo repos.HospitalRepo) *HospitalServiceStruct {
	return &HospitalServiceStruct{
		HospitalService: repo,
	}
}

func (ds *HospitalServiceStruct) CreateHospitalService(Hospital models.Hospital) error {
	return ds.HospitalService.Create(Hospital)
}

func (hs *HospitalServiceStruct) CreateHospitalsService(hospitals []models.Hospital) error {
	return hs.HospitalService.BulkCreate(hospitals)
}
func (ds *HospitalServiceStruct) GetAllHospitalsService() ([]models.Hospital, error) {
	return ds.HospitalService.GetAll()
}

func (ds *HospitalServiceStruct) GetOneHospitalService(h models.Hospital) (models.Hospital, error) {
	return ds.HospitalService.GetOne(h)
}

// ..................................................................................
type HospBranchServiceStruct struct {
	HospitalBranchService repos.HospitalBranchRepo
}

func NewHospBranchService(repo repos.HospitalBranchRepo) *HospBranchServiceStruct {
	return &HospBranchServiceStruct{
		HospitalBranchService: repo,
	}
}

func (ds *HospBranchServiceStruct) CreateHospBranchService(Hospital models.HospitalBranch) error {
	return ds.HospitalBranchService.Create(Hospital)
}

// CreateBranchesService handles the bulk creation of hospital branches
func (hbs *HospBranchServiceStruct) CreateBranchesService(branches []models.HospitalBranch) error {
	return hbs.HospitalBranchService.BulkCreate(branches)
}
func (ds *HospBranchServiceStruct) GetAllHospBranchsService(i int) ([]models.HospitalBranch, error) {
	return ds.HospitalBranchService.GetAll(i)
}

func (ds *HospBranchServiceStruct) GetOneHospBranchService(h models.HospitalBranch) (models.HospitalBranch, error) {
	return ds.HospitalBranchService.GetOne(h)
}
func (ds *HospBranchServiceStruct) UpdateHospBranchService(doctor models.HospitalBranch) error {
	return ds.HospitalBranchService.Modify(doctor)
}

// .................................................................................
// HospitalDoctorServiceStruct provides services related to HospitalDoctor
type HospitalDoctorService struct {
	HospitalDocService repos.HospitalDoctorRepo
}

func NewHospitalDoctorService(repo repos.HospitalDoctorRepo) *HospitalDoctorService {
	return &HospitalDoctorService{
		HospitalDocService: repo,
	}
}

func (hds *HospitalDoctorService) CreateHospitalDoctorService(hospitalDoctors []models.HospitalDoctor) error {
	return hds.HospitalDocService.BulkCreate(hospitalDoctors)
}
func (hds *HospitalDoctorService) GetAllHospitalDoctorsService(id int) ([]models.HospitalDoctor, error) {
	return hds.HospitalDocService.GetAll(id)
}

func (hds *HospitalDoctorService) GetOneHospitalDoctorService(id int) (models.HospitalDoctor, error) {
	return hds.HospitalDocService.GetOne(id)
}
