package handlers

import (
	"fmt"
	"net/http"
	masters "repogin/internal/models/masters"
	models "repogin/internal/models/masters"

	"github.com/gin-gonic/gin"
)

// # department routes
// CreateDepartmentRoute - route to create a department
func (ph *MainHandlers) CreateDepartmentRoute(c *gin.Context) {
	var department models.Department
	if err := c.Bind(&department); err != nil {
		fmt.Println("ERROR : CreateDepartmentRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ph.DepartmentService.CreateDeptService(department); err != nil {
		fmt.Println("ERROR : CreateDepartmentRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Department created successfully"})
}

// UpdateDepartmentRoute - route to update a department
func (ph *MainHandlers) UpdateDepartmentRoute(c *gin.Context) {
	var department models.Department
	if err := c.Bind(&department); err != nil {
		fmt.Println("ERROR : UpdateDepartmentRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ph.DepartmentService.UpdateDeptService(department); err != nil {
		fmt.Println("ERROR : UpdateDepartmentRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Department updated successfully"})
}

// GetAllDepartmentsRoute - route to get all departments
func (ph *MainHandlers) GetAllDepartmentsRoute(c *gin.Context) {
	departments, err := ph.DepartmentService.GetAllDeptsService()
	if err != nil {
		fmt.Println("ERROR : GetAllDepartmentsRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, departments)
}

// GetOneDepartmentRoute - route to get a single department by ID
func (ph *MainHandlers) GetOneDepartmentRoute(c *gin.Context) {
	var department models.Department
	if err := c.Bind(&department); err != nil {
		fmt.Println("ERROR : GetOneDepartmentRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	departmentDetails, err := ph.DepartmentService.GetDeptService(department)
	if err != nil {
		fmt.Println("ERROR : GetOneDepartmentRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, departmentDetails)
}

// DeleteDepartmentRoute - route to delete a department by ID
func (ph *MainHandlers) DeleteDepartmentRoute(c *gin.Context) {
	var department models.Department
	if err := c.Bind(&department); err != nil {
		fmt.Println("ERROR : DeleteDepartmentRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ph.DepartmentService.DeleteDeptService(department); err != nil {
		fmt.Println("ERROR : DeleteDepartmentRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Department deleted successfully"})
}

// ..................................................................................................
// CreateServiceRoute - route to create a new service
func (ph *MainHandlers) CreateServiceRoute(c *gin.Context) {
	var service models.Service
	if err := c.Bind(&service); err != nil {
		fmt.Println("ERROR : CreateServiceRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ph.LabServiceService.CreateService(service); err != nil {
		fmt.Println("ERROR : CreateServiceRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service created successfully"})
}
func (ph *MainHandlers) CreateBulkServicesRoute(c *gin.Context) {
	var labs []models.Service
	err := c.Bind(&labs)
	if err != nil {
		fmt.Println("ERROR : CreateBulkServicesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.LabServiceService.BulkCreateServices(labs)
	if err != nil {
		fmt.Println("ERROR : CreateBulkServicesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "services created successfully"})
}

// UpdateServiceRoute - route to update a service
func (ph *MainHandlers) UpdateServiceRoute(c *gin.Context) {
	var service models.Service
	if err := c.Bind(&service); err != nil {
		fmt.Println("ERROR : UpdateServiceRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ph.LabServiceService.UpdateService(service); err != nil {
		fmt.Println("ERROR : UpdateServiceRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service updated successfully"})
}

// GetAllServicesRoute - route to get all services
func (ph *MainHandlers) GetAllServicesRoute(c *gin.Context) {
	services, err := ph.LabServiceService.GetAllServices()
	if err != nil {
		fmt.Println("ERROR : GetAllServicesRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

// GetOneServiceRoute - route to get a single service by ID
func (ph *MainHandlers) GetOneServiceRoute(c *gin.Context) {
	var service models.Service
	if err := c.Bind(&service); err != nil {
		fmt.Println("ERROR : GetOneServiceRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	serviceDetails, err := ph.LabServiceService.GetServiceByID(service.ServiceID)
	if err != nil {
		fmt.Println("ERROR : GetOneServiceRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, serviceDetails)
}

// DeleteServiceRoute - route to delete a service by ID
func (ph *MainHandlers) DeleteServiceRoute(c *gin.Context) {
	var service models.Service
	if err := c.Bind(&service); err != nil {
		fmt.Println("ERROR : DeleteServiceRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ph.LabServiceService.DeleteService(service.ServiceID); err != nil {
		fmt.Println("ERROR : DeleteServiceRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}

// ..................................................................................................
func (ph *MainHandlers) CreatePatientRoute(c *gin.Context) {
	var patient models.Patient
	if err := c.Bind(&patient); err != nil {
		fmt.Println("ERROR : CreatePatientRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if patient.CityID == 0 && patient.StateID == 0 && patient.CountryID == 0 {
		fmt.Println("ERROR : CreatePatientRoute")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid CountryId/CityId,StateId"})
		return
	}
	if err := ph.PatientService.CreatePatient(patient); err != nil {
		fmt.Println("ERROR : CreatePatientRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Patient created successfully"})
}

func (ph *MainHandlers) CreateBulkPatientsRoute(c *gin.Context) {
	var pats []models.Patient
	err := c.Bind(&pats)
	if err != nil {
		fmt.Println("ERROR : CreateBulkServicesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	err = ph.PatientService.BulkCreatePatients(pats)
	if err != nil {
		fmt.Println("ERROR : CreateBulkServicesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "services created successfully"})
}

// UpdateServiceRoute - route to update a service
func (ph *MainHandlers) UpdatePatientRoute(c *gin.Context) {
	var p models.Patient
	if err := c.Bind(&p); err != nil {
		fmt.Println("ERROR : UpdatePatientRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ph.PatientService.UpdatePatient(p); err != nil {
		fmt.Println("ERROR : UpdatePatientRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "UpdatePatientRoute updated successfully"})
}

// GetAllServicesRoute - route to get all services
func (ph *MainHandlers) GetAllPatientsRoute(c *gin.Context) {
	services, err := ph.PatientService.GetAllPatients()
	if err != nil {
		fmt.Println("ERROR : GetAllServicesRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

// GetOneServiceRoute - route to get a single service by ID
func (ph *MainHandlers) GetOnePatintRoute(c *gin.Context) {
	var service models.Patient
	if err := c.Bind(&service); err != nil {
		fmt.Println("ERROR : GetOneServiceRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	serviceDetails, err := ph.PatientService.GetPatient(service)
	if err != nil {
		fmt.Println("ERROR : GetOneServiceRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, serviceDetails)
}

// DeleteServiceRoute - route to delete a service by ID
func (ph *MainHandlers) DeletePatientRoute(c *gin.Context) {
	var service models.Patient
	if err := c.Bind(&service); err != nil {
		fmt.Println("ERROR : DeletePatienteRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ph.PatientService.DeletePatient(service); err != nil {
		fmt.Println("ERROR : DeleteServiceRoute", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}

// ..................................................................................................
func (mh *MainHandlers) CreateVisitTestsRoute(c *gin.Context) {
	var visitDetails masters.VisitTest
	Error := c.Bind(&visitDetails)
	if Error != nil {
		fmt.Println("ERROR : CreateVisitTestsRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.VisitTestService.CreateVisitTest(visitDetails)
	if Err != nil {
		fmt.Println("ERROR : CreateVisitTestsRoute", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Branch created successfully"})
	return
}

func (mh *MainHandlers) UpdateVisitTestsRoute(c *gin.Context) {
	var vtDetails models.VisitTest
	Error := c.Bind(&vtDetails)
	if Error != nil {
		fmt.Println("ERROR : UpdateVisitTestsRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.VisitTestService.UpdateVisitTest(vtDetails)
	if Err != nil {
		fmt.Println("ERROR : UpdateVisitTestsRoute", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "VISIT TEsT updated successfully"})
	return
}

func (mh *MainHandlers) GetAllVisitTestsRoute(c *gin.Context) {
	vs, err := mh.VisitTestService.GetAllVisitTests()
	if err != nil {
		fmt.Println("ERROR : GetAllBranchesRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vs)
	return
}

// ...................................................................................................

func (mh *MainHandlers) CreateVisitRoute(c *gin.Context) {
	var vd models.Visit
	Error := c.Bind(&vd)
	if Error != nil {
		fmt.Println("ERROR : CreateBranchRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Visitservice.CreateVisit(vd)
	if Err != nil {
		fmt.Println("ERROR : CreateBranch", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Branch created successfully"})
	return
}

func (mh *MainHandlers) GetAllVisitsRoute(c *gin.Context) {

	Vists, Err := mh.Visitservice.GetAllVisits()
	if Err != nil {
		fmt.Println("ERROR : UpdateBranch", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, Vists)
	return
}

// func (mh *MainHandlers) GetOneBranchRoute(c *gin.Context) {
// 	var branch models.Branch
// 	bindError := c.Bind(&branch)
// 	if bindError != nil {
// 		fmt.Println("ERROR : GetOneBranchRoute", bindError)
// 	}
// 	branchDetails, err := mh.Branchservice.GetOneBranchService(branch)
// 	if err != nil {
// 		fmt.Println("ERROR : GetOneBranchRoute", err)
// 		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, branchDetails)
// 	return
// }

// func (mh *MainHandlers) DeleteBranchRoute(c *gin.Context) {
// 	var branch models.Branch
// 	bindError := c.Bind(&branch)
// 	if bindError != nil {
// 		fmt.Println("ERROR : DeleteBranchRoute", bindError)
// 	}
// 	err := mh.Branchservice.DeleteOneBranchService(branch)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, map[string]string{"message": "Branch deleted successfully"})
// 	return
// }

// ................................................................................
/*func (mh *MainHandlers) CreateTestRoute(c *gin.Context) {
	var vd models.Test
	Error := c.Bind(&vd)
	if Error != nil {
		fmt.Println("ERROR : CreateTestRoute", Error)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Error.Error()})
		return
	}
	Err := mh.Testservice.CreateTest(vd)
	if Err != nil {
		fmt.Println("ERROR : CreateTestRoute", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "Branch created successfully"})
	return
}

func (mh *MainHandlers) GetAllTestsRoute(c *gin.Context) {

	Vists, Err := mh.Testservice.GetAllTests()
	if Err != nil {
		fmt.Println("ERROR : GetAllTestsRoute", Err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": Err.Error()})
		return
	}
	c.JSON(http.StatusOK, Vists)
	return
}
*/
