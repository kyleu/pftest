-- {% func MixedCaseDrop() %}
drop table if exists "mixed_case";
-- {% endfunc %}

-- {% func MixedCaseCreate() %}
create table if not exists "mixed_case" (
  "id" text not null,
  "test_field" text not null,
  "another_field" text not null,
  primary key ("id")
);
-- {% endfunc %}
