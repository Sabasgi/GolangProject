package services

import (
	"errors"
	"fmt"
	masters "repogin/internal/models/masters"
	repos "repogin/internal/repositories/sql"
)

type ServiceService struct {
	repo repos.ServicesRepo
}

// NewServiceService initializes and returns a new ServiceService implementation.
func NewServiceService(repo repos.ServicesRepo) *ServiceService {
	return &ServiceService{repo: repo}
}
func (s *ServiceService) CreateService(service masters.Service) error {
	// Validate basic fields
	if service.ServiceName == "" || service.DepartmentID == 0 || service.BasicRate <= 0 {
		return errors.New("ERROR_INVALID_SERVICE_DATA")
	}

	err := s.repo.Create(service)
	if err != nil {
		fmt.Println("ERROR : ServiceService CreateService ", err)
		return err
	}
	fmt.Println("Service created successfully!")
	return nil
}
func (s *ServiceService) UpdateService(service masters.Service) error {
	// Validate ID and required fields
	if service.ServiceID == 0 || service.ServiceName == "" || service.DepartmentID == 0 || service.BasicRate <= 0 {
		return errors.New("ERROR_INVALID_SERVICE_UPDATE_DATA")
	}

	err := s.repo.Modify(service)
	if err != nil {
		fmt.Println("ERROR : ServiceService UpdateService ", err)
		return err
	}
	fmt.Println("Service updated successfully!")
	return nil
}
func (s *ServiceService) BulkCreateServices(services []masters.Service) error {
	// Check if services slice is not empty
	if len(services) == 0 {
		return errors.New("ERROR_EMPTY_SERVICES_LIST")
	}

	// Validation for each service
	// for _, service := range services {
	// 		if service.ServiceName == "" || service.DepartmentID == 0 || service.BasicRate <= 0 {
	// 				return errors.New("ERROR_INVALID_BULK_SERVICE_DATA")
	// 		}
	// }

	err := s.repo.BulkCreate(services)
	if err != nil {
		fmt.Println("ERROR : ServiceService BulkCreateServices ", err)
		return err
	}
	fmt.Println("Bulk services created successfully!")
	return nil
}
func (s *ServiceService) GetAllServices() ([]masters.Service, error) {
	services, err := s.repo.GetAll()
	if err != nil {
		fmt.Println("ERROR : ServiceService GetAllServices ", err)
		return nil, err
	}

	fmt.Println("Fetched all services successfully!")
	return services, nil
}
func (s *ServiceService) GetServiceByID(serviceID int64) (masters.Service, error) {
	if serviceID == 0 {
		return masters.Service{}, errors.New("ERROR_INVALID_SERVICE_ID")
	}

	service, err := s.repo.GetOne(masters.Service{ServiceID: serviceID})
	if err != nil {
		fmt.Println("ERROR : ServiceService GetServiceByID ", err)
		return service, err
	}

	fmt.Println("Service fetched successfully!")
	return service, nil
}
func (s *ServiceService) DeleteService(serviceID int64) error {
	if serviceID == 0 {
		return errors.New("ERROR_INVALID_SERVICE_ID")
	}

	err := s.repo.DeleteOne(masters.Service{ServiceID: serviceID})
	if err != nil {
		fmt.Println("ERROR : ServiceService DeleteService ", err)
		return err
	}

	fmt.Println("Service deleted successfully!")
	return nil
}

// .....................................................................................

type VisitService struct {
	Repo repos.VisitRepo
}

func NewVisitRepoService(repo repos.VisitRepo) *VisitService {
	return &VisitService{
		Repo: repo,
	}
}

func (s *VisitService) CreateVisit(vt masters.Visit) error {
	return s.Repo.Create(vt)
}

func (s *VisitService) UpdateVisit(vt masters.Visit) error {
	return s.Repo.Modify(vt)
}

func (s *VisitService) GetAllVisits() ([]masters.Visit, error) {
	return s.Repo.GetAll()
}

func (s *VisitService) GetVisitByID(vt masters.Visit) (masters.Visit, error) {
	return s.Repo.GetOne(vt)
}

func (s *VisitService) DeleteTest(vt masters.Visit) error {
	return s.Repo.DeleteOne(vt)
}

// ...................................................................................

type DepartmentService struct {
	DepartmentService repos.DepartmentRepo
}

func NewDepartmentService(repo repos.DepartmentRepo) *DepartmentService {
	return &DepartmentService{
		DepartmentService: repo,
	}
}

func (s *DepartmentService) CreateDeptService(vt masters.Department) error {
	return s.DepartmentService.Create(vt)
}

func (s *DepartmentService) UpdateDeptService(vt masters.Department) error {
	return s.DepartmentService.Modify(vt)
}

func (s *DepartmentService) GetAllDeptsService() ([]masters.Department, error) {
	return s.DepartmentService.GetAll()
}

func (s *DepartmentService) GetDeptService(vt masters.Department) (masters.Department, error) {
	return s.DepartmentService.GetOne(vt)
}

func (s *DepartmentService) DeleteDeptService(vt masters.Department) error {
	return s.DepartmentService.DeleteOne(vt)
}

// .........................................................................................................
type VisitTestService struct {
	Repo repos.VisitTestRepo
}

func NewVisitTestService(repo repos.VisitTestRepo) *VisitTestService {
	return &VisitTestService{
		Repo: repo,
	}
}

func (s *VisitTestService) CreateVisitTest(vt masters.VisitTest) error {
	return s.Repo.Create(vt)
}

func (s *VisitTestService) UpdateVisitTest(vt masters.VisitTest) error {
	return s.Repo.Modify(vt)
}

func (s *VisitTestService) GetAllVisitTests() ([]masters.VisitTest, error) {
	return s.Repo.GetAll()
}

func (s *VisitTestService) GetVisitTestByID(vt masters.VisitTest) (masters.VisitTest, error) {
	return s.Repo.GetOne(vt)
}

func (s *VisitTestService) DeleteVisitTest(vt masters.VisitTest) error {
	return s.Repo.DeleteOne(vt)
}

// .......................................................................................................................

// type TestService struct {
// 	Repo repos.TestRepo
// }

// func NewTestService(repo repos.TestRepo) *TestService {
// 	return &TestService{
// 		Repo: repo,
// 	}
// }

// func (s *TestService) CreateTest(vt masters.Test) error {
// 	return s.Repo.Create(vt)
// }

// func (s *TestService) UpdateTest(vt masters.Test) error {
// 	return s.Repo.Modify(vt)
// }

// func (s *TestService) GetAllTests() ([]masters.Test, error) {
// 	return s.Repo.GetAll()
// }

// func (s *TestService) GetTestByID(vt masters.Test) (masters.Test, error) {
// 	return s.Repo.GetOne(vt)
// }

// func (s *TestService) DeleteTest(vt masters.Test) error {
// 	return s.Repo.DeleteOne(vt)
// }

// .......................................................................................................................

type PatientService struct {
	Repo repos.PatientRepo
}

func NewPatientService(repo repos.PatientRepo) *PatientService {
	return &PatientService{
		Repo: repo,
	}
}

func (s *PatientService) CreatePatient(vt masters.Patient) error {
	return s.Repo.Create(vt)
}

func (s *PatientService) UpdatePatient(vt masters.Patient) error {
	return s.Repo.Modify(vt)
}

func (s *PatientService) GetAllPatients() ([]masters.Patient, error) {
	return s.Repo.GetAll()
}

func (s *PatientService) GetPatient(vt masters.Patient) (masters.Patient, error) {
	return s.Repo.GetOne(vt)
}

func (s *PatientService) DeletePatient(vt masters.Patient) error {
	return s.Repo.DeleteOne(vt)
}
func (s *PatientService) BulkCreatePatients(Pts []masters.Patient) error {
	// Check if services slice is not empty
	if len(Pts) == 0 {
		return errors.New("ERROR_EMPTY_Patients")
	}

	// Validation for each service
	// for _, service := range services {
	// 		if service.ServiceName == "" || service.DepartmentID == 0 || service.BasicRate <= 0 {
	// 				return errors.New("ERROR_INVALID_BULK_SERVICE_DATA")
	// 		}
	// }

	err := s.Repo.BulkCreate(Pts)
	if err != nil {
		fmt.Println("ERROR : Patient BulkCreateServices ", err)
		return err
	}
	fmt.Println("Bulk Patient created successfully!")
	return nil
}

// .......................................................................................................................
