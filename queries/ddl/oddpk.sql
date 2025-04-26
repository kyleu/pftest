-- {% func OddPKDrop() %}
drop table if exists "oddpk";
-- {% endfunc %}

-- {% func OddPKCreate() %}
create table if not exists "oddpk" (
  "project" uuid not null,
  "path" text not null,
  "name" text not null,
  primary key ("project", "path")
);

create index if not exists "oddpk__project_idx" on "oddpk" ("project");

create index if not exists "oddpk__path_idx" on "oddpk" ("path");
-- {% endfunc %}
