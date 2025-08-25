package database

import "database/sql"

type Models struct {
	User UserModel
	Event EventModel
	Attendee AttendeeModel

}			

func NewModels(db *sql.DB) Models {
	return Models{
		User: UserModel{DB: db},
		Event: EventModel{DB: db},
		Attendee: AttendeeModel{DB: db},
	}
}