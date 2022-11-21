
BEGIN;


-- learing center
DELETE FROM LearingCenter where id='5a97e683-8747-4054-a344-13c9e51891a1';
DELETE FROM LearingCenter where id='1628bd13-2bc9-4b08-9341-dec16ade38cd';

-- person
DELETE FROM person where id='dfe22f93-9a83-4a59-9afa-80c29c896c7f';
DELETE FROM person where id='85140919-c27d-48e3-bb67-5bf1363d9637';
DELETE FROM person where id='7c07bdc1-48b6-4140-b33c-45cede566703';
DELETE FROM person where id='e259ff05-ad8f-4d09-ae1c-308a5db29a79';
DELETE FROM person where id='dded95ac-b2f0-4d58-b047-472857fecde7';

COMMIT;