-- {% func CapitalDrop() %}
drop table if exists "Capital";
-- {% endfunc %}

-- {% func CapitalCreate() %}
create table if not exists "Capital" (
  "ID" text not null,
  "Name" text not null,
  "Birthday" timestamp not null,
  "Deathday" timestamp,
  primary key ("ID")
);
-- {% endfunc %}
