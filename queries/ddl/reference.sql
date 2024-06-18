-- {% func ReferenceDrop() %}
drop table if exists "reference";
-- {% endfunc %}

-- {% func ReferenceCreate() %}
create table if not exists "reference" (
  "id" uuid not null,
  "custom" jsonb not null,
  "self" jsonb not null,
  "created" timestamp not null default now(),
  primary key ("id")
);
-- {% endfunc %}
