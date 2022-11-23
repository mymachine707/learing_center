CREATE ROLE Islombek WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD '946506184';

CREATE DATABASE learning_centers WITH OWNER Islombek;

migrate -path ./storage/migrations -database 'postgres://islombek:946506184@127.0.0.1:5432/learning_centers?sslmode=disable' up


SELECT p.id, p.firstname, p.lastname, p.middlename, p.birthday, p.job, p.phoneNumber, p.created_at, p.updated_at, p.deleted_at FROM person AS p;
                 id       
 | firstname |
 lastname

 middlename
 |birthday
|
job

 | phonenumber
|       created_at         | updated_at | deleted_at 

update person set middlename=Bahodir qizi where id=7c07bdc1-48b6-4140-b33c-45cede566703;

SELECT * FROM person WHERE ((firstname ILIKE '%' || 'a' || '%') OR (lastname ILIKE '%' || 'a' || '%') OR (job ILIKE '%' || 'a' || '%') OR (phonenumber ILIKE '%' || 'a' || '%')) AND deleted_at is null LIMIT 1 OFFSET 2;


UPDATE person SET firstname='Amirxon', lastname='Erkinov', middlename='Baxtiyor ugli', birthday='12.07.2019', job='chaqaloq', phonenumber='71-220-61-84', updated_at=now() WHERE id='b5d9eead-e14e-4f02-a4bc-313864e55406' and deleted_at is null;


SELECT l.id, l.coursename, l.learingcenterid, l.price, l.created_at, l.updated_at, l.deleted_at FROM lessons AS l;