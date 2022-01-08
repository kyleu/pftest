-- {% func TimestampDrop() %}
drop table if exists "timestamp";
-- {% endfunc %}

-- {% func TimestampCreate() %}
create table if not exists "timestamp" (
  "id" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  "deleted" timestamp default now(),
  primary key ("id")
);
-- {% endfunc %}
