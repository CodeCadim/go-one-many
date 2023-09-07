package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	// Connect to the local demo database provided via docker
	db, err := sqlx.Connect("mysql", "root@tcp(localhost:3306)/demo")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Our simple business model with 2 data-structures and the
	// one-to-many relation : 1 author can have many posts.
	type Post struct {
		ID    int64
		Title string
	}
	type Author struct {
		ID    int64
		Name  string
		Posts []Post
	}

	// We provide an adhoc structure to collect the database results
	dest := []struct {
		ID        int64  `db:"a_id"`
		Name      string `db:"a_name"`
		PostID    int64  `db:"p_id"`
		PostTitle string `db:"p_title"`
	}{}
	db.Select(&dest, "SELECT a_id, a_name, p_id, p_title FROM authors LEFT JOIN posts ON p_author=a_id")

	// The result contains duplicate of author's datas, so we deduplicate
	// them with a map and an iteration on each result to group the posts
	// in the author's Posts field
	set := map[int64]*Author{}
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

	// Then we reassemble a valid business data-structure from our map
	var res []Author
	for _, v := range set {
		res = append(res, *v)
	}

	// And that's it !
	fmt.Printf("%+v\n", res)
}
