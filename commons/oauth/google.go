package oauth

type Google struct {
	*BaseProvider
}

func googleProvider(base *BaseProvider) *Google {
	return &Google{base}
}
