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
- ➖ Tuning markdown rendering and CSS

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
- ❌ Transcript Viewer

## Actual things that still need to be done:
- Markdown syntax for cross-bot/website variables.
    - From db variable to markdown sources: {variable_name}
    - From wiki article into db variable: 
    {new_var_name}(
        Text
        Text
    )

- Admin panel:
    - ✅ Discord login
    - Uploading blogs & Wiki articles
    - Managing wiki articles (changing their category, name...)
    - Managing blog articles (changing their name, author...)
    - Maybe something for managing the staff list in About Us

- Fixing current markdown issues:
    - Spacer doesnt work sometimes
    - Some of the margins arent what Id like them to be

- Other:
    - Designing and coding the landing page
    - Replacing 404 page logo with a new one
    - Normalizing some elements around the website to be consistent (e. g. vertial-spacer vs a pipe)
    - Making the wiki page usable on mobile
    - Adding images.creatorcoaster.com as a part of the website

- Before release:
    - Setting up the real prod database and making it accessible from pterodactyl
    - Somehow make the schema not be a part of the bot nor the website, so theres one source of truth for that
    - Fix up the bot and release the version that saves the staff list
    - Figure out the compilation properly

