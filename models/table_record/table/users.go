package table

import (
	"Go-Starter-Project/db"
	record "Go-Starter-Project/models/table_record"
)

// User - Struct che definisce la tabella "users"
// implementa TableRecordInterface
type User struct {
	tr       *record.TableRecord
	Name     string `json:"name" db:"name"`
	Lastname string `json:"lastname" db:"lastname"`
	Gender   string `json:"gender" db:"gender"`
}

// LoadAllUsers - Si occupa di restituire tutti gli utenti presenti nel database
func LoadAllUsers() ([]*User, error) {

	u := &User{}

	db := db.GetConnection()

	query := "SELECT " + record.AllField(u) + " FROM " + u.GetTableName()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*User

	for rows.Next() {

		u := NewUser()

		_, vField := record.GetFieldMapper(u)

		dest := append([]interface{}{&u.tr.RecordID}, vField...)

		err := rows.Scan(dest...)
		if err != nil {
			return nil, err
		}

		result = append(result, u)
	}

	return result, nil
}

// NewUser - Si occupa di istanziare un nuovo oggetto User istanziando il relativo TableRecord e impostandolo come "nuovo"
// È consigliato utilizzare sempre questo metodo per creare una nuova istanza di User
func NewUser() *User {

	u := new(User)
	u.tr = new(record.TableRecord)
	u.tr.SetIsNew(true)

	return u
}

// GetTableRecord - Restituisce l'istanza di TableRecord
func (u User) GetTableRecord() *record.TableRecord {
	return u.tr
}

// GetPrimaryKeyName - Restituisce il nome della chiave primaria
func (u User) GetPrimaryKeyName() string {
	return "record_id"
}

// GetTableName - Restituisce il nome della tabella
func (u User) GetTableName() string {
	return "users"
}

// New - Si occupa di istanziare una nuova struct andando ad istaziare table record e settanto il campo isNew a true
func (u *User) New() {

	u.tr = new(record.TableRecord)
	u.tr.SetIsNew(true)
}
