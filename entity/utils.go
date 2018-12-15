package entity

// Some utils
// hasID for check id in entity
func hasID(idList []int, id int) bool {
	for _, sid := range idList {
		if sid == id {
			return true
		}
	}
	return false
}
