package service

func GetDefaultWhere(where map[string]interface{}) map[string]interface{} {
	where["deleted_at"] = nil
	return where
}
