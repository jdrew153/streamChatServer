package types

type User struct {
	Id       string `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Role     Role   `bson:"role" json:"role"`
	Password string `bson:"password" json:"password"`
}

type Role string

const (
	admin Role = "ADMIN"
	user  Role = "USER"
)

type Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}
