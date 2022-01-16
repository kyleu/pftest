-- {% func TroubleDrop() %}
drop table if exists "trouble_selectcol";
drop table if exists "trouble";
-- {% endfunc %}

-- {% func TroubleCreate() %}
create table if not exists "trouble" (
  "from" text not null,
  "where" int not null,
  "current_selectcol" int not null default 1,
  "limit" text not null,
  "delete" timestamp default now(),
  primary key ("from", "where")
);

create index if not exists "trouble__from_idx" on "trouble"("from");

create index if not exists "trouble__where_idx" on "trouble"("where");

create table if not exists "trouble_selectcol" (
  "trouble_from" text not null,
  "trouble_where" int not null,
  "selectcol" int not null default 1,
  "group" text not null,
  foreign key ("trouble_from", "trouble_where") references trouble("from", "where"),
  primary key ("trouble_from", "trouble_where", "selectcol")
);
create index if not exists "trouble_selectcol__trouble_from_trouble_where_idx" on "trouble_selectcol"("trouble_from", "trouble_where");

create index if not exists "trouble_selectcol__trouble_from_idx" on "trouble_selectcol"("trouble_from");

create index if not exists "trouble_selectcol__trouble_where_idx" on "trouble_selectcol"("trouble_where");
-- {% endfunc %}
