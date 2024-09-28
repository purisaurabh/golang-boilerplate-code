package repository

type PostProfileRepo struct {
	Name       string `db:"name"`
	Email      string `db:"email"`
	Mobile     string `db:"mobile"`
	Created_At int64  `db:"created_at"`
	Updated_At int64  `db:"updated_at"`
}

type ListProfileData struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Email      string `db:"email"`
	Mobile     string `db:"mobile"`
	Created_At int64  `db:"created_at"`
	Updated_At int64  `db:"updated_at"`
}

type UpdateProfileData struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Email      string `db:"email"`
	Mobile     string `db:"mobile"`
	Updated_At int64  `db:"updated_at"`
}
