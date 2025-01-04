package utilify

func StrInSlice(value string, slice []string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func AddStrToSlice(slice []string, value string) []string {
	return append(slice, value)
}

func RmvStrInSlice(slice []string, value string) []string {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func I64InSlice(value int64, slice []int64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func AddI64ToSlice(slice []int64, value int64) []int64 {
	return append(slice, value)
}

func RmvI64InSlice(slice []int64, value int64) []int64 {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
