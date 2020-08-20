package main

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type table struct {
	tablename string
	feilds    []string
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Error: Please provide the rule Id")
		os.Exit(1)
	}
	var id = os.Args[1]

	var tables = []table{
		{"MpleRule", []string{"rule_id", "group_id", "rule_name", "status"}},
		{"MpleAction", []string{"action_id", "rule_id", "fire_rule_id", "status"}},
		{"MpleMatch", []string{"match_id", "rule_id", "match_name", "status"}},
		{"EventMatch", []string{"event_match_id", "rule_id", "event_match_name", "event_match_type", "status"}},
		{"ExclusionGroup", []string{"rule_id", "exclusion_group_id"}},
		{"MetaRuleMap", []string{"id", "customer_id", "rule_id", "meta_rule_id"}},
		{"MpleEvents", []string{"event_id", "customer_id", "rule_id", "match_id", "event_summary"}},
		{"MpleRuleTest", []string{"description", "group_id", "rule_id"}},
		{"TemplateAction", []string{"action_id", "rule_id", "options"}},
	}

	for _, tab := range tables {
		myQuery(tab, id)
	}
}

type credentials struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

func myQuery(tab table, id string) {
	// Reading credesntials from the file
	data, err := ioutil.ReadFile("cred.json")
	if err != nil {
		log.Fatal(err)
	}

	var cred credentials
	err = json.Unmarshal(data, &cred)
	if err != nil {
		log.Fatal(err)
	}

	var (
		dbUser = cred.DbUser
		dbHost = cred.DbHost
		dbPort = cred.DbPort
		dbName = cred.DbName
		dbPass = decrypt(cred.DbPassword)
	)

	//setting up the database connection
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@"+"tcp"+"("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name := tab.tablename
	flds := tab.feilds
	qry := "select " + strings.Join(flds, ",") + " from " + name + " where rule_id = " + id

	// queryinng the database
	rows, err := db.Query(qry)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var values []string

		for _, v := range tab.feilds {
			err := rows.Scan(&v)
			if err != nil {
				log.Fatal(err)
			}
			values = append(values, v)
		}
		fmt.Println(values)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func decrypt(p string) string {
	decoded, err := hex.DecodeString(p)
	if err != nil {
		log.Fatal(err)
	}
	return string(decoded)
}
