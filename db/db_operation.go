package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/neostefan/diet-backend/models"
)

type ConditionType string

const (
	DIABETES ConditionType = "DIABETES"
	ULCER    ConditionType = "ULCER"
)

//TODO create a function to check for an existing user in the database to check registrations

// creates the users table
func CreateUsersTable(db *sql.DB) {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY UNIQUE NOT NULL, firstName var(255) NOT NULL, lastName var(255) NOT NULL, password var(255) NOT NULL, age int NOT NULL, condition var(255))`)

	if err != nil {
		log.Fatalf("\n Unable to create users table: %v", err.Error())
	}

	if _, err := stmt.Exec(); err != nil {
		log.Fatalf("\n Unable to execute create table statement: %v", err.Error())
	}

	defer stmt.Close()
}

// adds a user to the users table
func AddUser(db *sql.DB, u *models.User) error {
	e := Error{}
	stmt, err := db.Prepare(`INSERT INTO users(firstName, lastName, password, age, condition) VALUES(?, ?, ?, ?, ?)`)

	if err != nil {
		e.errMsg = fmt.Sprintf("Unable to register user: %v", err.Error())
		return e
	}

	_, errE := stmt.Exec(u.FirstName, u.LastName, u.Password, u.Age, u.Condition)

	if errE != nil {
		e.errMsg = fmt.Sprintf("Unable to execute register db statement: %v", errE.Error())
		return e
	}

	defer stmt.Close()

	return nil
}

// gets a user from the users table
func GetUser(db *sql.DB, firstName, lastName string) (models.User, error) {
	stmt, err := db.Prepare(`SELECT * FROM users WHERE(firstName = ? AND lastName = ?)`)
	u := models.User{}
	e := Error{}

	if err != nil {
		e.errMsg = fmt.Sprintf("Unable to prepare get users query: %v", err.Error())
		return u, e
	}

	rows, err := stmt.Query(firstName, lastName)

	if err != nil {
		e.errMsg = fmt.Sprintf("Unable to execute prepared get users query: %v", err.Error())
		return u, e
	}

	for rows.Next() {
		err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Password, &u.Age, &u.Condition)
		if err != nil {
			e.errMsg = fmt.Sprintf("Unable to scan results from get users query: %v", err.Error())
			return u, e
		}
	}

	defer stmt.Close()
	return u, nil
}

func GetUserById(db *sql.DB, id int) (models.User, error) {
	stmt, err := db.Prepare(`SELECT * FROM users WHERE(Id = ?)`)
	u := models.User{}
	e := Error{}

	if err != nil {
		e.errMsg = fmt.Sprintf("Unable to prepare get users query: %v", err.Error())
		return u, e
	}

	rows, err := stmt.Query(id)

	if err != nil {
		e.errMsg = fmt.Sprintf("Unable to execute prepared get users query: %v", err.Error())
		return u, e
	}

	for rows.Next() {
		err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Password, &u.Age, &u.Condition)
		if err != nil {
			e.errMsg = fmt.Sprintf("Unable to scan results from get users query: %v", err.Error())
			return u, e
		}
	}

	defer stmt.Close()
	return u, nil
}

//TODO check that the foreign key constraint works

// creates the algorithm's diet recommendations table for a given day
func CreateRecommendationsTable(db *sql.DB) error {
	err := Error{}
	stmt, errP := db.Prepare(`CREATE TABLE IF NOT EXISTS recommendations (id INTEGER PRIMARY KEY UNIQUE NOT NULL, carbs var(255), proteins var(255), oils var(255), vegetables var(255), beverages var(255), fruits VAR(255), userId INTEGER)`)

	if errP != nil {
		err.errMsg = "unable to create users table: " + errP.Error()
		return err
	}

	if _, errE := stmt.Exec(); errE != nil {
		err.errMsg = "unable to execute create table statement: " + errE.Error()
		return err
	}

	defer stmt.Close()
	return nil
}

// add diet recommendation
func AddRecommendation(db *sql.DB, meal *models.Meal) error {
	err := Error{}
	stmt, errP := db.Prepare(`INSERT INTO recommendations(carbs, proteins, oils, vegetables, beverages, fruits) VALUES(?, ?, ?, ?, ?, ?)`)

	if errP != nil {
		err.errMsg = "unable to execute prepared recommendation query " + errP.Error()
		return err
	}

	if _, errE := stmt.Exec(meal.Carbs, meal.Proteins, meal.Oils, meal.Vegetables, meal.Beverages, meal.Fruits); errE != nil {
		err.errMsg = "unable to insert meal recommendation " + errE.Error()
		return err
	}

	defer stmt.Close()
	return nil
}

// get the meals recommended earlier in the day
func GetMealRecommendations(db *sql.DB, userId int32) ([]*models.Meal, error) {
	meals := []*models.Meal{}
	err := Error{}

	stmt, errP := db.Prepare(`SELECT carbs, proteins, vegetables, oils, beverages, fruits FROM recommendations WHERE(userId = ?)`)

	if errP != nil {
		err.errMsg = "unable to prepare select recommendations query: " + errP.Error()
		return nil, err
	}

	rows, errQ := stmt.Query(userId)

	if errQ != nil {
		err.errMsg = "unable to execute select recommendations query: " + errQ.Error()
		return nil, err
	}

	for rows.Next() {
		meal := models.Meal{}
		errS := rows.Scan(&meal.Carbs, &meal.Proteins, &meal.Vegetables, &meal.Oils, &meal.Oils, &meal.Beverages, &meal.Fruits)
		if errS != nil {
			err.errMsg = "error with reading recommendations: " + errS.Error()
			return nil, err
		}
	}

	defer stmt.Close()

	return meals, nil
}
