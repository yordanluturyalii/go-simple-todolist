package helper

func PanicIfNil(err error) {
	if err != nil {
		panic(err)
	}
}
