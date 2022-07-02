// $PF_IGNORE$
package migrations

import (
	"github.com/kyleu/pftest/app/lib/database/migrate"
	"github.com/kyleu/pftest/queries/seeddata"
)

func LoadMigrations(debug bool) {
	migrate.RegisterMigration("create initial database", Migration1InitialDatabase(debug))
	migrate.RegisterMigration("seed data", seeddata.SeedDataAll())
}
