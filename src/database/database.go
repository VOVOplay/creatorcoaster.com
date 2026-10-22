package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/VOVOplay/creatorcoaster.com/src/config"
	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init(databaseConfig config.DatabaseConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Address,
		databaseConfig.Name,
	)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return nil
}

func Close() {
	if db != nil {
		db.Close()
	}
}

func GetStaffMemberList() ([]myTypes.StaffMember, error) {
	if db == nil {
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
			return []myTypes.StaffMember{}, myTypes.ErrUnknownDatabaseError
		}
		staffList = append(staffList, staffMember)
	}

	if rawStaffList.Err() != nil {
		return []myTypes.StaffMember{}, myTypes.ErrUnknownDatabaseError
	}

	return staffList, nil
}

// Doesn't return a full blog, the HTML is missing
func GetAllBlogs() ([]myTypes.Article, error) {
	if db == nil {
		return []myTypes.Article{}, myTypes.ErrDatabseOffline
	}

	var allBlogs []myTypes.Article

	rawBlogs, err := db.Query("SELECT name, pretty_name, date, author, author_pfp_link, read_time FROM blog_posts")
	if err != nil {
		return []myTypes.Article{}, myTypes.ErrUnknownDatabaseError
	}
	defer rawBlogs.Close()

	for rawBlogs.Next() {
		var blog myTypes.Article
		err := rawBlogs.Scan(&blog.Name, &blog.PrettyName, &blog.Info.Date, &blog.Info.Author, &blog.Info.AuthorPFPLink, &blog.Info.ReadTime)
		if err != nil {
			return []myTypes.Article{}, myTypes.ErrUnknownDatabaseError
		}

		blog.Type = myTypes.BlogArticle
		allBlogs = append(allBlogs, blog)
	}

	if rawBlogs.Err() != nil {
		return []myTypes.Article{}, myTypes.ErrUnknownDatabaseError
	}

	return allBlogs, nil
}

func GetBlogByName(name string, renderer func(string) string) (myTypes.Article, error) {
	if db == nil {
		return myTypes.Article{}, myTypes.ErrDatabseOffline
	}

	query := "SELECT name, pretty_name, md_source, date, author, author_pfp_link, read_time FROM blog_posts WHERE name = ?"

	var blog myTypes.Article
	var markdownSource string

	err := db.QueryRow(query, name).Scan(&blog.Name, &blog.PrettyName, &markdownSource, &blog.Info.Date, &blog.Info.Author, &blog.Info.AuthorPFPLink, &blog.Info.ReadTime)
	if err != nil {
		return myTypes.Article{}, myTypes.ErrArticleNotFound
	}

	HTML := renderer(markdownSource)
	blog.HTML = HTML

	blog.Type = myTypes.BlogArticle

	return blog, nil
}

// Doesn't return full wiki articles, the HTML is missing.
func GetAllWikiArticles() ([]myTypes.Article, error) {
	if db == nil {
		return []myTypes.Article{}, myTypes.ErrDatabseOffline
	}

	var allWikiArticles []myTypes.Article

	rawWikiArticles, err := db.Query("SELECT name, pretty_name, category_path, last_edited_at, sort_order FROM wiki_articles")
	if err != nil {
		return []myTypes.Article{}, myTypes.ErrUnknownDatabaseError
	}
	defer rawWikiArticles.Close()

	for rawWikiArticles.Next() {
		var wikiArticle myTypes.Article
		err := rawWikiArticles.Scan(&wikiArticle.Name, &wikiArticle.PrettyName, &wikiArticle.CategoryPath, &wikiArticle.Info.Date, &wikiArticle.SortOrder)
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

func GetWikiArticleByName(name string, renderer func(string) string) (myTypes.Article, error) {
	if db == nil {
		return myTypes.Article{}, myTypes.ErrDatabseOffline
	}

	query := "SELECT name, pretty_name, md_source, category_path, last_edited_at FROM wiki_articles WHERE name = ?"

	var wikiArticle myTypes.Article
	var markdownSource string

	err := db.QueryRow(query, name).Scan(&wikiArticle.Name, &wikiArticle.PrettyName, &markdownSource, &wikiArticle.CategoryPath, &wikiArticle.Info.Date)
	if err != nil {
		return myTypes.Article{}, myTypes.ErrArticleNotFound
	}

	HTML := renderer(markdownSource)
	wikiArticle.HTML = HTML

	wikiArticle.Type = myTypes.WikiArticle
	return wikiArticle, nil
}

func SaveWikiArticle(wikiArticle myTypes.Article) error {
	if db == nil {
		return myTypes.ErrDatabseOffline
	}

	query := "INSERT INTO wiki_articles (name, pretty_name, md_source, category_path, last_edited_at, sort_order) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, wikiArticle.Name, wikiArticle.PrettyName, wikiArticle.MarkdownSource, wikiArticle.CategoryPath, wikiArticle.Info.Date, wikiArticle.SortOrder)
	if err != nil {
		return err
	}

	return nil
}

func SaveBlogPost(blog myTypes.Article) error {
	if db == nil {
		return myTypes.ErrDatabseOffline
	}

	query := "INSERT INTO blog_posts (name, pretty_name, md_source, date, author, author_pfp_link, read_time) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, blog.Name, blog.PrettyName, blog.MarkdownSource, blog.Info.Date, blog.Info.Author, blog.Info.AuthorPFPLink, blog.Info.ReadTime)
	if err != nil {
		return err
	}

	return nil
}

// --- Variables

func GetVariableByName(name string) (string, error) {
	if db == nil {
		return "", myTypes.ErrDatabseOffline
	}

	query := "SELECT variable_contents FROM markdown_variables WHERE variable_name = ?"

	var variableValue string
	err := db.QueryRow(query, name).Scan(&variableValue)
	if err != nil {
		return "", myTypes.ErrVariableNotFound
	}

	return variableValue, nil
}

func GetChannelVariableByName(name string) (myTypes.ChannelVariable, error) {
	if db == nil {
		return myTypes.ChannelVariable{}, myTypes.ErrDatabseOffline
	}

	query := "SELECT channel_name, channel_id FROM markdown_channel_variables WHERE variable_name = ?"

	var channel myTypes.ChannelVariable
	err := db.QueryRow(query, name).Scan(&channel.Name, &channel.ID)
	if err != nil {
		return myTypes.ChannelVariable{}, myTypes.ErrVariableNotFound
	}

	return channel, nil
}

func GetRoleUserVariableByName(name string) (string, error) {
	if db == nil {
		return "", myTypes.ErrDatabseOffline
	}

	query := "SELECT pretty_name FROM markdown_role_user_variables WHERE variable_name = ?"

	var variableValue string
	err := db.QueryRow(query, name).Scan(&variableValue)
	if err != nil {
		return "", myTypes.ErrVariableNotFound
	}

	return variableValue, nil
}
