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
   direktor_ID CHAR(255),
   administrator_ID CHAR(255),
   listcource VARCHAR(255) NOT NULL,
   listteachers VARCHAR(255) NOT NULL,
   address_s VARCHAR(255) NOT NULL,
   created_at TIMESTAMP DEFAULT now(),
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS teachers (
	person_id CHAR(36) PRIMARY KEY,
	fulname VARCHAR(255) NOT NULL,
	lessons_ID CHAR(255),
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);


CREATE TABLE lessons (
   id CHAR(36) PRIMARY KEY,
   course VARCHAR(255) UNIQUE NOT NULL,
   learingCenter_ID CHAR(255),
   listteachers VARCHAR(255) NOT NULL,
   created_at TIMESTAMP DEFAULT now(),
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);


ALTER TABLE learingCenter ADD CONSTRAINT fk_direktor_person FOREIGN KEY (direktor_ID) REFERENCES person (id);
ALTER TABLE learingCenter ADD CONSTRAINT fk_administrator_person FOREIGN KEY (administrator_ID) REFERENCES person (id);
ALTER TABLE teachers ADD CONSTRAINT fk_teachers_lessson FOREIGN KEY (lessons_ID) REFERENCES lessons (id);
ALTER TABLE lessons ADD CONSTRAINT fk_lessons_learningcenter FOREIGN KEY (learingCenter_ID) REFERENCES learingCenter (id);


