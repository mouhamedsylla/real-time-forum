package database

import (
	"database/sql"
	"real-time-forum/orm"
)

type PostDB struct {
	Storage *orm.ORM
}

var DbPost = &PostDB{}

func UpdateReaction(db *sql.DB, userId int, postId int, newValue string) error {
	// Début de la transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Vérifiez s'il y a une entrée existante dans ReactionPost
	var oldValue string
	err = tx.QueryRow(`SELECT Value FROM ReactionPost WHERE UserId = ? AND PostId = ?`, userId, postId).Scan(&oldValue)

	switch {
	case err == sql.ErrNoRows:
		// Si aucune entrée, insérer une nouvelle entrée
		_, err = tx.Exec(`INSERT INTO ReactionPost (UserId, PostId, Value) VALUES (?, ?, ?)`, userId, postId, newValue)
		if err != nil {
			return err
		}
		if newValue == "like" {
			_, err = tx.Exec(`UPDATE UserPosts SET Like = Like + 1 WHERE Id = ?`, postId)
		} else if newValue == "dislike" {
			_, err = tx.Exec(`UPDATE UserPosts SET Dislike = Dislike + 1 WHERE Id = ?`, postId)
		}
	case err != nil:
		return err
	default:
		// Si entrée existante, mettre à jour Value
		_, err = tx.Exec(`UPDATE ReactionPost SET Value = ? WHERE UserId = ? AND PostId = ?`, newValue, userId, postId)
		if err != nil {
			return err
		}
		// Mettre à jour les comptes dans la table Post
		if oldValue == "like" {
			_, err = tx.Exec(`UPDATE UserPosts SET Like = Like - 1 WHERE Id = ?`, postId)
		} else if oldValue == "dislike" {
			_, err = tx.Exec(`UPDATE UserPosts SET Dislike = Dislike - 1 WHERE Id = ?`, postId)
		}
		if newValue == "like" {
			_, err = tx.Exec(`UPDATE UserPosts SET Like = Like + 1 WHERE Id = ?`, postId)
		} else if newValue == "dislike" {
			_, err = tx.Exec(`UPDATE UserPosts SET Dislike = Dislike + 1 WHERE Id = ?`, postId)
		}
	}

	if err != nil {
		return err
	}

	// Commit de la transaction
	return tx.Commit()
}
