package users

type (
	Users struct {
		Username string `json:"username" db:"username"`
	}
)

func (usr *Users) CreateOneUser() error {

}

func (usr *Users) UpdateOneUser() error {

}

func (usr *Users) DeleteOneUser() error {

}

func (usr *Users) QueryAllUsers() error {

}
