-- {% func RelationDrop() %}
drop table if exists "relation";
-- {% endfunc %}

-- {% func RelationCreate() %}
create table if not exists "relation" (
  "id" uuid not null,
  "basic_id" uuid not null,
  "name" text not null,
  "created" timestamp not null default now(),
  foreign key ("basic_id") references "basic" ("id"),
  primary key ("id")
);

create index if not exists "relation__basic_id_idx" on "relation" ("basic_id");
-- {% endfunc %}
