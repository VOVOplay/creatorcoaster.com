package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VOVOplay/creatorcoaster.com/src/config"
	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
	_ "github.com/go-sql-driver/mysql"
)

var all_config config.Config = config.GetConfig()
var db_config config.DatabaseConfig = all_config.DatabaseConfig

var dsn string = fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s",
	db_config.User,
	db_config.Password,
	db_config.EntryPort,
	db_config.Name,
)

func GetStaffMemberList() ([]myTypes.StaffMember, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return []myTypes.StaffMember{}, myTypes.ErrDatabseOffline
	}

	var staffList []myTypes.StaffMember

	rawStaffList, err := db.Query("SELECT * FROM staff_list")
	if err != nil {
		return []myTypes.StaffMember{}, myTypes.ErrUnknownDatabaseError
	}
	defer rawStaffList.Close()

	for rawStaffList.Next() {
		var staffMember myTypes.StaffMember
		err := rawStaffList.Scan(&staffMember.UserID, &staffMember.Username, &staffMember.ProfilePictureLink, &staffMember.Position, &staffMember.PositionPrettyName)
		if err != nil {
			log.Fatal(err)
		}
		staffList = append(staffList, staffMember)
	}

	if rawStaffList.Err() != nil {
		return []myTypes.StaffMember{}, myTypes.ErrUnknownDatabaseError
	}

	return staffList, nil
}

func GetAllBlogs() ([]myTypes.Article, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return []myTypes.Article{}, myTypes.ErrDatabseOffline
	}

	var allBlogs []myTypes.Article

	rawBlogs, err := db.Query("SELECT name, pretty_name, generated_html, date, author, author_pfp_link, read_time FROM blog_posts")
	if err != nil {
		return []myTypes.Article{}, myTypes.ErrUnknownDatabaseError
	}
	defer rawBlogs.Close()

	for rawBlogs.Next() {
		var blog myTypes.Article
		err := rawBlogs.Scan(&blog.Name, &blog.PrettyName, &blog.HTML, &blog.Info.Date, &blog.Info.Author, &blog.Info.AuthorPFPLink, &blog.Info.ReadTime)
		if err != nil {
			log.Fatal(err)
		}

		blog.Type = myTypes.BlogArticle
		allBlogs = append(allBlogs, blog)
	}

	if rawBlogs.Err() != nil {
		return []myTypes.Article{}, myTypes.ErrUnknownDatabaseError
	}

	return allBlogs, nil
}

func GetBlogByName(name string) (myTypes.Article, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return myTypes.Article{}, myTypes.ErrDatabseOffline
	}

	query := "SELECT name, pretty_name, generated_html, date, author, author_pfp_link, read_time FROM blog_posts WHERE name = ?"

	var blog myTypes.Article

	err = db.QueryRow(query, name).Scan(&blog.Name, &blog.PrettyName, &blog.HTML, &blog.Info.Date, &blog.Info.Author, &blog.Info.AuthorPFPLink, &blog.Info.ReadTime)
	if err != nil {
		return myTypes.Article{}, myTypes.ErrArticleNotFound
	}

	blog.Type = myTypes.BlogArticle

	return blog, nil
}

func GetAllWikiArticles() ([]myTypes.Article, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return []myTypes.Article{}, myTypes.ErrDatabseOffline
	}

	var allWikiArticles []myTypes.Article

	rawWikiArticles, err := db.Query("SELECT name, pretty_name, generated_html, category_path, last_edited_at FROM wiki_articles")
	if err != nil {
		return []myTypes.Article{}, myTypes.ErrUnknownDatabaseError
	}
	defer rawWikiArticles.Close()

	for rawWikiArticles.Next() {
		var wikiArticle myTypes.Article
		err := rawWikiArticles.Scan(&wikiArticle.Name, &wikiArticle.PrettyName, &wikiArticle.HTML, &wikiArticle.CategoryPath, &wikiArticle.Info.Date)
		if err != nil {
			return []myTypes.Article{}, myTypes.ErrUnknownDatabaseError
		}

		wikiArticle.Type = myTypes.WikiArticle

		if wikiArticle.Name == myTypes.FirstArticleName {
			wikiArticle.Info.IsFirstWikiArticle = true
		}

		allWikiArticles = append(allWikiArticles, wikiArticle)
	}

	if rawWikiArticles.Err() != nil {
		return []myTypes.Article{}, myTypes.ErrUnknownDatabaseError
	}

	return allWikiArticles, nil
}

func GetWikiArticleByName(name string) (myTypes.Article, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return myTypes.Article{}, myTypes.ErrDatabseOffline
	}

	query := "SELECT name, pretty_name, generated_html, category_path, last_edited_at FROM wiki_articles WHERE name = ?"

	var wikiArticle myTypes.Article

	err = db.QueryRow(query, name).Scan(&wikiArticle.Name, &wikiArticle.PrettyName, &wikiArticle.HTML, &wikiArticle.CategoryPath, &wikiArticle.Info.Date)
	if err != nil {
		return myTypes.Article{}, myTypes.ErrArticleNotFound
	}

	wikiArticle.Type = myTypes.WikiArticle
	return wikiArticle, nil
}
