package masters

type SuperadminLab struct {
	LabID         int           `json:"lab_id" db:"lab_id"`
	LabName       string        `json:"lab_name" db:"lab_name"`
	LabCode       string        `json:"lab_code" db:"lab_code"`
	ValidityDate  string        `json:"validity_date" db:"validity_date"`
	CreatedOn     string        `json:"created_on" db:"created_on"`
	CreatedBy     string        `json:"created_by" db:"created_by"`
	BranchesDepts []BranchDepts `json:"branchDepts"`
}
