package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	println("Hello")
	db, err := sqlx.Connect("mysql", "root@tcp(localhost:3306)/demo")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	type Post struct {
		ID    int64
		Title string
	}
	type Author struct {
		ID    int64
		Name  string
		Posts []Post
	}

	dest := []struct {
		ID        int64  `db:"a_id"`
		Name      string `db:"a_name"`
		PostID    int64  `db:"p_id"`
		PostTitle string `db:"p_title"`
	}{}
	set := map[int64]*Author{}
	db.Select(&dest, "SELECT a_id, a_name, p_id, p_title FROM authors LEFT JOIN posts ON p_author=a_id")
	for _, item := range dest {
		auth, ok := set[item.ID]
		if !ok {
			set[item.ID] = &Author{
				ID:    item.ID,
				Name:  item.Name,
				Posts: []Post{{ID: item.PostID, Title: item.PostTitle}},
			}
		} else {
			auth.Posts = append(auth.Posts, Post{ID: item.PostID, Title: item.PostTitle})
		}
	}
	var res []Author
	for _, v := range set {
		res = append(res, *v)
	}
	fmt.Printf("%+v\n", res)
}
