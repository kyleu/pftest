-- {% func PathDrop() %}
drop table if exists "path";
-- {% endfunc %}

-- {% func PathCreate() %}
create table if not exists "path" (
  "id" uuid not null,
  "name" text not null,
  "status" text not null,
  "created" timestamp not null default now(),
  primary key ("id")
);
create index if not exists "basic_created_idx" on "basic" ("created");
-- {% endfunc %}
