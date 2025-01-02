package utilify

func StrInSlice(value string, slice []string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func I64InSlice(value int64, slice []int64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
