package modules

type person struct {
	ID          string
	Firstname   string
	Lastname    string
	Middlename  string
	Birthday    string // data
	Job         string // !!
	phoneNUmber string
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
