package repositories

import (
	"errors"
	"fmt"
	"os"
	"repogin/internal/db"
	"repogin/internal/models/masters"
	models "repogin/internal/models/masters"
	"strconv"
)

type PatientRepo interface {
	Create(models.Patient) error
	Modify(models.Patient) error
	GetAll() ([]models.Patient, error)
	GetOne(models.Patient) (models.Patient, error)
	DeleteOne(models.Patient) error
	BulkCreate([]models.Patient) error
}
type PatientSQLRepo struct {
	PRepo *db.SQLRepo
}

func NewPatientRepo(sr *db.SQLRepo) *PatientSQLRepo {
	return &PatientSQLRepo{
		PRepo: sr,
	}
}

// Create a new patient record
func (pr *PatientSQLRepo) Create(patient models.Patient) error {
	l, _ := strconv.Atoi(os.Getenv("codelength"))
	c, e := generateRandomString(l)
	if e != nil {
		fmt.Println("Hospital inserted successfully! with sql id - ")
		return errors.New("ERRORCDE_CODE_GENERATION_ERR")
	}
	patient.PatientCode = c
	query := `INSERT INTO Patient (
			first_name, last_name, age, gender, contact_number, email, address,
			patient_code, user_id, medical_history, blood_type, insurance_details,
			state_id, country_id, city_id
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, rerr := pr.PRepo.Session.Exec(query, patient.FirstName, patient.LastName, patient.Age, patient.Gender,
		patient.ContactNumber, patient.Email, patient.Address, patient.PatientCode,
		patient.UserID, patient.MedicalHistory, patient.BloodType, patient.InsuranceDetails,
		patient.StateID, patient.CountryID, patient.CityID)

	if rerr != nil {
		fmt.Println("ERROR : PatientSQLRepo Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("Patient inserted successfully! with SQL ID - ", inc)
		return nil
	}

	return errors.New("ERROR_PATIENT_NOT_INSERTED")
}

/*
UPDATE Patient SET
			first_name = ?, last_name = ?, age = ?, gender = ?, contact_number = ?,
			email = ?, address = ?, patient_code = ?, user_id = ?, medical_history = ?,
			blood_type = ?, insurance_details = ?, state_id = ?, country_id = ?, city_id = ?
		WHERE patient_id = ?

		Exec(query, patient.FirstName, patient.LastName, patient.Age, patient.Gender,
		patient.ContactNumber, patient.Email, patient.Address, patient.PatientCode,
		patient.UserID, patient.MedicalHistory, patient.BloodType, patient.InsuranceDetails,
		patient.StateID, patient.CountryID, patient.CityID, patient.PatientID,
	)
*/
// Update a patient record
func (pr *PatientSQLRepo) Modify(p models.Patient) error {

	query := `UPDATE patient
						SET first_name = ?, last_name = ?, age = ?, gender = ?, contact_number = ?, email = ?, address = ?, user_id = ?, medical_history = ?, blood_type = ?, insurance_details = ?
						WHERE patient_id = ?`

	res, rerr := pr.PRepo.Session.Exec(query, p.FirstName, p.LastName, p.Age, p.Gender, p.ContactNumber, p.Email, p.Address, p.PatientCode, p.UserID, p.MedicalHistory, p.BloodType, p.InsuranceDetails, p.PatientID)

	if rerr != nil {
		fmt.Println("ERROR : PatientSQLRepo Modify ", rerr)
		return rerr
	}

	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Patient updated successfully!")
		return nil
	}

	return errors.New("ERROR_PATIENT_NOT_UPDATED")

}

// Create inserts a array of city into the City table
func (pr *PatientSQLRepo) BulkCreate(pts []models.Patient) error {
	if len(pts) > 0 {
		query := `INSERT INTO patient
						(first_name, last_name, age, gender, contact_number, email, address, patient_code, user_id, medical_history, blood_type, insurance_details)
						VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		for _, p := range pts {
			res, rerr := pr.PRepo.Session.Exec(query, p.FirstName, p.LastName, p.Age, p.Gender, p.ContactNumber, p.Email, p.Address, p.PatientCode, p.UserID, p.MedicalHistory, p.BloodType, p.InsuranceDetails)
			if rerr != nil {
				fmt.Println("ERROR : PatientREPO Create ", rerr)
				return rerr
			}
			inc, _ := res.LastInsertId()
			rac, _ := res.RowsAffected()
			if inc > 0 || rac > 0 {
				fmt.Println("Patient inserted successfully! with sql id - ", inc)
			} else {
				fmt.Println("Patient failed to insert - ", p.FirstName, " - ", p.LastName, " -", p.Email)
				//LOGS Implementation remaining
			}
		}
	}
	return errors.New("ERROR_CITIES_LENGTH_IS_ZERO")
}

// Retrieve all patients
func (pr *PatientSQLRepo) GetAll() ([]models.Patient, error) {
	var patients []models.Patient
	query := `SELECT patient_id, first_name, last_name, age, gender, contact_number, email, address, patient_code, user_id, medical_history, blood_type, insurance_details
						FROM patient`

	rowsCount, rerr := pr.PRepo.Session.SelectBySql(query).Load(&patients)
	if rerr != nil {
		fmt.Println("ERROR : PatientSQLRepo GetAll ", rerr)
		return nil, rerr
	}

	if rowsCount > 0 && len(patients) > 0 {
		return patients, nil
	}

	return patients, errors.New("ERROR_PATIENTS_NOT_FOUND")
}

// Retrieve a single patient by ID
func (pr *PatientSQLRepo) GetOne(p models.Patient) (models.Patient, error) {
	query := `SELECT patient_id, first_name, last_name, age, gender, contact_number, email, address, patient_code, user_id, medical_history, blood_type, insurance_details
						FROM patient
						WHERE patient_id = ?`

	var patient models.Patient
	err := pr.PRepo.Session.SelectBySql(query).LoadOne(&patient)
	if err != nil {
		fmt.Println("ERROR : PatientSQLRepo GetOne ", err)
		return patient, err
	}

	return patient, nil
}

// Delete a patient record
func (pr *PatientSQLRepo) DeleteOne(p models.Patient) error {
	query := `DELETE FROM patient WHERE patient_id = ?`

	res, rerr := pr.PRepo.Session.Exec(query, p.PatientID)
	if rerr != nil {
		fmt.Println("ERROR : PatientSQLRepo DeleteOne ", rerr)
		return rerr
	}

	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Patient deleted successfully!")
		return nil
	}

	return errors.New("ERROR_PATIENT_NOT_DELETED")
}

// get patine details with visits tests
func (pr *PatientSQLRepo) GetPatientDetailsWithVisitsAndTests(patientID int) ([]models.PatientVisits, error) {
	q := `SELECT
			p.patient_id, p.first_name, p.last_name, p.age, p.gender, p.contact_number,
			p.email, p.address, p.patient_code, p.user_id, p.medical_history, p.blood_type, p.insurance_details,
			v.visit_id, v.visit_date, v.total_amount, v.status, v.payment_status,
			t.test_id, t.test_name, t.test_type, t.price AS test_price, t.description,
			vt.vt_id AS visit_test_id, vt.price AS price
		FROM
			Patient p
			LEFT JOIN Visit v ON p.patient_id = v.patient_id
			LEFT JOIN visits_tests vt ON v.visit_id = vt.visit_id
			LEFT JOIN Test t ON vt.test_id = t.test_id
		WHERE
			p.patient_id = ?`
	var patients []models.PatientVisits

	rowsCount, rerr := pr.PRepo.Session.SelectBySql(q).Load(&patients)
	if rerr != nil {
		fmt.Println("ERROR : PatientSQLRepo GetAll ", rerr)
		return nil, rerr
	}
	if rowsCount > 0 && len(patients) > 0 {
		return patients, nil
	}
	return patients, errors.New("Patient_Details_not_Found")
}

// ................................................................................

type DepartmentRepo interface {
	Create(models.Department) error
	Modify(models.Department) error
	GetAll() ([]models.Department, error)
	GetOne(models.Department) (models.Department, error)
	DeleteOne(models.Department) error
}
type DepartmentSQLRepo struct {
	DRepo *db.SQLRepo
}

func NewDepartmentRepo(sr *db.SQLRepo) *DepartmentSQLRepo {
	return &DepartmentSQLRepo{
		DRepo: sr,
	}
}

func (dr *DepartmentSQLRepo) Create(d models.Department) error {
	query := "INSERT INTO Department (branch_id, department_name, n,lab_id) VALUES (?, ?, ?,?)"
	res, rerr := dr.DRepo.Session.Exec(query, d.BranchID, d.DepartmentName, d.Description, d.LabID)
	if rerr != nil {
		fmt.Println("ERROR : DepartmentSQLRepo Create ", rerr)
		return rerr
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("Department inserted successfully with SQL ID - ", inc)
		return nil
	}
	return errors.New("ERROR_DEPARTMENT_NOT_INSERTED")
}

func (dr *DepartmentSQLRepo) Modify(d models.Department) error {
	query := "UPDATE Department SET branch_id = ?, department_name = ?, description = ? WHERE department_id = ?"
	res, rerr := dr.DRepo.Session.Exec(query, d.BranchID, d.DepartmentName, d.Description, d.DepartmentID)
	if rerr != nil {
		fmt.Println("ERROR : DepartmentSQLRepo Modify ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Department updated successfully!")
		return nil
	}
	return errors.New("ERROR_DEPARTMENT_NOT_UPDATED")
}

func (dr *DepartmentSQLRepo) GetAll() ([]models.Department, error) {
	var departments []models.Department
	query := "SELECT department_id, branch_id, department_name, description FROM Department"
	rows, rerr := dr.DRepo.Session.Query(query)
	if rerr != nil {
		fmt.Println("ERROR : DepartmentSQLRepo GetAll ", rerr)
		return nil, rerr
	}
	defer rows.Close()

	for rows.Next() {
		var department models.Department
		if err := rows.Scan(&department.DepartmentID, &department.BranchID, &department.DepartmentName, &department.Description); err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}

	if len(departments) > 0 {
		return departments, nil
	}
	return nil, errors.New("ERROR_DEPARTMENTS_NOT_FOUND")
}

func (dr *DepartmentSQLRepo) GetOne(d models.Department) (models.Department, error) {
	query := "SELECT department_id, branch_id, department_name, description FROM Department WHERE department_id = ?"
	var department models.Department
	err := dr.DRepo.Session.QueryRow(query, d.DepartmentID).Scan(&department.DepartmentID, &department.BranchID, &department.DepartmentName, &department.Description)
	if err != nil {
		fmt.Println("ERROR : DepartmentSQLRepo GetOne ", err)
		return department, err
	}
	return department, nil
}

func (dr *DepartmentSQLRepo) DeleteOne(d models.Department) error {
	query := "DELETE FROM Departments WHERE department_id = ?"
	res, rerr := dr.DRepo.Session.Exec(query, d.DepartmentID)
	if rerr != nil {
		fmt.Println("ERROR : DepartmentSQLRepo DeleteOne ", rerr)
		return rerr
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Department deleted successfully!")
		return nil
	}
	return errors.New("ERROR_DEPARTMENT_NOT_DELETED")
}

// ......................................................................................................

type ServicesRepo interface {
	Create(models.Service) error
	Modify(models.Service) error
	BulkCreate(services []models.Service) error
	GetAll() ([]models.Service, error)
	GetOne(models.Service) (models.Service, error)
	DeleteOne(models.Service) error
}
type ServicesSQLRepo struct {
	CRepo *db.SQLRepo
}

func NewServicesRepo(sr *db.SQLRepo) *ServicesSQLRepo {
	return &ServicesSQLRepo{
		CRepo: sr,
	}
}
func (sr *ServicesSQLRepo) Create(s models.Service) error {
	query := "INSERT INTO Service (department_id, service_name, description, basic_rate, duration_minutes, preparation_instructions) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := sr.CRepo.Session.Exec(query, s.DepartmentID, s.ServiceName, s.Description, s.BasicRate, s.DurationMinutes, s.PreparationInstructions)
	if err != nil {
		fmt.Println("ERROR : ServicesSQLRepo Create ", err)
		return err
	}
	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("Service inserted successfully! with SQL ID - ", inc)
		return nil
	}
	return errors.New("ERROR_SERVICE_NOT_INSERTED")
}
func (sr *ServicesSQLRepo) Modify(s models.Service) error {
	query := "UPDATE Services SET department_id = ?, service_name = ?, description = ?, basic_rate = ?, duration_minutes = ?, preparation_instructions = ? WHERE service_id = ?"
	res, err := sr.CRepo.Session.Exec(query, s.DepartmentID, s.ServiceName, s.Description, s.BasicRate, s.DurationMinutes, s.PreparationInstructions, s.ServiceID)
	if err != nil {
		fmt.Println("ERROR : ServicesSQLRepo Modify ", err)
		return err
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Service updated successfully!")
		return nil
	}
	return errors.New("ERROR_SERVICE_NOT_UPDATED")
}
func (sr *ServicesSQLRepo) BulkCreate(services []models.Service) error {
	if len(services) > 0 {
		query := "INSERT INTO Services (department_id, service_name, description, basic_rate, duration_minutes, preparation_instructions) VALUES (?, ?, ?, ?, ?, ?)"
		for _, service := range services {
			res, err := sr.CRepo.Session.Exec(query, service.DepartmentID, service.ServiceName, service.Description, service.BasicRate, service.DurationMinutes, service.PreparationInstructions)
			if err != nil {
				fmt.Println("ERROR : ServicesSQLRepo BulkCreate ", err)
				return err
			}
			inc, _ := res.LastInsertId()
			if inc > 0 {
				fmt.Println("Service inserted successfully! with SQL ID - ", inc)
			} else {
				fmt.Println("Failed to insert service - ", service.ServiceName)
			}
		}
		return nil
	}
	return errors.New("ERROR_SERVICES_LENGTH_IS_ZERO")
}
func (sr *ServicesSQLRepo) GetAll() ([]models.Service, error) {
	var services []models.Service
	query := "SELECT service_id, department_id, service_name, description, basic_rate, duration_minutes, preparation_instructions FROM Services"
	rowsCount, err := sr.CRepo.Session.SelectBySql(query).Load(&services)
	if err != nil {
		fmt.Println("ERROR : ServicesSQLRepo GetAll ", err)
		return nil, err
	}

	if rowsCount > 0 && len(services) > 0 {
		fmt.Println("SUCCESS : ServicesSQLRepo Len(services) ", len(services))
		return services, nil
	}
	return services, errors.New("NOT_FOUND")
}
func (sr *ServicesSQLRepo) DeleteOne(s models.Service) error {
	query := "DELETE FROM Services WHERE service_id = ?"
	res, err := sr.CRepo.Session.Exec(query, s.ServiceID)
	if err != nil {
		fmt.Println("ERROR : ServicesSQLRepo DeleteOne ", err)
		return err
	}
	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Service deleted successfully!")
		return nil
	}
	return errors.New("ERROR_SERVICE_NOT_DELETED")
}
func (sr *ServicesSQLRepo) GetOne(s models.Service) (models.Service, error) {
	query := "SELECT service_id, department_id, service_name, description, basic_rate, duration_minutes, preparation_instructions FROM Services WHERE service_id = ?"
	var service models.Service
	err := sr.CRepo.Session.SelectBySql(query, s.ServiceID).LoadOne(&service)
	if err != nil {
		fmt.Println("ERROR : ServicesSQLRepo GetOne ", err)
		return service, err
	}
	return service, nil
}
func (sr *ServicesSQLRepo) GetDeptsOfBranch(id int) (masters.Department, error) {
	var depts masters.Department
	q := "SELECT d.`department_name`,d.`department_id`,b.`branch_id` FROM department d 	INNER JOIN department b ON d.branch_id = b.branch_id 	WHERE d.branch_id = 1 ;"
	c, E := sr.CRepo.Session.SelectBySql(q, id).Load(&depts)
	if E != nil {
		fmt.Println("ERROR : ServicesSQLRepo GetOne ", E)
		return depts, E
	}
	if c > 0 {
		fmt.Println("COUNT : e ", c)
		return depts, nil
	}
	return depts, errors.New("ERRORCODE_DEPT_NOT_FOUND")
}

// ......................................................................................................

type VisitRepo interface {
	Create(models.Visit) error
	Modify(models.Visit) error
	GetAll() ([]models.Visit, error)
	GetOne(models.Visit) (models.Visit, error)
	DeleteOne(models.Visit) error
}
type VisitSQLRepo struct {
	VRepo *db.SQLRepo
}

func NewVisitRepo(sr *db.SQLRepo) *VisitSQLRepo {
	return &VisitSQLRepo{
		VRepo: sr,
	}
}

// Create a new visit record
func (vr *VisitSQLRepo) Create(v models.Visit) error {
	query := `INSERT INTO visit (patient_id, visit_date, total_amount, status, payment_status)
						VALUES (?, ?, ?, ?, ?)`
	res, rerr := vr.VRepo.Session.Exec(query, v.PatientID, v.VisitDate, v.TotalAmount, v.Status, v.PaymentStatus)

	if rerr != nil {
		fmt.Println("ERROR : VisitSQLRepo Create ", rerr)
		return rerr
	}

	inc, _ := res.LastInsertId()
	rac, _ := res.RowsAffected()
	if inc > 0 || rac > 0 {
		fmt.Println("Visit inserted successfully! with SQL ID - ", inc)
		return nil
	}

	return errors.New("ERROR_VISIT_NOT_INSERTED")
}

// Update a visit record
func (vr *VisitSQLRepo) Modify(v models.Visit) error {
	query := `UPDATE visit
						SET patient_id = ?, visit_date = ?, total_amount = ?, status = ?, payment_status = ?
						WHERE visit_id = ?`

	res, rerr := vr.VRepo.Session.Exec(query, v.PatientID, v.VisitDate, v.TotalAmount, v.Status, v.PaymentStatus, v.VisitID)

	if rerr != nil {
		fmt.Println("ERROR : VisitSQLRepo Modify ", rerr)
		return rerr
	}

	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Visit updated successfully!")
		return nil
	}

	return errors.New("ERROR_VISIT_NOT_UPDATED")
}

// Retrieve all visits
func (vr *VisitSQLRepo) GetAll() ([]models.Visit, error) {
	var visits []models.Visit
	query := `SELECT visit_id, patient_id, visit_date, total_amount, status, payment_status
						FROM visit`

	rowsCount, rerr := vr.VRepo.Session.SelectBySql(query).Load(&visits)
	if rerr != nil {
		fmt.Println("ERROR : VisitSQLRepo GetAll ", rerr)
		return nil, rerr
	}

	if rowsCount > 0 && len(visits) > 0 {
		return visits, nil
	}

	return visits, errors.New("ERROR_VISITS_NOT_FOUND")
}

// Retrieve a single visit by ID
func (vr *VisitSQLRepo) GetOne(v models.Visit) (models.Visit, error) {
	query := `SELECT visit_id, patient_id, visit_date, total_amount, status, payment_status
						FROM visit
						WHERE visit_id = ?`

	var visit models.Visit
	err := vr.VRepo.Session.SelectBySql(query).LoadOne(&visit)
	if err != nil {
		fmt.Println("ERROR : VisitSQLRepo GetOne ", err)
		return visit, err
	}

	return visit, nil
}

// Delete a visit record
func (vr *VisitSQLRepo) DeleteOne(v models.Visit) error {
	query := `DELETE FROM visit_master WHERE visit_id = ?`

	res, rerr := vr.VRepo.Session.Exec(query, v.VisitID)
	if rerr != nil {
		fmt.Println("ERROR : VisitSQLRepo DeleteOne ", rerr)
		return rerr
	}

	rac, _ := res.RowsAffected()
	if rac > 0 {
		fmt.Println("Visit deleted successfully!")
		return nil
	}

	return errors.New("ERROR_VISIT_NOT_DELETED")
}

// .........................................................................................

type VisitTestRepo interface {
	Create(models.VisitTest) error
	Modify(models.VisitTest) error
	GetAll() ([]models.VisitTest, error)
	GetOne(models.VisitTest) (models.VisitTest, error)
	DeleteOne(models.VisitTest) error
}
type VisitTestSQLRepo struct {
	DB *db.SQLRepo
}

func NewVisitTestRepo(db *db.SQLRepo) *VisitTestSQLRepo {
	return &VisitTestSQLRepo{
		DB: db,
	}
}

func (repo *VisitTestSQLRepo) Create(vt models.VisitTest) error {
	query := "INSERT INTO visits_tests (visit_id, test_id, price) VALUES (?, ?, ?)"
	_, err := repo.DB.Session.Exec(query, vt.VisitID, vt.TestID, vt.Price)
	if err != nil {
		fmt.Println("ERROR: VisitTestSQLRepo Create", err)
		return err
	}
	fmt.Println("Visit Test inserted successfully!")
	return nil
}

func (repo *VisitTestSQLRepo) Modify(vt models.VisitTest) error {
	query := "UPDATE visits_tests SET visit_id = ?, test_id = ?, price = ? WHERE vt_id = ?"
	_, err := repo.DB.Session.Exec(query, vt.VisitID, vt.TestID, vt.Price, vt.VisitID)
	if err != nil {
		fmt.Println("ERROR: VisitTestSQLRepo Modify", err)
		return err
	}
	fmt.Println("Visit Test updated successfully!")
	return nil
}

func (repo *VisitTestSQLRepo) GetAll() ([]models.VisitTest, error) {
	query := "SELECT vt_id, visit_id, test_id, price FROM visits_tests"
	var visitTests []models.VisitTest
	rows, err := repo.DB.Session.SelectBySql(query).Load(&visitTests)
	if err != nil {
		fmt.Println("ERROR: VisitTestSQLRepo GetAll", err)
		return nil, err
	}
	if rows > 0 {
		fmt.Println("VisitTestSQLRepo Found", rows)
		return visitTests, nil
	}
	fmt.Println("ERROR: VisitTestSQLRepo GetAll", err)
	return nil, err
}

func (repo *VisitTestSQLRepo) GetOne(vt models.VisitTest) (models.VisitTest, error) {
	query := "SELECT vt_id, visit_id, test_id, price FROM visits_tests WHERE vt_id = ?"
	var visitTest models.VisitTest
	e := repo.DB.Session.SelectBySql(query, vt.VisitID).LoadOne(&visitTest)
	if e != nil {
		fmt.Println("eOR: VisitTestSQLRepo GetOne", e)
		return visitTest, e
	}
	return visitTest, nil
}

func (repo *VisitTestSQLRepo) DeleteOne(vt models.VisitTest) error {
	query := "DELETE FROM visits_tests WHERE vt_id = ?"
	_, err := repo.DB.Session.Exec(query, vt.VisitID)
	if err != nil {
		fmt.Println("ERROR: VisitTestSQLRepo DeleteOne", err)
		return err
	}
	fmt.Println("Visit Test deleted successfully!")
	return nil
}
