package database

import (
	"back_project/structure"
	"database/sql"
)

func NewReviewStore(db *sql.DB) *ReviewStore {
	return &ReviewStore{
		db,
	}
}

type ReviewStore struct {
	*sql.DB
}

func (s *ReviewStore) AddReview(item structure.Review, id_user int) (int, error) {

	res, err := s.DB.Exec("INSERT INTO reviews (id_shop, id_user, rating, comment) VALUES (?, ?, ?, ?)", item.IDShop, id_user, item.Rating, item.Comment)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *ReviewStore) GetReviewByShop(id_shop int) ([]structure.Review, error) {
	var reviews []structure.Review

	rows, err := s.DB.Query("SELECT id_review, id_shop, id_user, rating, comment from reviews WHERE id_shop = ?", id_shop)
	if err != nil {
		return []structure.Review{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var review structure.Review
		if err = rows.Scan(
			&review.IDReview,
			&review.IDShop,
			&review.IDUser,
			&review.Rating,
			&review.Comment,
		); err != nil {
			return []structure.Review{}, err
		}
		reviews = append(reviews, review)
	}

	if err = rows.Err(); err != nil {
		return []structure.Review{}, err
	}

	return reviews, nil
}

func (s *ReviewStore) UpdateReview(id_review int, updated_review structure.Review) error {
	sqlStatement := `UPDATE reviews SET
		rating = ?,
		comment = ?
	WHERE id_review = ?
	`
	_, err := s.DB.Exec(sqlStatement, updated_review.Rating, updated_review.Comment, id_review)

	if err != nil {
		return err
	}
	return nil
}

func (s *ReviewStore) DeleteReview(id_review int) error {
	_, err := s.DB.Exec("DELETE from reviews WHERE id_review = ?", id_review)

	if err != nil {
		return err
	}
	return nil
}
