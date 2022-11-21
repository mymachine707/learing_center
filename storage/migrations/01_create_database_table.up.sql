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
   direktorID CHAR(255),
   administratorID CHAR(255),
   address_s VARCHAR(255) NOT NULL,
   created_at TIMESTAMP DEFAULT now(),
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS teachers (
	personID CHAR(36) PRIMARY KEY,
	fulname VARCHAR(255) NOT NULL,
	lessonsID CHAR(255),
    learingCenterID CHAR(255),
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);


CREATE TABLE lessons (
   id CHAR(36) PRIMARY KEY,
   courseName VARCHAR(255) UNIQUE NOT NULL,
   learingCenterID CHAR(255),
   Price VARCHAR(255) UNIQUE NOT NULL,
   created_at TIMESTAMP DEFAULT now(),
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);


ALTER TABLE learingCenter ADD CONSTRAINT fk_direktor_person FOREIGN KEY (direktorID) REFERENCES person (id);
ALTER TABLE learingCenter ADD CONSTRAINT fk_administrator_person FOREIGN KEY (administratorID) REFERENCES person (id);
ALTER TABLE teachers ADD CONSTRAINT fk_teachers_person FOREIGN KEY (personID) REFERENCES person (id);
ALTER TABLE teachers ADD CONSTRAINT fk_teachers_lessson FOREIGN KEY (lessonsID) REFERENCES lessons (id);
ALTER TABLE teachers ADD CONSTRAINT fk_teachers_learingCenter FOREIGN KEY (learingCenterID) REFERENCES learingCenter (id);
ALTER TABLE lessons ADD CONSTRAINT fk_lessons_learningcenter FOREIGN KEY (learingCenterID) REFERENCES learingCenter (id);


