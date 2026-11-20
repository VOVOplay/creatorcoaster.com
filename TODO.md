Legend:
- ✅: Done
- ➖: Working on
- ❌: Not done

## Actual things that still need to be done:
- Markdown:
    - ➖ Variables from wiki article into db variable:
        - ✅ Preprocessing function: Removes the syntax of this block from the input source (The contents stay there), this gets called in the generate html function and doesn't save anything into the database
        - ➖ Saving function: Saves all the the article sections into the database (maybe returns them as well depending on the requirements), this only gets called when actually submitting an article.
    - ➖ Rendering markdown meant for discord (e. g. Replacing #CHANNEL with <#CHANEL_ID>...) - seperate from goldmark completely

- Admin panel:
    - ❌ Managing/Uploading Categories
    - ❌ Adding "hints" to upload/update pages for blog and wiki such as existing variables or existing sort-orders

- Other:
    - ➖ Designing and coding the landing page
    - ❌ Mobile Improvements (Wiki & Blog pages, footer...)
    - ❌ Go over the security of the whole admin thing to make sure it's alright
    - ❌ Write the "About Us" page

- Before release:
    - ❌ Set up a proper database
    - ❌ Fix up the bot and release the version that saves the staff list

