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

func GetStaffMemberList() []myTypes.StaffMember {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	var staffList []myTypes.StaffMember

	rawStaffList, err := db.Query("SELECT * FROM staff_list")
	for rawStaffList.Next() {
		var staffMember myTypes.StaffMember
		err := rawStaffList.Scan(&staffMember.UserID, &staffMember.Username, &staffMember.ProfilePictureLink, &staffMember.Position, &staffMember.PositionPrettyName)
		if err != nil {
			log.Fatal(err)
		}
		staffList = append(staffList, staffMember)
	}

	if rawStaffList.Err() != nil {
		log.Fatal(err)
	}

	return staffList
}
