package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

var BLOGS []Blog

type Blog struct {
	Id          int
	Title       string
	Body        string
	Tags_string string
	created_at  string
	updated_at  string
}
type Comment struct {
	Id         int
	Comment    string
	Blog_id    int
	Username   string
	created_at string
	updated_at string
}
type User struct {
	Id              int
	Name            string
	Password_digest string
	Email           string
	created_at      string
	updated_at      string
}

func main() {
	readBlog()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	router.Static("/dist", "./dist")
	// API群
	router.GET("/api/v1/blogs", func(c *gin.Context) {
		bs := make([]*Blog, len(BLOGS))
		for _, b := range BLOGS {
			nb := new(Blog)
			nb.Title = b.Title
			nb.Tags_string = b.Tags_string
			nb.Body = string([]rune(b.Body)[:30])
			bs = append(bs, nb)
		}
		// 一覧の取得
		c.JSON(200, gin.H{"blogs": bs})
	})
	router.GET("/api/v1/blog/:id", func(c *gin.Context) {
		// 記事の取得
		var fb Blog
		for _, b := range BLOGS {
			id, _ := strconv.Atoi(c.Param("id"))
			if b.Id == id {
				fb = b
				break
			}
		}
		if &fb.Id != nil {
			c.JSON(200, gin.H{"blogs": BLOGS})
		} else {
			c.JSON(404, gin.H{"error": "not found"})
		}
	})
	router.POST("/api/v1/blog", func(c *gin.Context) {
		// 登録
		title := c.PostForm("title")
		body := c.PostForm("body")
		tags_string := c.PostForm("tags_string")

		db, err := sql.Open("sqlite3", "./blog.db")
		defer db.Close()
		_, _ := db.Exec(
			"insert into blog(title, body, tags_string) values(?,?,?)",
			title,
			body,
			tags_string,
		)
		if err != nil {
			panic(err)
		}
	})
	router.DELETE("/api/v1/blog/:id", func(c *gin.Context) {
		// 削除
		id := strconv.Atoi(c.Param("id"))
		db, err := sql.Open("sqlite3", "./blog.db")
		defer db.Close()
		_, _ := db.Exec(
			"delete from blog where id = ?",
			id,
		)
		defer rows.Close()
	})
	router.PUT("/api/v1/blog/:id", func(c *gin.Context) {
		// 更新
		id := strconv.Atoi(c.PostForm("id"))
		title := c.PostForm("title")
		body := c.PostForm("body")
		tags_string := c.PostForm("tags_string")

		db, err := sql.Open("sqlite3", "./blog.db")
		defer db.Close()
		_, _ := db.Exec(
			"update blog set title = ?, body = ?, tags_string = ? where id = ?",
			title,
			body,
			tags_string,
			id,
		)
		if err != nil {
			panic(err)
		}
	})
	router.Run()
}

func readBlog() {
	db, err := sql.Open("sqlite3", "./blog.db")
	defer db.Close()
	rows, err := db.Query(
		`SELECT * FROM blog order by updated_at;`,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	var id int
	var title, body, tags_string string
	for rows.Next() {
		rows.Scan(&id, &title, &body, &tags_string)
		BLOGS = append(BLOGS, Blog{Id: id, Title: title, Body: body,
			Tags_string: tags_string})
	}
}
