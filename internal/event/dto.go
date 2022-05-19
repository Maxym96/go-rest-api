package event

type CreateEventDTO struct {
	Name        string `db:"name"`
	Description string `db:"description"`
	DateAndTime string `db:"date_and_time"`
}
