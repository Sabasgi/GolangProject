package masters

type Country struct {
	ID             int    `json:"id" db:"country_id"`
	Code           string `json:"code" db:"country_code"` //required_in_payload_
	Name           string `json:"name" db:"country_name"` //required_in_payload
	GST_Percentage int    `json:"gst_percentage" db:"gst_percentage"`
}

type State struct {
	StateID   int    `json:"state_id" db:"state_id"`
	StateName string `json:"state_name" db:"state_name"` //required_in_payload_
	CountryID int    `json:"country_id" db:"country_id"` //required_in_payload_
}

type City struct {
	CityID         int    `json:"city_id" db:"city_id"`
	CityName       string `json:"city_name" db:"city_name"` //required_in_payload_
	StateID        int    `json:"state_id" db:"state_id"`   //required_in_payload_
	GST_Percentage int    `json:"gst_percentage" db:"gst_percentage"`
}

type Role struct {
	RoleID      int    `json:"role_id" db:"role_id"`
	RoleName    string `json:"role_name" db:"role_name"`
	Description string `json:"description" db:"description"`
}
type Userr struct {
	UserID      int    `json:"user_id" db:"user_id"`
	LabID       int    `json:"lab_id" db:"user_id"`
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	Role        string `json:"role" db:"role"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	// Address     string `json:"address" db:"address"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Username  string `json:"username" db:"username"`
}
type Doctor struct {
	DoctorID          int    `json:"doctor_id" db:"doctor_id"`
	UserID            int    `json:"user_id" db:"user_id"`
	Specialization    string `json:"specialization" db:"specialization"`
	YearsOfExperience int    `json:"years_of_experience" db:"years_of_experience"`
	LicenseNumber     string `json:"license_number" db:"license_number"`
	ClinicHours       string `json:"clinic_hours" db:"clinic_hours"`
	// AssignedLabID     int    `json:"assigned_lab_id" db:"assigned_lab_id"`
}

// rrepo implementation is remaining

type Nurse struct {
	NurseID          int    `json:"nurse_id" db:"nurse_id"`
	UserID           int    `json:"user_id" db:"user_id"`
	Shift            string `json:"shift" db:"shift"`
	AssignedDoctorID int    `json:"assigned_doctor_id" db:"assigned_doctor_id"`
	Qualification    string `json:"qualification" db:"qualification"`
}

type Menu struct {
	MenuID       int    `json:"menu_id"`
	Label        string `json:"label"`
	ToURL        string `json:"to_url"`
	Icon         string `json:"icon"`
	ParentMenuID *int   `json:"parent_menu_id"`
}

// Permission struct
type Permission struct {
	PermissionID int  `json:"permission_id"`
	RoleID       int  `json:"role_id"`
	MenuID       int  `json:"menu_id"`
	Allowed      bool `json:"allowed"`
}

// Hospital struct
type Hospital struct {
	HospitalID   int    `json:"hospital_id"`
	HospitalName string `json:"hospital_name"`
	CityID       int    `json:"city_id"`
	StateID      int    `json:"state_id"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	HospitalCode string `json:"hospital_code"`
}

// HospitalBranch struct defines the structure for the hospital_branch table
type HospitalBranch struct {
	BranchID    int    `json:"branch_id" db:"branch_id"`
	HospitalID  int    `json:"hospital_id" db:"hospital_id"`
	BranchName  string `json:"branch_name" db:"branch_name"`
	Address     string `json:"address" db:"address"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	BranchCode  string `json:"branch_code" db:"branchCode"`
}

// HospitalDoctor represents the relationship between Hospital and Doctor
type HospitalDoctor struct {
	HospitalDoctorID int    `json:"hd_id" db:"hd_id"`
	HospitalID       int    `json:"hospital_id" db:"hospital_id"`
	DoctorID         int    `json:"doctor_id" db:"doctor_id"`
	Role             string `json:"role" db:"role"`
	StartDate        string `json:"start_date" db:"start_date"` // Assuming date is in string format, you may convert it to time.Time
	EndDate          string `json:"end_date" db:"end_date"`     // Same as above
}
