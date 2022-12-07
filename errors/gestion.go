package errors

func GestionError(err error) {
	if err != nil {
		panic(err)
	}
}

