/*
   Copyright 2019 Septian Wibisono

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

*/
package main

import (
	"database/sql"
	// "errors"
	"log"
	// "path/filepath"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rubenv/sql-migrate"
	pak "github.com/septianw/jas/common"
	"github.com/septianw/jas/types"
)

type database string

func (db database) PingDb(d types.Dbconf) (bool, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.User, d.Pass, d.Host, d.Port, d.Database)

	dbi, err := sql.Open("mysql", dsn)
	if err != nil {
		return false, err
	}

	err = dbi.Ping()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (db database) OpenDb(d types.Dbconf) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		d.User, d.Pass, d.Host, d.Port, d.Database)

	database, err := sql.Open("mysql", dsn)

	// fmt.Printf("\n%+v   %+v\n", d, database)

	return database, err
}

func (db database) Migrate(location string, d types.Dbconf) bool {
	fmt.Printf("\n%+v  %+v\n", "file://"+location, d)
	var reval bool = false
	migration := &migrate.FileMigrationSource{
		Dir: location,
	}

	dbase, err := db.OpenDb(d)
	defer dbase.Close()
	fmt.Printf("\n\nmigrate: %+v\n\n", migration)
	pak.ErrHandler(err)

	version, err := migrate.Exec(dbase, "mysql", migration, migrate.Up)
	pak.ErrHandler(err)
	if err == nil {
		reval = true
	}
	log.Printf("DB version :  %d", version)
	log.Printf("DB migration succeed :  %v", reval)
	// // Read migrations from /home/mattes/migrations and connect to a local postgres database.
	// dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s", d.User, d.Pass, d.Host, d.Port, d.Database)
	// fmt.Println(dsn)
	// m, err := migrate.New("file://"+location, dsn)
	// if err != nil {
	// 	log.Fatal("New migration ", err)
	// 	reval = false
	// }

	// // Migrate all the way up ...

	// if err := m.Up(); err != nil {
	// 	v, dirty, e := m.Version()
	// 	log.Printf("\n version: %+v, dirty: %+v, err: %+v\n", v, dirty, e)

	// 	log.Fatal("Migrating ", err)
	// 	reval = false
	// }
	// v, dirty, err := m.Version()
	// log.Printf("\n version: %+v, dirty: %+v, err: %+v\n", v, dirty, err)

	return reval
}

// func (db database) SetupDb(location string, d types.Dbconf) bool {
// 	return db.Migrate(location, d)
// fmt.Printf("\n%+v  %+v\n", "file://"+filepath.Join(location, "schema"), d)
// var reval bool = false
// migration := &migrate.FileMigrationSource{
// 	Dir: filepath.Join(location, "schema"),
// }

// dbase, err := db.OpenDb(d)
// pak.ErrHandler(err)

// version, err := migrate.Exec(dbase, "mysql", migration, migrate.Up)
// pak.ErrHandler(err)
// if err == nil {
// 	reval = true
// }
// log.Printf("DB version :  %d", version)

// return reval
// fmt.Printf("\n%+v %+v\n", "file://"+filepath.Join(location, "schema"), d)
// var reval bool = true
// // Read migrations from /home/mattes/migrations and connect to a local postgres database.
// dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s", d.User, d.Pass, d.Host, d.Port, d.Database)
// fmt.Println(dsn)
// m, err := migrate.New("file://"+filepath.Join(location, "schema"), dsn)
// if err != nil {
// 	log.Fatal(err)
// 	reval = false
// }

// // Migrate all the way up ...
// if err := m.Up(); err != nil {
// 	v, dirty, e := m.Version()
// 	log.Printf("\n version: %+v, dirty: %+v, err: %+v\n", v, dirty, e)
// 	log.Fatal(err)
// 	reval = false
// }
// v, dirty, err := m.Version()
// log.Printf("\n version: %+v, dirty: %+v, err: %+v\n", v, dirty, err)

// return reval
// }

var Database database
