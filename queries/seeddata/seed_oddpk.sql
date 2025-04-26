-- {% func OddPKSeedData() %}
insert into "oddpk" (
  "project", "path", "name"
) values (
  '90000000-0000-0000-0000-000000000000', 'foo/bar', 'Project 1'
) on conflict do nothing;
-- {% endfunc %}
