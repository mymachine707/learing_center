Storage 
Handlar
Crud

// Kodlani githubga push qimaslik kere



// This link for install migration 
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

// This code for create migration 
migrate create -ext sql -dir ./storage/migrations -seq -digits 2 insert_info_table

//!! groups database ochishim kere


// migration UP 

migrate -path ./storage/migrations -database 'postgres://islombek:946506184@127.0.0.1:5432/learning_centers?sslmode=disable' up