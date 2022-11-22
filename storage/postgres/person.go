package postgres

import (
	"errors"
	"mymachine707/models"
)

// AddPerson ...
func (stg Postgres) AddPerson(id string, entity models.CreatePersonModul) error {
	var err error
	_, err = stg.db.Exec(`
	INSERT into person (
		id, 
		firstname, 
		lastname, 
		middlename, 
		birthday, 
		job, 
		phoneNumber
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
			) ON CONFLICT DO NOTHING`,
		id,
		entity.Firstname,
		entity.Lastname,
		entity.Middlename,
		entity.Birthday,
		entity.Job,
		entity.PhoneNumber,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetPersonByID ...
func (stg Postgres) GetPersonByID(id string) (models.Person, error) {

	var a models.Person

	err := stg.db.QueryRow(`
	SELECT 
	p.id,
	p.firstname,
	p.lastname,
	p.middlename,
	p.birthday,
	p.job,
	p.phoneNumber,
	p.created_at,
	p.updated_at,
	p.deleted_at FROM person AS p WHERE id=$1 AND deleted_at is null`, id).Scan(
		&a.ID,
		&a.Firstname,
		&a.Lastname,
		&a.Middlename,
		&a.Birthday,
		&a.Job,
		&a.PhoneNumber,
		&a.CreatedAt,
		&a.UpdatedAt,
		&a.DeletedAt,
	)

	if err != nil {
		return a, err
	}

	return a, nil
}

// GetPersonList ...
func (stg Postgres) GetPersonList(offset, limit int, search string) (resp []models.Person, err error) {

	rows, err := stg.db.Queryx(`
	SELECT * FROM person WHERE 
	((firstname ILIKE '%' || %1 || '%') OR 
	(lastname ILIKE '%' || $1 || '%') OR 
	(job ILIKE '%' || $1 || '%')) OR 
	(phonenumber ILIKE '%' || $1 || '%'))
	
	AND deleted_at is null LIMIT $2 OFFSET $3`, search, limit, offset)

	if err != nil {
		return resp, err
	}

	for rows.Next() {

		var a models.Person

		err = rows.Scan(
			&a.ID,
			&a.Firstname,
			&a.Lastname,
			&a.Middlename,
			&a.Birthday,
			&a.Job,
			&a.PhoneNumber,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.DeletedAt,
		)
		if err != nil {
			return resp, err
		}

		resp = append(resp, a)
	}
	return resp, nil
}

// UpdatePerson ...
func (stg Postgres) UpdatePerson(person models.UpdatePersonModel) error {

	rows, err := stg.db.NamedExec(`
	UPDATE person SET firstname:=f, lastname:=l, middlename:=m, birthday:=b, job:=j, phonenumber:=p, updated_at=now() WHERE id:=id and where deleted_at is null`, map[string]interface{}{
		"id": person.ID,
		"f":  person.Firstname,
		"l":  person.Lastname,
		"m":  person.Middlename,
		"b":  person.Birthday,
		"j":  person.Job,
		"p":  person.PhoneNumber,
	})

	if err != nil {
		return err
	}

	n, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}
	return errors.New("person not found")
}

// DeletePerson ...
func (stg Postgres) DeletePerson(id string) error {
	rows, err := stg.db.Exec(`UPDATE person SET deleted_at=now() WHERE id=$1 and deleted_at is null`, id)

	if err != nil {
		return err
	}

	n, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("Person cannot deleted becouse person not found")
}
