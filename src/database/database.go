package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VOVOplay/creatorcoaster.com/src/config"
	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
	_ "github.com/go-sql-driver/mysql"
)

var all_config config.Config = config.GetConfig()
var db_config config.DatabaseConfig = all_config.DatabaseConfig

var dsn string = fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s",
	db_config.User,
	db_config.Password,
	db_config.EntryPort,
	db_config.Name,
)

func GetStaffMemberList() []myTypes.TeamMemberInfo {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	rawStaffList, err := db.Query("SELECT * FROM staff_list")

	staffList := []myTypes.TeamMemberInfo{}
	for rawStaffList.Next() {
		var name string
		if err := rawStaffList.Scan(&name); err != nil {
			log.Fatal(err)
		}
	}

	if rawStaffList.Err() != nil {
		log.Fatal(err)
	}

	return []myTypes.TeamMemberInfo{}
}
