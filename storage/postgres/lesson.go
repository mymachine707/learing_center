package postgres

import (
	"mymachine707/models"

	"golang.org/x/text/search"
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

func (stg Postgres) GetListLesson(offset, limit int, search string)(resp []models.Lesson, error){
		
	rows, err := stg.db.Queryx(`SELECT * FROM lessons where 
	((coursname ILIKE '%' || $1 || '%')) AND deleted_at is null LIMIT $2 OFFSET $3`, search, limit, offset)

	


}
//func (stg Postgres) UpdateLessonByID() {}
//func (stg Postgres) DeleteLessonByID() {}
