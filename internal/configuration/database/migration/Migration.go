package migration

import (
	"fmt"
	"os"

	"github.com/bisudoh/gokapi/internal/configuration/database"
	"github.com/bisudoh/gokapi/internal/configuration/database/dbabstraction"
	"github.com/bisudoh/gokapi/internal/environment/flagparser"
)

// Do checks the passed flags for a migration and then executes it
func Do(flags flagparser.MigrateFlags) {
	oldDb, err := database.ParseUrl(flags.Source, true)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		osExit(1)
		return
	}
	newDb, err := database.ParseUrl(flags.Destination, false)
	if err != nil {
		fmt.Println(err.Error())
		osExit(2)
		return
	}
	fmt.Printf("Migrating %s database %s to %s database %s\n", getType(oldDb.Type), oldDb.HostUrl, getType(newDb.Type), newDb.HostUrl)
	database.Migrate(oldDb, newDb)
}

func getType(input int) string {
	switch input {
	case dbabstraction.TypeSqlite:
		return "SQLite"
	case dbabstraction.TypeRedis:
		return "Redis"
	}
	return "Invalid"
}

// Declared for testing
var osExit = os.Exit
