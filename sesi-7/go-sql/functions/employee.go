package functions

import (
	"database/sql"
	"fmt"
	"go-sql/models"

	_ "github.com/lib/pq"
)

type DBRepository struct {
	db *sql.DB
}

func CreateEmployee(db *sql.DB) {
	var employee = models.Employee{}

	sqlStatement := `
		INSERT INTO employees (full_name, email, age, division)
		VALUES ($1, $2, $3, $4)
		RETURNING *;
	`

	err := db.QueryRow(sqlStatement, "Leonardo Siagian", "leonardosiagian2001@gmail.com", 20, "Back End").
		Scan(&employee.Id, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New Data of Employee : %+v\n", employee)
}

func GetEmployees(db *sql.DB) {
	var result = []models.Employee{}

	sqlStatement := `
		SELECT * FROM employees;
	`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee = models.Employee{}
		err = rows.Scan(&employee.Id, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)
		if err != nil {
			panic(err)
		}

		result = append(result, employee)
	}

	fmt.Printf("Employees : %+v\n", result)
}

func UpdateEmployee(db *sql.DB) {
	sqlStatement := `
		UPDATE employees
		SET full_name = $1, email = $2, age = $3, division = $4
		WHERE id = $5
		RETURNING *;
	`
	res, err := db.Exec(sqlStatement, "Leonardo Siagian", "leosiagiannn@gmail.com", 20, "Back End", 1)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Update Successfully", count)
}

func DeleteEmplyoee(db *sql.DB) {
	sqlStatement := `
		DELETE FROM employees
		WHERE id = $1
		RETURNING *;
	`
	res, err := db.Exec(sqlStatement, 1)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Delete Successfully", count)
}
