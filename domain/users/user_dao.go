package users

import (
	"fmt"
	"github.com/ArturoAguilar1/bookstore_users-api/utils/date_utils"
	"github.com/ArturoAguilar1/bookstore_users-api/utils/errors"
)

//Work with methods instead of functions
//Later we profundizar en porqué es mejor los métodos que las funciones


var (
	usersDB = make(map[int64] *User)
)

func (user User)Save() *errors.RestErr{
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered",user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exits",user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = &user
	return nil
}


//Get by id, or get by primary key
//Puedo devolver el usuario pedido, o algun error(ejemplo no existe el usuario)
func (user User)Get() (*errors.RestErr) {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	//Temporal hasta que este la base de datos:
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}