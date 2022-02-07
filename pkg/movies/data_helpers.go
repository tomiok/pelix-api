package movies

func SafeString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func SafeInt(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}
