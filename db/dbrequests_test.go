package db

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func Equal(a, b []Station) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestDbase_GetAllStans(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"stan_id", "dor_kod", "st_kod", "name", "flag"}).AddRow(1, 1, 1, "kek", "flg")
	//mock.ExpectRollback()
	mock.ExpectQuery("^select stan_id, dor_kod, st_kod, name, flag from gredit_schema.stan$").WillReturnRows(rows)

	dbb := Dbase{db}

	var expctd []Station

	exp := Station{
		Stan_id: 1,
		Dor_kod: 1,
		St_kod:  1,
		Name:    "kek",
		Flag:    "flg",
	}

	expctd = append(expctd, exp)

	tstcs, err := dbb.GetAllStans()

	if !Equal(tstcs, expctd) {
		t.Errorf("Should have %v, got %v", expctd, tstcs)
	}
}

func TestDbase_GetStans(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"stan_id", "dor_kod", "st_kod", "name", "flag"}).AddRow(1, 1, 1, "kek", "flg")
	//mock.ExpectRollback()
	mock.ExpectQuery("^select (.+) from gredit_schema.stan where dor_kod = 1$").WillReturnRows(rows)

	dbb := Dbase{db}

	var expctd []Station

	exp := Station{
		Stan_id: 1,
		Dor_kod: 1,
		St_kod:  1,
		Name:    "kek",
		Flag:    "flg",
	}

	expctd = append(expctd, exp)

	tstcs, err := dbb.GetStans(1, "")

	if !Equal(tstcs, expctd) {
		t.Errorf("Should have %v, got %v", expctd, tstcs)
	}
}


func BenchmarkDbase_GetAllStans(b *testing.B) {

	b.N = 100000
	for i := 0; i < b.N; i++ {
		db, mock, err := sqlmock.New()

		if err != nil {
			b.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"stan_id", "dor_kod", "st_kod", "name", "flag"}).AddRow(1, 1, 1, "kek", "flg")
		//mock.ExpectRollback()
		mock.ExpectQuery("^select stan_id, dor_kod, st_kod, name, flag from gredit_schema.stan$").WillReturnRows(rows)

		dbb := Dbase{db}

		dbb.GetAllStans()
	}
}
