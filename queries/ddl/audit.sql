-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func AuditDrop() %}
drop table if exists "audit";
-- {% endfunc %}

-- {% func AuditCreate() %}
create table if not exists "audit" (
  "id" uuid not null,
  "app" text not null,
  "act" text not null,
  "client" text not null,
  "server" text not null,
  "user" text not null,
  "metadata" jsonb not null,
  "message" text not null,
  "started" timestamp not null default now(),
  "completed" timestamp not null default now(),
  primary key ("id")
);
-- {% endfunc %}
