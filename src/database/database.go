package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/VOVOplay/creatorcoaster.com/src/config"
	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initializes the database and opens a connection for the lifetime of the program.
// database.Close() must be called when the process ends.
func Init(config config.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.DatabaseConfig.User,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Address,
		config.DatabaseConfig.Name,
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

	if !config.IsProduction {
		createDatabase()
	}

	return nil
}

// Closes the database connection
func Close() {
	if db != nil {
		db.Close()
	}
}

func createDatabase() {
	if db == nil {
		return
	}

	schemaPath := filepath.Join("./database/schema.sql")
	data, err := os.ReadFile(schemaPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	dataString := string(data)

	queries := strings.Split(dataString, ";")

	for _, query := range queries {
		db.Exec(query)
	}

}

// Gets the staff member list from the table 'staff_list'
// If the staff_list is empty, it returns []myTypes.StaffMember{}, this behavior has to be handled in the callsite.
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

// Gets all of the blogs.
// The HTML field is missing in all of the articles.
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

// Gets a blog by its name.
// renderer is the method that will render the HTML, aka. markdown.GenerateHTMLFromString()
func GetBlogByName(name string, renderer func(string) string) (myTypes.Article, error) {
	if db == nil {
		return myTypes.Article{}, myTypes.ErrDatabseOffline
	}

	query := "SELECT name, pretty_name, md_source, date, author, author_pfp_link, read_time FROM blog_posts WHERE name = ?"

	var blog myTypes.Article

	err := db.QueryRow(query, name).Scan(&blog.Name, &blog.PrettyName, &blog.MarkdownSource, &blog.Info.Date, &blog.Info.Author, &blog.Info.AuthorPFPLink, &blog.Info.ReadTime)
	if err != nil {
		return myTypes.Article{}, myTypes.ErrArticleNotFound
	}

	HTML := renderer(blog.MarkdownSource)
	blog.HTML = HTML

	blog.Type = myTypes.BlogArticle

	return blog, nil
}

// Gets all of the wiki articles.
// The HTML field is missing in all of the articles.
func GetAllWikiArticles() ([]myTypes.Article, error) {
	if db == nil {
		return []myTypes.Article{}, myTypes.ErrDatabseOffline
	}

	var allWikiArticles []myTypes.Article

	rawWikiArticles, err := db.Query("SELECT name, pretty_name, parent_category_id, last_edited_at, sort_order FROM wiki_articles")
	if err != nil {
		return []myTypes.Article{}, myTypes.ErrUnknownDatabaseError
	}
	defer rawWikiArticles.Close()

	for rawWikiArticles.Next() {
		var wikiArticle myTypes.Article
		var parentCategoryId *int
		err := rawWikiArticles.Scan(&wikiArticle.Name, &wikiArticle.PrettyName, &parentCategoryId, &wikiArticle.Info.Date, &wikiArticle.SortOrder)
		if err != nil {
			return []myTypes.Article{}, err
		}

		if parentCategoryId == nil {
			wikiArticle.ParentCategoryId = myTypes.RootCategoryId
		} else {
			wikiArticle.ParentCategoryId = *parentCategoryId
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

// Gets a wiki article by its name.
// renderer is the method that will render the HTML, aka. markdown.GenerateHTMLFromString()
func GetWikiArticleByName(name string, renderer func(string) string) (myTypes.Article, error) {
	if db == nil {
		return myTypes.Article{}, myTypes.ErrDatabseOffline
	}

	query := "SELECT name, pretty_name, md_source, parent_category_id, last_edited_at FROM wiki_articles WHERE name = ?"

	var wikiArticle myTypes.Article
	var parentCategoryId *int

	err := db.QueryRow(query, name).Scan(&wikiArticle.Name, &wikiArticle.PrettyName, &wikiArticle.MarkdownSource, &parentCategoryId, &wikiArticle.Info.Date)
	if err != nil {
		return myTypes.Article{}, err
	}

	if parentCategoryId == nil {
		wikiArticle.ParentCategoryId = myTypes.RootCategoryId
	} else {
		wikiArticle.ParentCategoryId = *parentCategoryId
	}

	HTML := renderer(wikiArticle.MarkdownSource)
	wikiArticle.HTML = HTML

	wikiArticle.Type = myTypes.WikiArticle
	return wikiArticle, nil
}

func GetRawWikiCategories() ([]myTypes.RawCategory, error) {
	if db == nil {
		return []myTypes.RawCategory{}, myTypes.ErrDatabseOffline
	}

	var allCategories []myTypes.RawCategory

	query := "SELECT id, pretty_name, sort_order, parent_category_id FROM wiki_categories"
	rawCategories, err := db.Query(query)
	if err != nil {
		return []myTypes.RawCategory{}, err
	}
	defer rawCategories.Close()

	for rawCategories.Next() {
		var rawCategory myTypes.RawCategory
		var parentCategoryId *int
		err := rawCategories.Scan(&rawCategory.Id, &rawCategory.Name, &rawCategory.SortOrder, &parentCategoryId)
		if err != nil {
			return []myTypes.RawCategory{}, err
		}

		if parentCategoryId == nil {
			rawCategory.ParentCategoryId = myTypes.RootCategoryId
		} else {
			rawCategory.ParentCategoryId = *parentCategoryId
		}

		allCategories = append(allCategories, rawCategory)
	}

	if rawCategories.Err() != nil {
		return []myTypes.RawCategory{}, err
	}

	return allCategories, nil
}

// Saves a wiki article to the database.
func SaveWikiArticle(wikiArticle myTypes.Article) error {
	if db == nil {
		return myTypes.ErrDatabseOffline
	}

	// this converts the root category id into `nil` so that it properly works in the database as having no parent category
	var parentCategoryId *int
	if wikiArticle.ParentCategoryId != myTypes.RootCategoryId {
		parentCategoryId = &wikiArticle.ParentCategoryId
	}

	query := "INSERT INTO wiki_articles (name, pretty_name, md_source, parent_category_id, last_edited_at, sort_order) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, wikiArticle.Name, wikiArticle.PrettyName, wikiArticle.MarkdownSource, parentCategoryId, wikiArticle.Info.Date, wikiArticle.SortOrder)
	if err != nil {
		return err
	}

	return nil
}

// Updates a wiki article in the database.
// The previousName value is the old name. Even if the name hasn't changed, this value has to be provided.
func UpdateWikiArticle(newWikiArticle myTypes.Article, previousName string) error {
	if db == nil {
		return myTypes.ErrDatabseOffline
	}

	// this converts the root category id into `nil` so that it properly works in the database as having no parent category
	var parentCategoryId *int
	if newWikiArticle.ParentCategoryId != myTypes.RootCategoryId {
		parentCategoryId = &newWikiArticle.ParentCategoryId
	}

	query := "UPDATE wiki_articles SET name = ?, pretty_name = ?, md_source = ?, parent_category_id = ?, last_edited_at = ?, sort_order = ? WHERE name = ?"
	_, err := db.Exec(query, newWikiArticle.Name, newWikiArticle.PrettyName, newWikiArticle.MarkdownSource, parentCategoryId, newWikiArticle.Info.Date, newWikiArticle.SortOrder, previousName)
	if err != nil {
		return err
	}

	return nil
}

// Deletes a wiki article from the database by name.
func DeleteWikiArticle(name string) error {
	if db == nil {
		return myTypes.ErrDatabseOffline
	}

	query := "DELETE FROM wiki_articles WHERE name = ?"
	_, err := db.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}

// Saves a blog post to the database.
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

// Updates a blog post in the database.
// The previousName value is the old name. Even if the name hasn't changed, this value has to be provided.
func UpdateBlogPost(newBlogPost myTypes.Article, previousName string) error {
	if db == nil {
		return myTypes.ErrDatabseOffline
	}

	query := "UPDATE blog_posts SET name = ?, pretty_name = ?, md_source = ?, author = ?, author_pfp_link = ?, read_time = ? WHERE name = ?"
	_, err := db.Exec(query, newBlogPost.Name, newBlogPost.PrettyName, newBlogPost.MarkdownSource, newBlogPost.Info.Author, newBlogPost.Info.AuthorPFPLink, newBlogPost.Info.ReadTime, previousName)
	if err != nil {
		return err
	}

	return nil
}

// Deletes a blog post in the database.
func DeleteBlogArticle(name string) error {
	if db == nil {
		return myTypes.ErrDatabseOffline
	}

	query := "DELETE FROM blog_posts WHERE name = ?"
	_, err := db.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}

// --- Variables

// Returns a string of a standard variable's contents.
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

// Returns a channel variable.
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

// Returns a role-user variable.
func GetRoleUserVariableByName(name string) (myTypes.RoleUserVariable, error) {
	if db == nil {
		return myTypes.RoleUserVariable{}, myTypes.ErrDatabseOffline
	}

	query := "SELECT role_user_name, role_user_id, variable_type FROM markdown_role_user_variables WHERE variable_name = ?"

	var roleUserVariable myTypes.RoleUserVariable
	var variableType int
	err := db.QueryRow(query, name).Scan(&roleUserVariable.Name, &roleUserVariable.ID, &variableType)
	if err != nil {
		return myTypes.RoleUserVariable{}, myTypes.ErrVariableNotFound
	}

	switch variableType {
	case int(myTypes.Role):
		roleUserVariable.Type = myTypes.Role
	case int(myTypes.User):
		roleUserVariable.Type = myTypes.User
	default:
		fmt.Println("Warning: Got value other than 0/1 from database in table 'markdown_role_user_variables' for field `variable_type`")
		roleUserVariable.Type = myTypes.User
	}

	return roleUserVariable, nil
}
