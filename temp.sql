CREATE ROLE Islombek WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD '946506184';

CREATE DATABASE learning_centers WITH OWNER Islombek;

migrate -path ./storage/migrations -database 'postgres://islombek:946506184@127.0.0.1:5432/learning_centers?sslmode=disable' up