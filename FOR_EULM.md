# Wiki Markdown -> Discord Markdown

I added some custom features to my version of markdown, specifically custom variables that get pulled from the database and are defined ahead of time. These use custom syntax, but since the wiki articles have to be compatible with discord markdown, I need to be able to transpile them to discord markdown.

I probably won't be translating entire articles into discord all the time, so I want a simple function like `func GenerateDiscordMarkdownFromString(source string) string {}` in the `markdown` package.

## Standard Variables:
These are nothing special, one thing to mention is that they don't support multi-line strings.
Syntax: `{variable_name}`
How to get them: `database.GetVariableByName("variable_name")` - Returns a string and an error.


## Channel Variables:
When these get rendered, they look like a discord channel and are clickable.
When you are rendering them, the format needs to be the discord compatible `<#ID>`
Syntax: `{#channel_variable_name}`
Type: `myTypes.ChannelVariable`
How to get them: `database.GetChannelVariableByName("channel_variable_name")`

## Role/User Variables
These represent both since there's little functional difference between a user and a role.
When rendering them, you'll need to differentiate between users (`<@ID>`) and roles (`<@&ID>`)
Syntax: `{@variable_name}`
Type: `myTypes.RoleUserVariable`
How to get them: `database.GetRoleUserVariableByName("variable_name")`

Note: The `GetRoleUserVariableByName` function doesnt properly return `RoleUserVariable` yet, I'll get around to that later.



**Notes:**
If a variable is not found, keep the text as if that custom syntax never existed.
The definitions for types are in myTypes/markdownVariables.go
