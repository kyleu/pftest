-- {% func TypesDrop() %}
drop type if exists "foo";
drop type if exists "bar";
-- {% endfunc %}

-- {% func TypesCreate() %}
do $$ begin
  create type "bar" as enum ('first_value', 'second_value', 'unknown');
exception
  when duplicate_object then null;
end $$;
do $$ begin
  create type "foo" as enum ('a', 'b', 'c', 'd');
exception
  when duplicate_object then null;
end $$;
-- {% endfunc %}
