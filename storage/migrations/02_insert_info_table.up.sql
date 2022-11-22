BEGIN;

-- person
INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('dfe22f93-9a83-4a59-9afa-80c29c896c7f', 'Islombek', 'Oripov', 'Erkinjon ugli', '09-05-1998', 'Direktor','94-650-61-84') ON CONFLICT DO NOTHING;
INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('85140919-c27d-48e3-bb67-5bf1363d9637', 'Nilufar', 'Yunusjonova', 'Farxod qizi', '02-05-2004', 'Menedjer','99-142-70-33') ON CONFLICT DO NOTHING;
INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('7c07bdc1-48b6-4140-b33c-45cede566703', 'Aziza', 'Tuychiyeva', 'Nematulla qizi', '27-06-1975', 'Massaj','94-642-27-43') ON CONFLICT DO NOTHING;
INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('e259ff05-ad8f-4d09-ae1c-308a5db29a79', 'Aziza', 'Ubaydullayeva',  'Dilshod qizi', '02-06-1992', 'Pishiriq','99-333-22-87') ON CONFLICT DO NOTHING;
INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('dded95ac-b2f0-4d58-b047-472857fecde7', 'Ziyoda', 'Xusniddinova',  'Nuriddin qizi', '08-08-1992', 'Pishiriq','94-121-35-84') ON CONFLICT DO NOTHING;
INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('c6214094-973a-446f-afcf-e3630a25a1e2', 'Saidamir', 'Botirov', 'Bahodir ugli', '20-04-1998', 'Golang developer','94-154-34-84') ON CONFLICT DO NOTHING;
INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('b153f810-edb5-4a23-abcb-2e8352d5a924', 'Zafar', 'Saidov', 'Rustam ugli', '14-05-1988', 'Devops Programmist','94-655-54-82') ON CONFLICT DO NOTHING;

-- learning center
INSERT into LearingCenter (id, names, direktorID, administratorID, address_s) VALUES ('5a97e683-8747-4054-a344-13c9e51891a1', 'Uktamxon_nur_ziyo','dfe22f93-9a83-4a59-9afa-80c29c896c7f', '85140919-c27d-48e3-bb67-5bf1363d9637', 'Yunusobod 3-1-29') ON CONFLICT DO NOTHING;
INSERT into LearingCenter (id, names, direktorID, administratorID, address_s) VALUES ('1628bd13-2bc9-4b08-9341-dec16ade38cd', 'Uacademy','85140919-c27d-48e3-bb67-5bf1363d9637', 'dded95ac-b2f0-4d58-b047-472857fecde7', 'Mirzo Ulugbek tumani') ON CONFLICT DO NOTHING;


-- lesson
INSERT into lessons (id, courseName, learingCenterID, price) VALUES ('03db9d40-7bd9-4382-b6cd-34d597bc20b5', 'Massaj', '5a97e683-8747-4054-a344-13c9e51891a1', '450000') ON CONFLICT DO NOTHING;
INSERT into lessons (id, courseName, learingCenterID, price) VALUES ('adf4fa60-e10d-4d9e-bec4-3304347931fc', 'Golang', '1628bd13-2bc9-4b08-9341-dec16ade38cd', '1500000') ON CONFLICT DO NOTHING;
INSERT into lessons (id, courseName, learingCenterID, price) VALUES ('b8023df8-60b9-481a-94f6-010a231faeb6', 'Uygur taomlari', '5a97e683-8747-4054-a344-13c9e51891a1', '350000') ON CONFLICT DO NOTHING;
INSERT into lessons (id, courseName, learingCenterID, price) VALUES ('033ec38a-f895-477c-8313-71d9c16058c3', 'Oliy tort Mastika', '5a97e683-8747-4054-a344-13c9e51891a1', '450000') ON CONFLICT DO NOTHING;
INSERT into lessons (id, courseName, learingCenterID, price) VALUES ('f6fd1b78-88d4-402d-9d9a-c613378b521f', 'Devops', '1628bd13-2bc9-4b08-9341-dec16ade38cd', '1500000') ON CONFLICT DO NOTHING;
INSERT into lessons (id, courseName, learingCenterID, price) VALUES ('8f327d8e-6582-42e7-9504-0233e5c3f7db', 'Kompyuter saboqlari', '1628bd13-2bc9-4b08-9341-dec16ade38cd', '350000') ON CONFLICT DO NOTHING;

-- teacher
INSERT into teachers (id, personID, fulname, lessonsID, learingCenterID) VALUES ('d94458ba-a81e-46b3-9ac4-cb8f1a2a463b', '7c07bdc1-48b6-4140-b33c-45cede566703','Aziza Tuychiyeva Dilshod qizi', '03db9d40-7bd9-4382-b6cd-34d597bc20b5', '5a97e683-8747-4054-a344-13c9e51891a1') ON CONFLICT DO NOTHING;
INSERT into teachers (id, personID, fulname, lessonsID, learingCenterID) VALUES ('b5bf5601-ca27-43b4-bc38-a6a6784a9e97', 'c6214094-973a-446f-afcf-e3630a25a1e2','Botirov Saidamir Bahodir ugli', 'adf4fa60-e10d-4d9e-bec4-3304347931fc', '1628bd13-2bc9-4b08-9341-dec16ade38cd') ON CONFLICT DO NOTHING;
INSERT into teachers (id, personID, fulname, lessonsID, learingCenterID) VALUES ('2c7d07ba-d006-4761-b730-e15b0f63267f', 'c6214094-973a-446f-afcf-e3630a25a1e2','Botirov Saidamir Bahodir ugli', '8f327d8e-6582-42e7-9504-0233e5c3f7db', '5a97e683-8747-4054-a344-13c9e51891a1') ON CONFLICT DO NOTHING;
INSERT into teachers (id, personID, fulname, lessonsID, learingCenterID) VALUES ('933c9ba9-8913-4bd1-9043-e85054d2c551', 'b153f810-edb5-4a23-abcb-2e8352d5a924','Zafar Saidov Rustam ugli', 'f6fd1b78-88d4-402d-9d9a-c613378b521f', '1628bd13-2bc9-4b08-9341-dec16ade38cd') ON CONFLICT DO NOTHING;

COMMIT;
