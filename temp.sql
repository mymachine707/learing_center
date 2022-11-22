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