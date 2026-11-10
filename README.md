# creatorcoaster.com
A multi-purpose website for Creator Coaster.

### Main features:
- **Wiki:** A central hub for all information surrounding Creator Coaster
- **Blog:** The outlet for announcements, explanations and our random ideas.

### Technologies:
- `Go`: The main language used on the backend for, well, everything
- `Templ`: A Go templating library/framework. This lets me easily build dynamic HTML pages with Go-like syntax
- `HTMX`: Some parts of this website use HTMX, this is mostly to avoid unnecessary JavaScript
- `HTML` + `CSS` + `JavaScript`: The only notable thing to mention here is that I only use JavaScript for tiny client-side interactive elements, not for communicating with the backend.

### Project Structure
```
src
├── config/                         # Configuration and static markdown files
│   └── .env                        # Configuration
│   └── static-markdown-content     # Some of the text on the website that isn't a wiki article or a blog
├── database/                       # Interaction with database
├── handlers/                       # Request Handlers
├── markdown/                       # Everything about generating HTML from markdown
├── myTypes/                        # Some types that are honestly only there because of circular import errors
├── static/                         # Static files served directly to the browser
│   ├── assets/                     # SVGs
│   ├── fonts/                      # Locally downloaded fonts
│   ├── script.js                   # All of the javascript
│   └── style.css                   # All styling
├── views/                          # Templ templates
└── main.go                         # Entrypoint & router config.
```

### Development:
- If you're interested in running the code locally and modifying it, please check out the [development](DEVELOPMENT.md) file.

### Contributing:
- Contributions aren't publicly open, but if you want to help me out, contact me on discord (`@vovoplay`)

### License clarification:
- This project is mainly open source so people could learn from it, but you can host it yourself provided you follow the [license](LICENSE) and the clarification below:
> If you host this website, you MUST remove all mentions of Creator Coaster and all Creator Coaster branding e.g. `--cc-brand-color` in CSS. You CANNOT impersonate the website and claim to represent Creator Coaster.