package database

import (
	"back_project/helper"
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

func (shop_store *ShopStore) AddShopAndUser(new_shop structure.NewShopAndUser) error {
	hashPassword, _ := helper.HashPassword(new_shop.Password)

	new_shop.Password = hashPassword
	res, err := shop_store.DB.Exec("INSERT INTO users (firstName, lastName, phone, email, password, role) VALUES (?, ?, ?, ?, ?, ?)", new_shop.FirstName, new_shop.LastName, new_shop.Phone, new_shop.Email, new_shop.Password, "trader")
	if err != nil {
		return err
	}
	id_user, err := res.LastInsertId()
	if err != nil {
		return err
	}
	res, err = shop_store.DB.Exec(
		"INSERT INTO shops (name, zip_code, city,address, country, phone, email, description, id_user) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		new_shop.Name,
		new_shop.ZipCode,
		new_shop.City,
		new_shop.Address,
		// new_shop.Lat,
		// new_shop.Long,
		new_shop.Country,
		new_shop.Phone,
		new_shop.Email,
		new_shop.Description,
		id_user)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	for _, element := range new_shop.KindID {
		_, err := shop_store.DB.Exec(
			"INSERT INTO shop_kind (id_shop, id_kind) VALUES (?,?)",
			id,
			element)
		if err != nil {
			return err
		}
	}
	return nil
}

// func (shop_store *ShopStore) AddShop(new_shop structure.NewShop, id_user int) (int, error) {
// 	res, err := shop_store.DB.Exec(
// 		"INSERT INTO shops (name, zip_code, city, country, phone, email, description, id_user) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
// 		new_shop.Name,
// 		new_shop.ZipCode,
// 		new_shop.City,
// 		// new_shop.Lat,
// 		// new_shop.Long,
// 		new_shop.Country,
// 		new_shop.Phone,
// 		new_shop.Email,
// 		new_shop.Description,
// 		id_user)
// 	if err != nil {
// 		return 0, err
// 	}

// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		return 0, err
// 	}
// 	for _, element := range new_shop.KindID {
// 		_, err := shop_store.DB.Exec(
// 			"INSERT INTO shop_kind (id_shop, id_kind) VALUES (?,?)",
// 			id,
// 			element)
// 		if err != nil {
// 			return 0, err
// 		}
// 	}
// 	return int(id), nil
// }

func (shop_store *ShopStore) GetAllShopByKindAndCity(id_kind int, city string) ([]structure.Shop, error) {
	var shops []structure.Shop
	rows, err := shop_store.DB.Query("SELECT shops.id_shop, name, address, zip_code, city, latitude, longitude, country, phone, email, description, id_user FROM shops  INNER JOIN shop_kind on shop_kind.id_shop = shops.id_shop WHERE shop_kind.id_kind = ? AND city = ? ", id_kind, city)
	if err != nil {
		return []structure.Shop{}, err
	}
	for rows.Next() {
		var shop structure.Shop
		if err = rows.Scan(
			&shop.ID,
			&shop.Name,
			&shop.Address,
			&shop.ZipCode,
			&shop.City,
			&shop.Lat,
			&shop.Long,
			&shop.Country,
			&shop.Phone,
			&shop.Email,
			&shop.Description,
			&shop.UserID); err != nil {
			return []structure.Shop{}, err
		}
		shops = append(shops, shop)
	}
	if err = rows.Err(); err != nil {
		return []structure.Shop{}, err

	}
	defer rows.Close()

	return shops, nil
}

func (shop_store *ShopStore) DeleteShop(id_shop int) error {

	_, err := shop_store.DB.Exec("DELETE FROM shops WHERE id_shop = ?", id_shop)
	if err != nil {
		return err
	}
	return nil
}

func (shop_store *ShopStore) UpdateShop(id_shop int, updated_shop structure.Shop) error {
	sqlStatement := `UPDATE shops SET 
	name = ?,
	zip_code = ?,
	city = ?,
	latitude = ?
	longitude = ?
	country = ?,
	phone = ?,
	email = ?,
	description = ? 
	WHERE id_shop = ?`

	_, err := shop_store.DB.Exec(sqlStatement,
		updated_shop.Name,
		updated_shop.ZipCode,
		updated_shop.City,
		updated_shop.Lat,
		updated_shop.Long,
		updated_shop.Country,
		updated_shop.Phone,
		updated_shop.Email,
		updated_shop.Description,
		id_shop,
	)
	if err != nil {
		return err
	}
	return nil
}

func (shop_store *ShopStore) GetAllShopByUser(id_user int) ([]structure.Shop, error) {
	var shops []structure.Shop

	rows, err := shop_store.DB.Query("SELECT id_shop, name, zip_code, city, latitude, longitude, country, phone, email, description, id_user FROM shops where id_user = ?", id_user)
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

func (shop_store *ShopStore) GetAllShopNear(lat float64, long float64, kind string) ([]structure.ShopsNearReturn, error) {
	var shops []structure.ShopsNearReturn
	rows, err := shop_store.DB.Query("SELECT id_shop, shops.name, address, zip_code, city, latitude, longitude, country, phone, email, description, ST_Distance_Sphere( point (?, ?), point(longitude, latitude)) * .000621371192 AS distance_in_miles FROM shops INNER JOIN kinds ON shops.id_kind = kinds.id_kind WHERE kinds.name = ? having distance_in_miles <= 15 order by distance_in_miles asc", long, lat, kind)
	if err != nil {
		return []structure.ShopsNearReturn{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var shop structure.ShopsNearReturn
		if err = rows.Scan(
			&shop.ID,
			&shop.Name,
			&shop.Address,
			&shop.ZipCode,
			&shop.City,
			&shop.Lat,
			&shop.Long,
			&shop.Country,
			&shop.Phone,
			&shop.Email,
			&shop.Description,
			&shop.DistanceInMiles,
		); err != nil {
			return []structure.ShopsNearReturn{}, err
		}
		shops = append(shops, shop)
	}

	if err = rows.Err(); err != nil {
		return []structure.ShopsNearReturn{}, err
	}

	return shops, nil
}

func (shop_store *ShopStore) GetShopById(id int) (structure.ShopRO, error) {
	var shop structure.ShopRO
	var openings []structure.ShowOpening
	var benefits []structure.BenefitRO
	var reviews []structure.ReviewRO

	var sqlStatement = "SELECT id_shop,name,address,zip_code,city, latitude, longitude, country,phone,email,description FROM shops WHERE id_shop = ?"
	rows := shop_store.DB.QueryRow(sqlStatement, id)
	err := rows.Scan(
		&shop.ID,
		&shop.Name,
		&shop.Address,
		&shop.ZipCode,
		&shop.City,
		&shop.Lat,
		&shop.Long,
		&shop.Country,
		&shop.Phone,
		&shop.Email,
		&shop.Description)
	if err != nil {
		return structure.ShopRO{}, err
	}

	row, err := shop_store.DB.Query("SELECT open, close, id_day FROM openings where id_shop = ?", id)
	if err != nil {
		return structure.ShopRO{}, err
	}
	defer row.Close()

	for row.Next() {
		var opening structure.ShowOpening
		if err = row.Scan(
			&opening.Open,
			&opening.Close,
			&opening.DayID); err != nil {
			return structure.ShopRO{}, err
		}
		openings = append(openings, opening)
	}
	row, err = shop_store.DB.Query("SELECT id_benefit, name, description, duration, price FROM benefits where id_shop = ?", id)
	if err != nil {
		return structure.ShopRO{}, err
	}
	defer row.Close()

	for row.Next() {
		var benefit structure.BenefitRO
		if err = row.Scan(
			&benefit.IDBenefit,
			&benefit.Name,
			&benefit.Description,
			&benefit.Duration,
			&benefit.Price); err != nil {
			return structure.ShopRO{}, err
		}
		benefits = append(benefits, benefit)
	}
	row, err = shop_store.DB.Query("SELECT id_review, rating, comment FROM reviews where id_shop = ?", id)
	if err != nil {
		return structure.ShopRO{}, err
	}
	defer row.Close()

	for row.Next() {
		var review structure.ReviewRO
		if err = row.Scan(
			&review.IDReview,
			&review.Rating,
			&review.Comment); err != nil {
			return structure.ShopRO{}, err
		}
		reviews = append(reviews, review)
	}
	shop.Reviews = reviews
	shop.Benefits = benefits
	shop.Openings = openings

	return shop, err
}
