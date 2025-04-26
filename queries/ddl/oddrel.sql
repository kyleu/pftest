-- {% func OddrelDrop() %}
drop table if exists "oddrel";
-- {% endfunc %}

-- {% func OddrelCreate() %}
create table if not exists "oddrel" (
  "id" uuid not null,
  "project" uuid not null,
  "path" text not null,
  primary key ("id")
);
-- {% endfunc %}
