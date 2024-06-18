-- {% func TroubleDrop() %}
drop table if exists "trouble";
-- {% endfunc %}

-- {% func TroubleCreate() %}
create table if not exists "trouble" (
  "from" text not null,
  "where" jsonb not null,
  "selectcol" int not null default 1,
  "limit" text not null,
  "group" text not null,
  "delete" timestamp default now(),
  primary key ("from", "where")
);

create index if not exists "trouble__from_idx" on "trouble" ("from");

create index if not exists "trouble__where_idx" on "trouble" ("where");
-- {% endfunc %}
