package model

type Department struct {
	BaseFields
	DepartmentName string `json:"department_name"`
	Pid            string `json:"pid"`
	LeadId         string `json:"lead_id"`
}
