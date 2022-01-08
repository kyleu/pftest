-- {% func VersionDrop() %}
drop table if exists "version_revision";
drop table if exists "version";
-- {% endfunc %}

-- {% func VersionCreate() %}
create table if not exists "version" (
  "id" text not null,
  "current_revision" int not null default 1,
  "const" text not null,
  "updated" timestamp default now(),
  "deleted" timestamp default now(),
  primary key ("id")
);

create table if not exists "version_revision" (
  "version_id" text not null,
  "revision" int not null default 1,
  "var" jsonb not null,
  "created" timestamp not null default now(),
  foreign key ("version_id") references version ("id"),
  primary key ("version_id", "revision")
);
create index if not exists "version_revision__version_id_idx" on "version_revision" ("version_id");

create index if not exists "version_revision__version_id_idx" on "version_revision" ("version_id");
-- {% endfunc %}
