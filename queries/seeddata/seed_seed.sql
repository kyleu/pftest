-- {% func SeedSeedData() %}
insert into "seed" (
  "id", "name", "size", "obj"
) values (
  '00000000-0000-0000-0000-000000000001', 'A', 1, '{"foo":"a"}'
), (
  '00000000-0000-0000-0000-000000000002', 'B', 2, '{"foo":"b"}'
), (
  '00000000-0000-0000-0000-000000000003', 'C', 3, '{"foo":"c"}'
), (
  '00000000-0000-0000-0000-000000000004', 'D', 4, '{"foo":"d"}'
) on conflict do nothing;
-- {% endfunc %}
