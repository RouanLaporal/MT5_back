package database

import (
	"back_project/structure"
	"database/sql"
)

func NewShopStore(db *sql.DB) *ShopStore {
	return &ShopStore{
		db,
	}
}

type ShopStore struct {
	*sql.DB
}

func (shop_store *ShopStore) AddShop(new_shop structure.Shop) (int, error) {
	res, err := shop_store.DB.Exec(
		"INSERT INTO shops (name, zip_code, city, latitude, longitude, country, phone, email, description, id_kind, id_user) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		new_shop.Name,
		new_shop.ZipCode,
		new_shop.City,
		new_shop.Lat,
		new_shop.Long,
		new_shop.Country,
		new_shop.Phone,
		new_shop.Email,
		new_shop.Description,
		new_shop.KindID,
		new_shop.UserID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (shop_store *ShopStore) GetAllShopByKindAndCity(id_kind int, city string) ([]structure.Shop, error) {
	var shops []structure.Shop

	rows, err := shop_store.DB.Query("SELECT id_shop, name, zip_code, city, latitude, longitude, country, phone, email, description, id_kind, id_user FROM shops where id_kind = ? AND city = ? ", id_kind, city)
	if err != nil {
		return []structure.Shop{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var shop structure.Shop
		if err = rows.Scan(
			&shop.ID,
			&shop.Name,
			&shop.ZipCode,
			&shop.City,
			&shop.Lat,
			&shop.Long,
			&shop.Country,
			&shop.Phone,
			&shop.Email,
			&shop.Description,
			&shop.KindID,
			&shop.UserID); err != nil {
			return []structure.Shop{}, err
		}
		shops = append(shops, shop)
	}

	if err = rows.Err(); err != nil {
		return []structure.Shop{}, err
	}

	return shops, nil
}

func (shop_store *ShopStore) DeleteShop(id_shop int) error {
	_, err := shop_store.DB.Exec("DELETE FROM shops WHERE id_shop = ?", id_shop)
	if err != nil {
		return err
	}
	return nil
}
