-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func GroupDrop() %}
drop table if exists "group";
-- {% endfunc %}

-- {% func GroupCreate() %}
create table if not exists "group" (
  "id" text not null,
  "child" text not null,
  "data" jsonb not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  "deleted" timestamp default now(),
  primary key ("id")
);
-- {% endfunc %}
