package entites

type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	AvatarID   string `json:"avatar_id"`
	Permission int32  `json:"permission"`
	Status     string `json:"status"`
}
