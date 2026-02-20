package data

import (
	"log"
	db "wordGame/pkg"
)

type Word struct {
	Wordindex int
	Hint      string
	Answer    string
}

func GetWords() []Word {
	db.DataBase()
	defer db.DB.Close()
	var words []Word

	//READ
	rows, err := db.DB.Query("SELECT * FROM capitals")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var w Word
		if err := rows.Scan(&w.Wordindex, &w.Hint, &w.Answer); err != nil {
			log.Fatal(err)
		}
		words = append(words, w)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return words

}
