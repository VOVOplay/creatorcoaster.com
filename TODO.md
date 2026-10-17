Legend:
- ✅: Done
- ➖: Working on
- ❌: Not done

**Wiki:**
- ✅ Wiki page
- ✅ Articles
- ✅ Categories
- ✅ Sidebar
- ✅ Loading from database
- ❌ Upload wiki articles from website
- ❌ Markdown syntax for variables (from database: {HOME_CHANNEL}, and to database: ?)
- ✅ Tuning markdown rendering and CSS

**Blogs:**
- ✅ Blog page
- ✅ Markdown Rendering
- ✅ Article setup

**Other pages:**
- ✅ About Us page
- ✅ About this website page
- ✅ Privacy Policy Page
- ➖ Home Page

**Misc:**
- ✅ Discord Login

## Actual things that still need to be done:
- Markdown:
    - ✅ Variables from db variable to markdown sources: {variable_name}
    - ❌ Variables from wiki article into db variable: 
    {new_var_name}(
        Text
        Text
    )
    - ❌ @ROLE Variables
    - ✅ #CHANNEL Variables
    - Rendering markdown meant for discord (e. g. Replacing #CHANNEL with <#CHANEL_ID>...)

- Admin panel:
    - Uploading blogs & Wiki articles
    - Managing wiki articles (changing their category, name...)
    - Managing blog articles (changing their name, author...)
    - Maybe something for managing the staff list in About Us

- Other:
    - Designing and coding the landing page
    - Making the wiki page usable on mobile
    - Fix issue where the sidebar state resets when clicking to a different wiki article
    - Figure out sorting for categories

- Before release:
    - Setting up the real prod database and making it accessible from pterodactyl
    - Somehow make the schema not be a part of the bot nor the website, so theres one source of truth for that
    - Fix up the bot and release the version that saves the staff list
    - Figure out the compilation properly
