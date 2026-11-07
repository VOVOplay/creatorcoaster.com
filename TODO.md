Legend:
- ✅: Done
- ➖: Working on
- ❌: Not done

## Actual things that still need to be done:
- Markdown:
    - ❌ Variables from wiki article into db variable: 
    {new_var_name}(
        Text
        Text
    )
    - ➖ Rendering markdown meant for discord (e. g. Replacing #CHANNEL with <#CHANEL_ID>...) - seperate from goldmark completely

- Admin panel:
    - ❌ Uploading all types of variables
    - ❌ Managing all types of variables
    - ➖ Add pre-defined categories (instead of each article having its own path, they'll have to choose from one of the already existing categories)
        - ❌ Add it in the admin panel
        - ✅ Add it in the database
        - ✅ Update code currently using category paths
        - ➖ Update the wiki article upload/edit pages

- Other:
    - ➖ Designing and coding the landing page
    - ❌ Making the wiki page usable on mobile
    - ❌ Figure out sorting for categories
    - ❌ Applying a UGC policy to articles

- Before release:
    - ❌ Setting up the real prod database and making it accessible from pterodactyl
    - ❌ Somehow make the schema not be a part of the bot nor the website, so theres one source of truth for that
    - ❌ Fix up the bot and release the version that saves the staff list
    - ❌ Figure out the compilation properly