-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func HistDrop() %}
drop table if exists "hist_history";
drop table if exists "hist";
-- {% endfunc %}

-- {% func HistCreate() %}
create table if not exists "hist" (
  "id" text not null,
  "data" jsonb not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  primary key ("id")
);

create table if not exists "hist_history" (
  "id" uuid,
  "hist_id" text not null,
  "o" jsonb not null,
  "n" jsonb not null,
  "c" jsonb not null,
  "created" timestamp not null default now(),
  foreign key ("hist_id") references "hist" ("id"),
  primary key ("id")
);
-- {% endfunc %}
