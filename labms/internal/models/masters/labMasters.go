package masters

import "time"

type Lab struct {
	LabID        int    `json:"lab_id" db:"lab_id"`
	LabName      string `json:"lab_name" db:"lab_name"`
	LabCode      string `json:"lab_code" db:"lab_code"`
	ValidityDate string `json:"validity_date" db:"validity_date"`
	CreatedOn    string `json:"created_on" db:"created_on"`
	CreatedBy    string `json:"created_by" db:"created_by"`
}
type Branch struct {
	BranchID   int    `json:"branch_id" db:"branch_id"`
	BranchName string `json:"branch_name" db:"branch_name"`
	LabID      int    `json:"lab_id" db:"lab_id"`
	BranchCode string `json:"branch_code" db:"branch_code"`
	Address    string `json:"address" db:"address"`
	CityID     int    `json:"city_id" db:"city_id"`
	LabName    string `json:"lab_name" db:"lab_name"`
	LabCode    string `json:"lab_code" db:"lab_code"`
}
type BranchDepts struct {
	BranchID    int          `json:"branch_id" db:"branch_id"`
	BranchName  string       `json:"branch_name" db:"branch_name"`
	LabID       int          `json:"lab_id" db:"lab_id"`
	BranchCode  string       `json:"branch_code" db:"branch_code"`
	Address     string       `json:"address" db:"address"`
	CityID      int          `json:"city_id" db:"city_id"`
	LabName     string       `json:"lab_name" db:"lab_name"`
	LabCode     string       `json:"lab_code" db:"lab_code"`
	Departments []Department `json:"departments"`
}
type Department struct {
	DepartmentID   int    `json:"department_id" db:"department_id"`
	BranchID       int    `json:"branch_id" db:"branch_id"`
	DepartmentName string `json:"department_name" db:"department_name"`
	Description    string `json:"description" db:"description"`
	LabID          int    `json:"lab_id" db:"lab_id"` //not required but added just for normalization
}

type DeptServices struct {
	DepartmentID   int     `json:"department_id" db:"department_id"`
	BranchID       int     `json:"branch_id" db:"branch_id"`
	DepartmentName string  `json:"department_name" db:"department_name"`
	Description    string  `json:"description" db:"description"`
	LabID          int     `json:"lab_id" db:"lab_id"` //not required but added just for normalization
	Services       Service `json:"services" db:"services"`
}

type Service struct {
	ServiceID               int64   `db:"service_id"`
	DepartmentID            int64   `json:"department_id"  db:"department_id"`
	ServiceName             string  `json:"service_name"  db:"service_name"`
	Description             string  `json:"description" db:"description"`
	BasicRate               float64 `json:"basic_rate" db:"basic_rate"`
	DurationMinutes         int     `json:"duration_minutes" db:"duration_minutes"`
	PreparationInstructions string  `json:"preparation_instructions" db:"preparation_instructions"`
}
type Test struct {
	TestID      int     `json:"test_id" db:"test_id"`
	TestName    string  `json:"test_name" db:"test_name"`
	TestType    string  `json:"test_type" db:"test_type"`
	Price       float64 `json:"price" db:"price"`
	Description string  `json:"description" db:"description"`
}
type Visit struct {
	VisitID       int       `json:"visit_id" db:"visit_id"`
	PatientID     int       `json:"patient_id" db:"patient_id"`
	VisitDate     time.Time `json:"visit_date" db:"visit_date"`
	TotalAmount   float64   `json:"total_amount" db:"total_amount"`
	Status        string    `json:"status" db:"status"`                 //status completed , noncompleted
	PaymentStatus string    `json:"payment_status" db:"payment_status"` // done , not done
}
type VisitTest struct {
	VisitTestID int     `json:"visit_test_id" db:"visit_test_id"`
	VisitID     int     `json:"visit_id" db:"visit_id"`
	TestID      int     `json:"test_id" db:"test_id"`
	Price       float64 `json:"price" db:"price"`
}

type Patient struct {
	PatientID        int    `json:"patient_id" db:"patient_id"`
	FirstName        string `json:"first_name" db:"first_name"`
	LastName         string `json:"last_name" db:"last_name"`
	Age              int    `json:"age" db:"age"`
	Gender           string `json:"gender" db:"gender"`
	ContactNumber    string `json:"contact_number" db:"contact_number"`
	Email            string `json:"email" db:"email"`
	Address          string `json:"address" db:"address"`
	PatientCode      string `json:"patient_code" db:"patient_code"`
	UserID           int    `json:"user_id" db:"user_id"`
	MedicalHistory   string `json:"medical_history" db:"medical_history"`
	BloodType        string `json:"blood_type" db:"blood_type"`
	InsuranceDetails string `json:"insurance_details" db:"insurance_details"`
	StateID          int    `json:"state_id" db:"state_id"`
	CountryID        int    `json:"country_id" db:"country_id"`
	CityID           int    `json:"city_id" db:"city_id"`
}

type Report struct {
	ReportID      int       `json:"report_id" db:"report_id"`
	VisitID       int       `json:"visit_id" db:"visit_id"`
	TestID        int       `json:"test_id" db:"test_id"`
	ReportFile    string    `json:"report_file" db:"report_file"`
	ResultStatus  string    `json:"result_status" db:"result_status"`
	GeneratedDate time.Time `json:"generated_date" db:"generated_date"`
}
type Billing struct {
	BillID        int       `json:"bill_id" db:"bill_id"`
	VisitID       int       `json:"visit_id" db:"visit_id"`
	TotalAmount   float64   `json:"total_amount" db:"total_amount"`
	Discount      float64   `json:"discount" db:"discount"`
	Tax           float64   `json:"tax" db:"tax"`
	FinalAmount   float64   `json:"final_amount" db:"final_amount"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	PaymentDate   time.Time `json:"payment_date" db:"payment_date"`
}

type PatientVisits struct {
	PatientID     int       `json:"patient_id" db:"patient_id"`
	FirstName     string    `json:"first_name" db:"first_name"`
	LastName      string    `json:"last_name" db:"last_name"`
	PatientCode   string    `json:"patient_code" db:"patient_code"`
	UserID        int       `json:"user_id" db:"user_id"`
	Price         float64   `json:"price" db:"price"`
	TestName      string    `json:"test_name" db:"test_name"`
	TestType      string    `json:"test_type" db:"test_type"`
	ActualPrice   float64   `json:"actual_price" db:"actual_price"`
	Status        string    `json:"status" db:"status"`
	PaymentStatus string    `json:"payment_status" db:"payment_status"`
	VisitTestID   int       `json:"visit_test_id" db:"visit_test_id"`
	VisitID       int       `json:"visit_id" db:"visit_id"`
	TestID        int       `json:"test_id" db:"test_id"`
	VisitDate     time.Time `json:"visit_date" db:"visit_date"`
	TotalAmount   float64   `json:"total_amount" db:"total_amount"`
}
