<!--- Content managed by Project Forge, see [projectforge.md] for details. -->
# Environment Variables

All variable names can be set in uppercase or lowercase

| Name                          | Type   | Description                                                                      |
|-------------------------------|--------|----------------------------------------------------------------------------------|
| `app_display_name`            | string | replaces the UI title                                                            |
| `app_display_name_append`     | string | added to the end of the UI title                                                 |
| `app_nav_color_dark`          | string | sets the navigation color for users with dark mode, defaults to theme color      |
| `app_nav_color_light`         | string | sets the navigation color for users with light mode, defaults to theme color     |
| `compression_enabled`         | bool   | when set, compresses HTTP responses based on the request headers                 |
| `controller_metrics_disabled` | bool   | when set, skips metrics for controller methods                                   |
| `db_database`                 | string | database for the default MySQL connection                                        |
| `db_debug`                    | bool   | when set, logs all database interactions                                         |
| `db_file`                     | string | filename and path for the default SQLite connection                              |
| `db_host`                     | string | hostname for the default MySQL connection, defaults to `localhost`               |
| `db_metrics_disabled`         | bool   | when set, skips metrics for database methods                                     |
| `db_password`                 | string | password for the default MySQL connection                                        |
| `db_port`                     | int    | port for the default MySQL connection, defaults to `3306`                        |
| `db_schema`                   | string | schema for the default MySQL connection                                          |
| `db_ssl`                      | bool   | when set, uses SSL sessions for connections                                      |
| `db_user`                     | string | user for the default MySQL connection                                            |
| `logging_format`              | string | When set to `json`, forces the logging format                                    |
| `logging_level`               | string | minimum logging level to display, one of [`debug`, `info`, `warn`, `error`]      |
| `max_connections`             | int    | max connection count for the default MySQL connection                            |
| `notebook_base_url`           | string | base URL of the notebook server, defaults to [app port + 10]                     |
| `oauth_protocol`              | string | protocol to use for OAuth callbacks, defaults to the request's hostname          |
| `oauth_redirect`              | string | final URL to use for OAuth callbacks, overrides other options                    |
| `openid_connect_name`         | string | when OpenID is enabled, this controls the display name                           |
| `openid_connect_url`          | string | when OpenID is enabled, this determines the URL                                  |
| `read_db_database`            | string | database for the read-only PostgreSQL connection                                 |
| `read_db_debug`               | bool   | when set, logs all database interactions for the read-only database              |
| `read_db_host`                | string | hostname for the read-only PostgreSQL connection, defaults to `localhost`        |
| `read_db_password`            | string | password for the read-only PostgreSQL connection                                 |
| `read_db_port`                | int    | port for the read-only PostgreSQL connection, defaults to `3306`                 |
| `read_db_schema`              | string | schema for the read-only PostgreSQL connection                                   |
| `read_db_user`                | string | user for the read-only PostgreSQL connection                                     |
| `read_max_connections`        | int    | max connection count for the read-only PostgreSQL connection                     |
| `telemetry_disabled`          | bool   | when set, disables all telemetry                                                 |
| `telemetry_endpoint`          | string | address of OpenTelemetry collector (when enabled), defaults to `localhost:55681` |
| `pftest_encryption_key`       | string | encryption key for web sessions, defaults to `pftest_secret`, warns if missing   |
