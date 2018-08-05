package model

type UserRole struct {
	BaseFields
	RoleName     string `json:"role_name"`
	DepartmentId string `json:"department_id"`
	LeadId       string `json:"lead_id"`
}
