package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) any

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (any, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, db *sql.DB) *Uow {
	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

// Register adds a new repository factory to the unit of work.
// The factory function should return a repository instance when called with a transaction.
// The name is used to identify the repository.
// If a repository with the same name already exists, it will be overwritten.
func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}

// UnRegister removes a repository factory from the unit of work by its name.
// If the repository does not exist, it will do nothing.
func (u *Uow) UnRegister(name string) {
	delete(u.Repositories, name)
}

// GetRepository retrieves a repository by its name.
func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	// Check if the repository exists
	_, exists := u.Repositories[name]
	if !exists {
		return nil, fmt.Errorf("repository %s not found", name)
	}

	// If a transaction is not already started, begin a new transaction
	// and set it to the unit of work.
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}
	// Retrieve the repository factory and create the repository instance
	repo := u.Repositories[name](u.Tx)
	// If the repository factory returns nil, return an error
	if repo == nil {
		return nil, fmt.Errorf("repository %s factory returned nil", name)
	}
	// Return the repository instance
	// Note: The returned type is `any` to allow flexibility in repository types.
	return repo, nil
}

// Do executes the provided function within a transaction context.
func (u *Uow) Do(ctx context.Context, fn func(Uow *Uow) error) error {
	// Check if a transaction is already started
	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}
	// Start a new transaction
	tx, err := u.Db.BeginTx(ctx, nil)
	// If there is an error starting the transaction, return it
	if err != nil {
		return err
	}
	// Set the transaction
	u.Tx = tx

	// Execute the provided function with the unit of work
	err = fn(u)
	if err != nil {
		// If there is an error, rollback the transaction
		errRb := u.Rollback()
		if errRb != nil {
			// If there is an error during rollback, return both errors
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		// Return the original error
		return err
	}
	// If everything is successful, commit the transaction
	return u.CommitOrRollback()
}

// Rollback rolls back the current transaction if it exists.
func (u *Uow) Rollback() error {
	// Check if there is an active transaction
	if u.Tx == nil {
		// If there is no transaction, return an error
		return errors.New("no transaction to rollback")
	}
	// Attempt to rollback the transaction
	err := u.Tx.Rollback()
	if err != nil {
		// If there is an error during rollback, return it
		return err
	}
	// Set the transaction to nil after rollback
	u.Tx = nil

	// Return nil to indicate successful rollback
	return nil
}

// CommitOrRollback commits the current transaction if it exists, or rolls it back if there is an error.
func (u *Uow) CommitOrRollback() error {
	// Commit the transaction if it exists
	if u.Tx == nil {
		return errors.New("no transaction to commit")
	}
	// Attempt to commit the transaction
	err := u.Tx.Commit()
	if err != nil {
		// If there is an error during commit, rollback the transaction
		errRb := u.Rollback()
		if errRb != nil {
			// If there is an error during rollback, return both errors
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		// Return the original error
		return err
	}
	// Set the transaction to nil after commit
	u.Tx = nil
	// Return nil to indicate successful commit
	return nil
}
