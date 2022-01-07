-- {% func CreateDatabase() %}
create role "pftest" with login password 'pftest';

create database "pftest";
alter database "pftest" set timezone to 'utc';
grant all privileges on database "pftest" to "pftest";
-- {% endfunc %}
