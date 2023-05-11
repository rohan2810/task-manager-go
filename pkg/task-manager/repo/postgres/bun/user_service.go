package bun

import (
	"context"
	"github.com/uptrace/bun"
	"rohan.com/task-manager/pkg/task-manager/repo/models"
)

func (db *DBClient) CreateUserTable(ctx context.Context) error {
	_, err := db.DB.NewCreateTable().
		IfNotExists().Model(&models.User{}).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBClient) CreateUser(ctx context.Context, u *models.User) (*models.User, error) {
	err := db.DB.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := db.DB.NewInsert().Model(u).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (db *DBClient) GetUser(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	err := db.DB.NewSelect().
		Model(user).
		Where("username = ?", username).
		Scan(ctx)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DBClient) DeleteUser(ctx context.Context, username string) error {
	_, err := db.DB.NewDelete().
		Model((*models.User)(nil)).
		Where("username = ?", username).Exec(ctx)

	if err != nil {
		return err
	}
	return nil
}

func (db *DBClient) ListUser(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	err := db.DB.NewSelect().
		Model(&users).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *DBClient) UpdateUser(ctx context.Context, u *models.User) (*models.User, error) {

	_, err := db.DB.NewUpdate().
		Model(u).
		Set("fullname = ?", u.Fullname).
		Set("email = ?", u.Email).
		Set("role = ?", u.Role).
		Where("username = ?", u.Username).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	// Get the updated user from the database
	user := &models.User{}
	err = db.DB.NewSelect().
		Model(user).
		Where("username = ?", u.Username).
		Limit(1).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}
