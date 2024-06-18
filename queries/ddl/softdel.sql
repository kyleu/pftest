-- {% func SoftdelDrop() %}
drop table if exists "softdel";
-- {% endfunc %}

-- {% func SoftdelCreate() %}
create table if not exists "softdel" (
  "id" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  "deleted" timestamp default now(),
  primary key ("id")
);
-- {% endfunc %}
