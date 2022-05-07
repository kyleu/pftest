-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func HistoryDrop() %}
drop table if exists "history_history";
drop table if exists "history";
-- {% endfunc %}

-- {% func HistoryCreate() %}
create table if not exists "history" (
  "id" text not null,
  "data" jsonb not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  primary key ("id")
);

create table if not exists "history_history" (
  "id" uuid,
  "history_id" text not null,
  "o" jsonb not null,
  "n" jsonb not null,
  "c" jsonb not null,
  "created" timestamp not null default now(),
  foreign key ("history_id") references "history" ("id"),
  primary key ("id")
);
-- {% endfunc %}
