-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func TypesDrop() %}
drop type if exists "foo";
-- {% endfunc %}

-- {% func TypesCreate() %}
create type "foo" as enum ('a', 'b', 'c', 'd');
-- {% endfunc %}
