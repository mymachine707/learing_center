BEGIN;

-- person

INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('dfe22f93-9a83-4a59-9afa-80c29c896c7f', 'Islombek', 'Oripov', 'Erkinjon ugli', '09-05-1998', 'Direktor','94-650-61-84') ON CONFLICT DO NOTHING;
INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('85140919-c27d-48e3-bb67-5bf1363d9637', 'Nilufar', 'Yunusjonova', 'Farxod qizi', '02-05-2004', 'Menedjer','99-142-70-33') ON CONFLICT DO NOTHING;
INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('7c07bdc1-48b6-4140-b33c-45cede566703', 'Tuychiyeva', 'Aziza', 'Nematulla qizi', '27-06-1975', 'Massaj','94-642-27-43') ON CONFLICT DO NOTHING;
INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('e259ff05-ad8f-4d09-ae1c-308a5db29a79', 'Ubaydullayeva', 'Aziza', 'Dilshod qizi', '14-07-1989', 'Pishiriq','99-333-22-87') ON CONFLICT DO NOTHING;
INSERT into person (id, firstname, lastname, middlename, birthday, job, phoneNumber) VALUES ('dded95ac-b2f0-4d58-b047-472857fecde7', 'Xusniddinova', 'Ziyoda', 'Nuriddin qizi', '14-05-1992', 'Pishiriq','94-111-85-84') ON CONFLICT DO NOTHING;


-- learning center

INSERT into LearingCenter (id, names, direktorID, AdministratorID, address_s) VALUES ('5a97e683-8747-4054-a344-13c9e51891a1', 'Uktamxon_nur_ziyo','dfe22f93-9a83-4a59-9afa-80c29c896c7f', '85140919-c27d-48e3-bb67-5bf1363d9637', 'Yunusobod 3-1-29') ON CONFLICT DO NOTHING;
INSERT into LearingCenter (id, names, direktorID, AdministratorID, address_s) VALUES ('1628bd13-2bc9-4b08-9341-dec16ade38cd', 'Milliy_kasblar','85140919-c27d-48e3-bb67-5bf1363d9637', 'dded95ac-b2f0-4d58-b047-472857fecde7', 'Yunusobod Minor metro') ON CONFLICT DO NOTHING;



COMMIT;
