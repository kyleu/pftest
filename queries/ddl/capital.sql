<!-- Content managed by Project Forge, see [projectforge.md] for details. -->
-- {% func CapitalDrop() %}
drop table if exists "Capital_Version";
drop table if exists "Capital";
-- {% endfunc %}

-- {% func CapitalCreate() %}
create table if not exists "Capital" (
  "ID" text not null,
  "current_Version" int not null default 1,
  primary key ("ID")
);

create table if not exists "Capital_Version" (
  "Capital_ID" text not null,
  "Version" int not null default 1,
  "Name" text not null,
  "Birthday" timestamp not null,
  "Deathday" timestamp,
  foreign key ("Capital_ID") references "Capital" ("ID"),
  primary key ("Capital_ID", "Version")
);
create index if not exists "Capital_Version__Capital_ID_idx" on "Capital_Version" ("Capital_ID");

create index if not exists "Capital_Version__Capital_ID_idx" on "Capital_Version" ("Capital_ID");
-- {% endfunc %}
