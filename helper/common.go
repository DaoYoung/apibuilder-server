package helper
func Contains(intSlice []string, searchInt string) bool {
	if len(intSlice) == 0 {
		return false
	}
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}

func SetForbidUpdateFields(fs ...string) []string {
	res := []string{"id", "created_at", "updated_at", "deleted_at"}
	for _, value := range fs {
		res = append(res, value)
	}
	return res
}
