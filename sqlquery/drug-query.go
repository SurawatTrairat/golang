package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

var server = "cortex-non-prod-east-asia.database.windows.net"
var port = 1433
var user = "cortexadmin"
var password = "PB4&kwk[0c"
var database = "dev_medication"

func main() {
	//create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	//create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Conn pool fail: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	queryMedicationTest()
}

func queryMedicationTest() (int, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf("SELECT TOP (5) [created_at], [updated_at], [hospital_id], [hospital_code], [drug_name], [display_name], [strength_value], [prescribe_unit_code] FROM [dbo].[medication];")
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		// var hospital_id sql.NullString
		var hospital_id string
		// var hospital_code sql.NullString
		// var drug_name sql.NullString
		// var display_name sql.NullString
		// var strength_value sql.NullFloat64
		// var prescribe_unit_code sql.NullString
		fmt.Print(rows.ColumnTypes())
		fmt.Printf("id: %s", hospital_id)
		// fmt.Printf("hospital_id: %s, hospital_code: %s, drug_name: %s, display_name: %s, strength_value: %f, prescribe_unit_code: %s\n", hospital_id.String, hospital_code.String, drug_name.String, display_name.String, strength_value.Float64, prescribe_unit_code.String)
		count++
	}

	return count, nil
}

func queryMedicationFull() (int, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf("SELECT TOP (5) [created_at], [created_by], [updated_at], [updated_by], [id], [org_id], [hospital_id], [hospital_code], [tmt_id], [tmt_concept_id], [drug_name], [display_name], [display_th], [alias_name], [trade_name], [generic_name], [medication_generic_id], [dosage_value], [dosage_unit_id], [dosage_form_id], [strength_value], [strength_per_prescribe], [strength_unit_id], [strength_display], [dispense_unit_id], [dosage_mg_per_ml], [dosage_mg_per_dispense], [active_search], [remark], [prescribe_unit_code], [prescribe_unit_display], [sale_effective_from], [sale_effective_to] FROM [dbo].[medication];")
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		var created_at string
		var created_by string
		var updated_at string
		var updated_by string
		var id string
		var org_id string
		var hospital_id string
		var hospital_code string
		var tmt_id string
		var tmt_concept_id string
		var drug_name string
		var display_name string
		var display_th string
		var alias_name string
		var trade_name string
		var generic_name string
		var medication_generic_id string
		var dosage_value string
		var dosage_unit_id string
		var dosage_form_id string
		var strength_value string
		var strength_per_prescribe string
		var strength_unit_id string
		var strength_display string
		var dispense_unit_id string
		var dosage_mg_per_ml string
		var dosage_mg_per_dispense string
		var active_search string
		var remark string
		var prescribe_unit_code string
		var prescribe_unit_display string
		var sale_effective_from string
		var sale_effective_to string

		fmt.Printf("created_at: %s, created_by: %s, updated_at: %s, updated_by: %s, id: %s, org_id: %s, hospital_id: %s, hospital_code: %s, tmt_id: %s, tmt_concept_id: %s, drug_name: %s, display_name: %s, display_th: %s, alias_name: %s, trade_name: %s, generic_name: %s, medication_generic_id: %s, dosage_value: %s, dosage_unit_id: %s, dosage_form_id: %s, strength_value: %s, strength_per_prescribe: %s, strength_unit_id: %s, strength_display: %s, dispense_unit_id: %s, dosage_mg_per_ml: %s, dosage_mg_per_dispense: %s, active_search: %s, remark: %s, prescribe_unit_code: %s, prescribe_unit_display: %s, sale_effective_from: %s, sale_effective_t: %s\n", created_at, created_by, updated_at, updated_by, id, org_id, hospital_id, hospital_code, tmt_id, tmt_concept_id, drug_name, display_name, display_th, alias_name, trade_name, generic_name, medication_generic_id, dosage_value, dosage_unit_id, dosage_form_id, strength_value, strength_per_prescribe, strength_unit_id, strength_display, dispense_unit_id, dosage_mg_per_ml, dosage_mg_per_dispense, active_search, remark, prescribe_unit_code, prescribe_unit_display, sale_effective_from, sale_effective_to)
		count++
	}

	return count, nil
}
