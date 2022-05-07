-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func AuditedDrop() %}
drop table if exists "audited";
-- {% endfunc %}

-- {% func AuditedCreate() %}
create table if not exists "audited" (
  "id" uuid not null,
  "name" text not null,
  primary key ("id")
);
-- {% endfunc %}
