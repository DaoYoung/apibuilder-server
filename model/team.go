package model

type Team struct {
	BaseFields
	LeadId   int    `json:"lead_id"`
	TeamName string `json:"team_name"`
}
func (mod *Team) Leader() *User {
	leader := &User{}
	ByID(leader, mod.LeadId)
	return leader
}