package auth

type Payload struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
	Role  string `json:"role"`
	Exp   int64  `json:"exp"`
}
