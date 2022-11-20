CREATE TABLE IF NOT EXISTS person (
	id CHAR(36) PRIMARY KEY,
	firstname VARCHAR(255) NOT NULL,
	lastname VARCHAR(255) NOT NULL,
    middlename VARCHAR(255) NOT NULL,
    birthday VARCHAR(255) NOT NULL,
    job VARCHAR(255) NOT NULL,
    phoneNumber VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE learingCenter (
   id CHAR(36) PRIMARY KEY,
   direktor VARCHAR(255) UNIQUE NOT NULL,
   administrator VARCHAR(255) NOT NULL,
   listcource VARCHAR(255) NOT NULL,
   listteachers VARCHAR(255) NOT NULL,
   address_s VARCHAR(255) NOT NULL,
   created_at TIMESTAMP DEFAULT now(),
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS teachers (
	id CHAR(36) PRIMARY KEY,
	fulname VARCHAR(255) NOT NULL,
	coursInfo VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);



//!!!
ALTER TABLE article ADD CONSTRAINT fk_article_author FOREIGN KEY (author_id) REFERENCES author (id);
