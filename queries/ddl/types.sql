-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func TypesDrop() %}
drop type if exists "foo";
-- {% endfunc %}

-- {% func TypesCreate() %}
do $$ begin
  create type "foo" as enum ('a', 'b', 'c', 'd');
exception
  when duplicate_object then null;
end $$;
-- {% endfunc %}
