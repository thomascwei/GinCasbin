package internal

var sampleAccountPasswords = map[string]string{
	"thomas": "123456",
	"jim":    "123456",
}

func auth(reqAccount, reqPassword string) bool {
	password, ok := sampleAccountPasswords[reqAccount]
	if !ok {
		return false
	}
	if password != reqPassword {
		return false
	}
	return true
}
