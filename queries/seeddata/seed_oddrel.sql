-- {% func OddrelSeedData() %}
insert into "oddrel" (
  "id", "project", "path"
) values (
  '90000000-0000-0000-0000-100000000000', '90000000-0000-0000-0000-000000000000', 'foo/bar'
) on conflict do nothing;
-- {% endfunc %}
