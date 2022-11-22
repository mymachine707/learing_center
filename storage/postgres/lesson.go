package postgres

import (
	"errors"
	"mymachine707/models"
)

// AddLesson ...
func (stg Postgres) AddLesson(id, entity models.CreateLessonModels) error {
	var err error
	_, err = stg.db.Exec(`INSERT into lessons (
		id, 
		courseName, 
		learingCenterID, 
		price
		) VALUES (
		$1,
		$2, 
		$3, 
		$4) ON CONFLICT DO NOTHING;`,
		id,
		entity.CourseName,
		entity.LearingCenterID,
		entity.Price,
	)

	if err != nil {
		return err
	}
	return nil
}

// GetLessonByID ...
func (stg Postgres) GetLessonByID(id string) (models.Lesson, error) {

	var a models.Lesson

	err := stg.db.QueryRow(`
		Select 
			l.id, 
			l.coursname,
			l.learingcenterid,
			l.price,
			l.created_at,
			l.updated_at,
			l.deleted_at
			from lessons as l where id=$1 AND deleted_at is not null`, id).Scan(
		a.ID,
		a.CourseName,
		a.LearingCenterID,
		a.Price,
		a.CreatedAt,
		a.UpdatedAt,
		a.DeletedAt,
	)

	if err != nil {
		return a, err
	}
	return a, nil
}

// GetListLesson ...
func (stg Postgres) GetListLesson(offset, limit int, search string) (resp []models.Lesson, err error) {

	rows, err := stg.db.Queryx(`SELECT * FROM lessons where 
	((coursname ILIKE '%' || $1 || '%')) AND deleted_at is null LIMIT $2 OFFSET $3`, search, limit, offset)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var a models.Lesson

		err = rows.Scan(
			&a.ID,
			&a.CourseName,
			&a.LearingCenterID,
			&a.Price,
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

// UpdateLessonByID ...
func (stg Postgres) UpdateLessonByID(lessonw models.UpdateLessonModels) error {

	rows, err := stg.db.NamedExec(`
	Update lessons Set coursname:=c, price:=p, updated_at=now() where id:=ids and deleted_at is null`, map[string]interface{}{
		"ids": lessonw.ID,
		"c":   lessonw.CourseName,
		"p":   lessonw.Price,
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
	return errors.New("lesson not found")
}

//func (stg Postgres) DeleteLessonByID() {}
