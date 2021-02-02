package db

import (
	"errors"
	"fmt"
	"log"
)

type Dstr interface {
	GetAllStans() ([]Station, error)
	GetStans(dorCode int, flag string) ([]Station, error)
}

//GetAllStations send a request and return an array of []Station with all stans
func (db *Dbase) GetAllStans() ([]Station, error) {
	rows, err := db.Query("select stan_id, dor_kod, st_kod, name, flag from gredit_schema.stan")
	if err != nil {
		//panic(err)
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	stations := []Station{}
	for rows.Next() {
		st := Station{}
		err := rows.Scan(&st.Cnsi, &st.Road, &st.Esr, &st.Name, &st.Flag)
		if err != nil {
			fmt.Println(err)
			continue
		}
		stations = append(stations, st)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	//fmt.Println(stations[:3])

	return stations, nil
}

//GetStans func send a request and return array of filtered stans by road and/or by flag
func (db *Dbase) GetStans(dorCode int, flag string) ([]Station, error) {
	var query string

	switch {
	case dorCode == 0 || flag != "":
		query = fmt.Sprintf("select stan_id, dor_kod, st_kod, name, flag from gredit_schema.stan where flag = %s", flag)
	case dorCode != 0 || flag == "":
		query = fmt.Sprintf("select stan_id, dor_kod, st_kod, name, flag from gredit_schema.stan where dor_kod = %v", dorCode)
	case dorCode != 0 || flag != "":
		query = fmt.Sprintf("select stan_id, dor_kod, st_kod, name, flag from gredit_schema.stan where dor_kod = %v and flag = %s", dorCode, flag)
	case dorCode == 0 || flag == "":
		return nil, errors.New("empty parameters")
	}

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	stations := []Station{}
	for rows.Next() {
		st := Station{}
		err := rows.Scan(&st.Cnsi, &st.Road, &st.Esr, &st.Name, &st.Flag)
		if err != nil {
			fmt.Println(err)
			continue
		}
		stations = append(stations, st)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	//fmt.Println(stations[:3])

	return stations, nil
}

//GetNextStans method return list of all stations around the current stan
func (db *Dbase) GetNextStans(cnsi int) ([]Station, error) {
	//var query string

	query := fmt.Sprintf(`with
      end_stans as (select SCG.name_end as name, SCG.esr_end as esr, SCG.cnsi_end as cnsi, SCG.dor_kod_end as road, SCG.len from gredit_schema.set_cnsi_graf SCG where SCG.cnsi_begin = %v),
      strt_stans as (select SCG.name_begin as name, SCG.esr_begin as esr, SCG.cnsi_begin as cnsi, SCG.dor_kod_begin as road, SCG.len from gredit_schema.set_cnsi_graf SCG where SCG.cnsi_end = %v)
select name, esr, cnsi, road, len from end_stans
union
select name, esr, cnsi, road, len  from strt_stans`, cnsi, cnsi)

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	stations := []Station{}
	for rows.Next() {
		st := Station{}
		err := rows.Scan(&st.Cnsi, &st.Road, &st.Esr, &st.Name, &st.Len)
		if err != nil {
			fmt.Println(err)
			continue
		}
		stations = append(stations, st)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stations, nil
}