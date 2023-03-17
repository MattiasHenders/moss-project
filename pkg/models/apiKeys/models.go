package users

type ApiKey struct {
	ID          *string `json:"id" db:"id"`
	UserID      *string `json:"userID" db:"user_id"`
	Permissions *string `json:"permissions" db:"permissions"`
	Name        *string `json:"name" db:"name"`
	CreatedOn   *string `json:"createdOn" db:"created_on"`
}
