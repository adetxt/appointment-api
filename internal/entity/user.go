package entity

type User struct {
	ID       int64
	Name     string
	Role     Role
	Timezone string
}

type Role string

const (
	ROLE_USER  Role = "USER"
	ROLE_COACH Role = "COACH"
)

type GetUsersRequest struct {
	Page     int
	PageSize int
	Fields   []string
}
