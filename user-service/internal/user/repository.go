package user

import (
	"database/sql"
	"go-microservice/user-service/internal/common"
	"go-microservice/user-service/internal/database"
)

func CreateUser(data User) error {
	// Start transaction
	tx, _ := database.DB.BeginTx(database.Ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	defer tx.Rollback()

	query, err := tx.PrepareContext(database.Ctx, "INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	common.LogIfError(err)
	defer query.Close()

	_, err = query.ExecContext(database.Ctx, data.Name, data.Email, data.Password)
	if err != nil {
		common.LogIfError(err)
		return err
	}

	err = tx.Commit()
	common.LogIfError(err)

	return nil
}

func FindById(id int) (User, error) {
	var user User

	query, err := database.DB.Prepare("SELECT * FROM users WHERE id = ?")
	common.LogIfError(err)

	defer query.Close()

	err = query.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	common.LogIfError(err)
	// TODO: Find user
	return user, err
}
