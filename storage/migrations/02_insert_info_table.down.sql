BEGIN;
-- teacher
DELETE FROM teachers where id='d94458ba-a81e-46b3-9ac4-cb8f1a2a463b';
DELETE FROM teachers where id='b5bf5601-ca27-43b4-bc38-a6a6784a9e97';
DELETE FROM teachers where id='2c7d07ba-d006-4761-b730-e15b0f63267f';
DELETE FROM teachers where id='933c9ba9-8913-4bd1-9043-e85054d2c551';

-- lesson
DELETE FROM lessons where id='03db9d40-7bd9-4382-b6cd-34d597bc20b5';
DELETE FROM lessons where id='adf4fa60-e10d-4d9e-bec4-3304347931fc';
DELETE FROM lessons where id='b8023df8-60b9-481a-94f6-010a231faeb6';
DELETE FROM lessons where id='033ec38a-f895-477c-8313-71d9c16058c3';
DELETE FROM lessons where id='f6fd1b78-88d4-402d-9d9a-c613378b521f';
DELETE FROM lessons where id='8f327d8e-6582-42e7-9504-0233e5c3f7db';

-- learning center
DELETE FROM LearingCenter where id='5a97e683-8747-4054-a344-13c9e51891a1';
DELETE FROM LearingCenter where id='1628bd13-2bc9-4b08-9341-dec16ade38cd';

-- person
DELETE FROM person where id='dfe22f93-9a83-4a59-9afa-80c29c896c7f';
DELETE FROM person where id='85140919-c27d-48e3-bb67-5bf1363d9637';
DELETE FROM person where id='7c07bdc1-48b6-4140-b33c-45cede566703';
DELETE FROM person where id='e259ff05-ad8f-4d09-ae1c-308a5db29a79';
DELETE FROM person where id='dded95ac-b2f0-4d58-b047-472857fecde7';
DELETE FROM person where id='c6214094-973a-446f-afcf-e3630a25a1e2';
DELETE FROM person where id='b153f810-edb5-4a23-abcb-2e8352d5a924';
COMMIT;
