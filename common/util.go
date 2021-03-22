package common

// CheckError - check error occurred
func CheckError(err error) {

	if err != nil {
		panic(err)
	}

}
