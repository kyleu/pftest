-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func SeedDrop() %}
drop table if exists "seed";
-- {% endfunc %}

-- {% func SeedCreate() %}
create table if not exists "seed" (
  "id" uuid not null,
  "name" text not null,
  "size" int not null,
  "obj" jsonb not null,
  primary key ("id")
);
-- {% endfunc %}
