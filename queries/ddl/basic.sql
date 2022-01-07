-- {% func BasicDrop() %}
drop table if exists "basic";
-- {% endfunc %}

-- {% func BasicCreate() %}
create table if not exists "basic" (
  "id" uuid not null,
  "name" text not null,
  "created" timestamp not null default now(),
  primary key ("id", "name")
);

create index if not exists "basic__id_idx" on "basic" ("id");

create index if not exists "basic__name_idx" on "basic" ("name");
-- {% endfunc %}
