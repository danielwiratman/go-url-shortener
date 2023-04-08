package handlers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/danielwiratman/go-url-shortener/database"
	"github.com/danielwiratman/go-url-shortener/helper"
	"github.com/gofiber/fiber/v2"
)

func GetShortUrl(c *fiber.Ctx) error {
	db := database.GetDbConnection()
	table_name := os.Getenv("TABLE_NAME")
	short_url := c.Params("shortURL")
	sql := fmt.Sprintf("SELECT (url) FROM %s WHERE short_url='%s'", table_name, short_url)
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatalln(err)
		return c.JSON("An error occured")
	}
	var url string
	for rows.Next() {
		rows.Scan(&url)
	}
	if url == "" {
		return c.JSON(struct {
			Error string `json:"error"`
		}{"URL Not Found :("})
	} else {
		return c.Redirect(url)
	}
}

func PostNewUrl(c *fiber.Ctx) error {
	db := database.GetDbConnection()
	table_name := os.Getenv("TABLE_NAME")
	newUrl := struct {
		Url      string
		ShortUrl string
	}{}
	if err := c.BodyParser(&newUrl); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	if newUrl.Url != "" {
		length, _ := strconv.ParseInt(os.Getenv("SHORT_URL_LENGTH"), 10, 64)
		var short_url string
		switch newUrl.ShortUrl {
		case "":
			short_url = helper.GenerateRandomString(int(length))
		default:
			short_url = newUrl.ShortUrl
			if database.IsExistShortURL(short_url) {
				return c.JSON(struct {
					Error string `json:"error"`
				}{"Short URL used, try another one"})
			}
		}
		sql := fmt.Sprintf("INSERT INTO %s (url, short_url) VALUES ('%s', '%s')", table_name, newUrl.Url, short_url)
		fmt.Println(sql)
		_, err := db.Exec(sql)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
		return c.JSON(struct {
			Url      string `json:"url"`
			ShortURL string `json:"short_url"`
		}{newUrl.Url, short_url})
	} else {
		return c.JSON(struct {
			Error string `json:"error"`
		}{"Please supply minimum a url, optionally a preferred short URL"})
	}
}

func HandleIndex(c *fiber.Ctx) error {
	now := time.Now()
	return c.Render("index", fiber.Map{
		"Timetz":   now.Format(time.RFC1123),
		"Port":     os.Getenv("PORT"),
		"Hostname": c.Hostname(),
	}, "layouts/main")
}
