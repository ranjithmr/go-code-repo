package main

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

func updatedb(brokers []broker, datacenter string, wg *sync.WaitGroup) error {
	defer wg.Done()
	var (
		dbUser = "root"
		dbHost = "localhost"
		dbPort = "3306"
		dbName = "my_dashboard"
		dbPass = "guessmypassword"
	)

	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@"+"tcp"+"("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		return errors.New("Error : Cannot Establiesh DB Connection")
	}
	defer db.Close()

	stmt, err := db.Prepare("replace br_map (datacenter, broker, fqdn, mapping) Values(?,?,?,?)")
	if err != nil {
		return errors.New("Error : Query Execution failed")
	}
	defer stmt.Close()
	for _, s := range brokers {
		brokername := "os-" + s.Name
		for _, inner := range s.Services {
			if inner.Service == "@value" {
				mapper := inner.Service + ":" + inner.Container.(string)
				stmt.Exec(datacenter, brokername, "", mapper)
			} else {

				innermost, ok := inner.Container.(map[string]interface{})
				if ok {
					for key, value := range innermost {
						containername := key
						var mapper string
						newmap := value.(map[string]interface{})
						for newkey, newvalue := range newmap {
							mapper = newkey + ":" + newvalue.(string)
						}
						stmt.Exec(datacenter, brokername, containername, mapper)
					}
				}
			}
		}
	}
	return nil
}
