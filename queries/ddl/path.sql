-- Content managed by Project Forge, see [projectforge.md] for details.
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
-- {% endfunc %}
