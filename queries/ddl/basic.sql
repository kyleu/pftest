<!-- Content managed by Project Forge, see [projectforge.md] for details. -->
-- {% func BasicDrop() %}
drop table if exists "basic";
-- {% endfunc %}

-- {% func BasicCreate() %}
create table if not exists "basic" (
  "id" uuid not null,
  "name" text not null,
  "created" timestamp not null default now(),
  primary key ("id")
);
-- {% endfunc %}
