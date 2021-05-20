package db

import (
	"database/sql"
	"time"

	"github.com/cauabernardino/bucket-list/models"
)

// GetItems returns all items in database
func (db Database) GetItems() (*models.ItemList, error) {
	list := &models.ItemList{}

	rows, err := db.Conn.Query("SELECT * FROM items ORDER BY created_at desc;")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var item models.Item

		if err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Creation,
		); err != nil {
			return nil, err
		}

		list.Items = append(list.Items, item)
	}

	return list, nil
}

// NewItem handles the insertion of a new item in the database
func (db Database) NewItem(item *models.Item) error {
	var id string
	var createdAt time.Time

	query := `INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id, created_at;`

	if err := db.Conn.QueryRow(query, item.Name, item.Description).Scan(&id, &createdAt); err != nil {
		return nil
	}

	item.ID = id
	item.Creation = createdAt

	return nil
}

// SearchItemByID handles the searching of an item by its ID
func (db Database) SearchItemByID(itemID string) (models.Item, error) {
	item := models.Item{}

	query := `SELECT * FROM items WHERE id = $1;`

	row := db.Conn.QueryRow(query, itemID)

	switch err := row.Scan(&item.ID, &item.Name, &item.Description, &item.Creation); err {
	case sql.ErrNoRows:
		return item, ErrNoMatch
	default:
		return item, err
	}
}

// DeleteItem handles the deletion of an item by its ID
func (db Database) DeleteItem(itemID string) error {
	query := `DELETE FROM items WHERE id = $1;`

	_, err := db.Conn.Exec(query, itemID)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

// UpdateItem handles the update of an item by its ID
func (db Database) UpdateItem(itemID string, itemData models.Item) (models.Item, error) {
	item := models.Item{}
	query := `UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`

	if err := db.Conn.QueryRow(
		query, itemData.Name, itemData.Description, itemID,
	).Scan(
		&item.ID, &item.Name, &item.Description, &item.Creation,
	); err != nil {
		if err == sql.ErrNoRows {
			return item, ErrNoMatch
		}
		return models.Item{}, err
	}

	return item, nil
}
