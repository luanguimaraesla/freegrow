package gadget

import (
	"fmt"

	"github.com/luanguimaraesla/freegrow/internal/database"
	"github.com/luanguimaraesla/freegrow/internal/log"
	"go.uber.org/zap"
)

// insertGadget inserts the gadget in the database
func insertGadget(gadget *Gadget) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlStatement := `INSERT INTO gadgets (gadget_uuid, user_id, enabled) VALUES ($1, $2, $3) RETURNING gadget_uuid`

	if err := db.QueryRow(
		sqlStatement,
		gadget.UUID,
		gadget.UserID,
		gadget.Enabled,
	).Scan(&gadget.UUID); err != nil {
		log.L.Fatal("unable to execute the query", zap.Error(err))
	}

	log.L.Info("created gadget record", zap.String("gadget_uuid", gadget.UUID))

	return nil
}

// getGadgetByUUID gets only one gadget from the DB by its gadget_uuid
func getGadgetByUUID(UUID string) (*Gadget, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sqlStatement := `SELECT gadget_uuid,user_id,enabled FROM gadgets WHERE gadget_uuid=$1`

	gadget := New()

	// execute the sql statement
	row := db.QueryRow(sqlStatement, UUID)

	// unmarshal the row object to gadget
	if err := row.Scan(&gadget.UUID, &gadget.UserID, &gadget.Enabled); err != nil {
		return nil, fmt.Errorf(
			"unable to execute `%s` with gadget_uuid %s: %v",
			sqlStatement, UUID, err,
		)
	}

	return gadget, nil
}

// getAllGadgets returns a list with the all gadgets registered
func getAllGadgets() ([]*Gadget, error) {
	var gadgets []*Gadget

	db, err := database.Connect()
	if err != nil {
		return gadgets, err
	}

	defer db.Close()

	sqlStatement := `SELECT gadget_uuid,user_id,enabled FROM gadgets`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return gadgets, fmt.Errorf(
			"unable to execute `%s`: %v",
			sqlStatement, err,
		)
	}

	defer rows.Close()

	for rows.Next() {
		gadget := New()

		if err := rows.Scan(&gadget.UUID, &gadget.UserID, &gadget.Enabled); err != nil {
			return gadgets, fmt.Errorf("unable to scan a gadget row: %v", err)
		}

		gadgets = append(gadgets, gadget)
	}

	return gadgets, err
}

// updateGadget changes the gadget row in the database
// according to the new given User model
func updateGadget(gadget *Gadget) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlStatement := `UPDATE gadgets SET enabled=$2 WHERE gadget_uuid=$1`

	res, err := db.Exec(sqlStatement, gadget.UUID, gadget.Enabled)
	if err != nil {
		return fmt.Errorf(
			"unable to execute `%s` with gadget_uuid `%s`: %v",
			sqlStatement, gadget.UUID, err,
		)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed checking the affected rows: %v", err)
	}

	log.L.Info("gadget rows updated", zap.Int64("total", rowsAffected))

	return nil
}

// deleteGadget deletes an gadget from DB
func deleteGadget(UUID string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlStatement := `DELETE FROM gadgets WHERE gadget_uuid=$1`

	res, err := db.Exec(sqlStatement, UUID)
	if err != nil {
		return fmt.Errorf(
			"unable to execute `%s` with gadget_uuid `%s`: %v",
			sqlStatement, UUID, err,
		)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed checking the affected rows: %v", err)
	}

	log.L.Info("gadget rows deleted", zap.Int64("total", rowsAffected))

	return nil
}

// getUserGadgets returns a list with the all gadgets registered for an
// specific user
func getUserGadgets(ID int64) ([]*Gadget, error) {
	var gadgets []*Gadget

	db, err := database.Connect()
	if err != nil {
		return gadgets, err
	}

	defer db.Close()

	sqlStatement := `SELECT gadget_uuid,user_id,enabled FROM gadgets WHERE user_id=$1`

	rows, err := db.Query(sqlStatement, ID)
	if err != nil {
		return gadgets, fmt.Errorf(
			"unable to execute `%s`: %v",
			sqlStatement, err,
		)
	}

	defer rows.Close()

	for rows.Next() {
		g := New()

		if err := rows.Scan(&g.UUID, &g.UserID, &g.Enabled); err != nil {
			return gadgets, fmt.Errorf("unable to scan a gadget row: %v", err)
		}

		gadgets = append(gadgets, g)
	}

	return gadgets, err
}
