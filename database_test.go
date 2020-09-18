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
	"testing"

	// "database/sql"
	// "path/filepath"
	"os"
	"path"
	"reflect"

	"github.com/septianw/jas/types"
)

var D = types.Dbconf{
	"mysql",
	"localhost",
	3306,
	"asep",
	"dummypass",
	"ipointtest",
}

func TestPingDb(t *testing.T) {
	ok, err := Database.PingDb(D)
	if !ok {
		t.Fail()
		t.Logf("err : %+v", err)
	}
}

func TestOpenDb(t *testing.T) {
	db, err := Database.OpenDb(D)
	if err != nil {
		t.Fail()
		t.Logf("err : %+v", err)
	}
	t.Logf("db : |%+v|", reflect.TypeOf(db).String())
	if reflect.TypeOf(db).String() != "*sql.DB" {
		t.Fail()
		t.Logf("type : %+v", reflect.TypeOf(db).String())
	}
	// db.(sql.DB)
}

// func TestSetupDb(t *testing.T) {
// 	ok := Database.SetupDb("/home/asep/workspace/go/github.com/septianw/jas-mysql/test", D)
// 	if !ok {
// 		t.Fail()
// 		t.Logf("ok : %+v", ok)
// 	}
// }

func TestMigrate(t *testing.T) {
	pwd, _ := os.Getwd()
	ok := Database.Migrate(path.Join(pwd, "test", "schema"), D)
	if !ok {
		t.Fail()
		t.Logf("ok : %+v", ok)
	}
}
