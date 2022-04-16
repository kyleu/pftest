-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func SizeInfo() %}
with recursive
  pg_inherit(inhrelid, inhparent) as (
    select inhrelid, inhparent
    from pg_inherits
    union
    select child.inhrelid, parent.inhparent
    from pg_inherit child, pg_inherits parent
    where child.inhparent = parent.inhrelid
  ),
  pg_inherit_short as (
    select *
    from pg_inherit
    where inhparent not in (select inhrelid from pg_inherit)
  )

select
  table_schema,
  table_name,
  row_estimate,
  total_bytes as "total",
  pg_size_pretty(total_bytes) as "total_pretty",
  case when index_bytes is null then 0 else index_bytes end as "index",
  pg_size_pretty(case when index_bytes is null then 0 else index_bytes end) as "index_pretty",
  case when toast_bytes is null then 0 else toast_bytes end as "toast",
  pg_size_pretty(case when toast_bytes is null then 0 else toast_bytes end) as "toast_pretty",
  case when table_bytes is null then 0 else table_bytes end as "table",
  pg_size_pretty(case when table_bytes is null then 0 else table_bytes end) as "table_pretty"
from (
  select *, total_bytes - index_bytes - coalesce(toast_bytes, 0) as table_bytes
  from (
    select c.oid, nspname as table_schema, relname as table_name,
      sum(c.reltuples) over (partition by parent) as row_estimate, sum(pg_total_relation_size(c.oid)) over (partition by parent) as total_bytes,
      sum(pg_indexes_size(c.oid)) over (partition by parent) as index_bytes,
      sum(pg_total_relation_size(reltoastrelid)) over (partition by parent) as toast_bytes, parent
    from (
      select pg_class.oid, reltuples, relname, relnamespace, pg_class.reltoastrelid, coalesce(inhparent, pg_class.oid) parent
      from pg_class left join pg_inherit_short on inhrelid = oid
      where relkind in ('r', 'p')
    ) c left join pg_namespace n on n.oid = c.relnamespace
    where nspname != 'pg_catalog' and nspname != 'information_schema'
  ) a
  where oid = parent
) a
order by total_bytes desc;
-- {% endfunc %}
