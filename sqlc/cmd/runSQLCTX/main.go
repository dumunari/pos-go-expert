package main

import (
	"context"
	"database/sql"
	"fmt"

	"sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

// NewCourseDB initializes a new CourseDB instance with the provided database connection.
// It wraps the db.Queries to provide additional methods for course and category operations.
func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

// CourseParams and CategoryParams are structures that define the parameters
// required for creating a course and a category, respectively.
type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

// callTx is a helper function that starts a transaction, executes the provided function with the Queries,
// and commits the transaction if successful. If an error occurs, it rolls back the transaction.
func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	// Start a new transaction
	// Use BeginTx to allow for more control over the transaction options
	// such as isolation level or read-only mode.
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Create a new Queries instance that uses the transaction
	// This allows us to execute queries within the transaction context.
	q := db.New(tx)
	// Call the provided function with the Queries instance
	err = fn(q)
	// If the function returns an error, we roll back the transaction.
	if err != nil {
		// Rollback the transaction and return the error
		// If the rollback itself fails, we return a wrapped error.
		// This ensures that we do not lose the original error context.
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}
	// If the function executed successfully, we commit the transaction.
	// This makes all changes made during the transaction permanent.
	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	// Call the transaction helper function with the context and a function
	// that performs the necessary database operations.
	// This function will create a category and a course within a single transaction.
	// If any operation fails, the transaction will be rolled back.
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		// Create the category with the provided parameters.
		// This operation is part of the transaction.
		// If it fails, the transaction will be rolled back.
		// If the category already exists, it will not be created again.
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			// If creating the category fails, we return the error.
			// // The transaction will be rolled back automatically by the callTx function.
			// This ensures that the database remains in a consistent state.
			return err
		}
		// Create the course with the provided parameters.
		// This operation is also part of the transaction.
		// If it fails, the transaction will be rolled back.
		// If the course already exists, it will not be created again.
		// The course is associated with the category created above.
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})
		if err != nil {
			// If creating the course fails, we return the error.
			// The transaction will be rolled back automatically by the callTx function.
			// This ensures that the database remains in a consistent state.
			return err
		}
		// If both operations succeed, we return nil to indicate success.
		// The transaction will be committed by the callTx function.
		return nil
	})
	if err != nil {
		// If an error occurs during the transaction, we return it.
		// This could be due to a failure in creating the category or the course.
		// The transaction will be rolled back automatically by the callTx function.
		return err
	}
	// If the transaction was successful, we return nil to indicate success.
	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	// courseArgs := CourseParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Go",
	// 	Description: sql.NullString{String: "Go Course", Valid: true},
	// 	Price:       100.0,
	// }

	// categoryArgs := CategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend Course", Valid: true},
	// }

	// courseDB := NewCourseDB(dbConn)
	// err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	// if err != nil {
	// 	panic(err)
	// }

	queries := db.New(dbConn)

	// List all courses and their categories
	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %.2f",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}
}
