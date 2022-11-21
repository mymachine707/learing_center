ALTER TABLE article DROP CONSTRAINT IF EXISTS fk_article_author;


ALTER TABLE learingCenter DROP CONSTRAINT IF EXISTS fk_direktor_person 
ALTER TABLE learingCenter DROP CONSTRAINT IF EXISTS fk_administrator_person 
ALTER TABLE teachers DROP CONSTRAINT IF EXISTS fk_teachers_person 
ALTER TABLE teachers DROP CONSTRAINT IF EXISTS fk_teachers_lessson 
ALTER TABLE teachers DROP CONSTRAINT IF EXISTS fk_teachers_learingCenter 
ALTER TABLE lessons DROP CONSTRAINT IF EXISTS fk_lessons_learningcenter 


DROP TABLE lessons;
DROP TABLE teachers;
DROP TABLE learingCenter;
DROP TABLE person;



