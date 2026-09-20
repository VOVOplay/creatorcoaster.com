# <img src="https://images.creatorcoaster.com/cc-animated-logo-circle.gif" height="28" alt="Logo">creatorcoaster.com
A multi-purpose website for Creator Coaster. This website extends the discord server instead of replacing it.

### Main features:
- **Wiki:** A central hub for all information surrounding Creator Coaster
- **Blogs:** The outlet for announcements, explanations and just our random ideas.
- **Privacy Policy:** A formal place for people to learn how we process their data on the discord server.

### Tech stack:
- **Go:** The backend and most of the logic
- **Templ:** The templating library for rendering HTML with Go syntax
- **HTMX:** A super simple frontend script that allows me to write almost no javascript
- **Goldmark:** A highly-extensible markdown renderer
- **Standard web technologies:** The only notable thing to say here is that thanks to HTMX & Templ I barely ever need to use Javascript
> You can learn more about the decision that led to these in the `About this website` section.

### Project Structure

```
src
├── config/        # Configuration and static markdown files
├── database/      # Interaction with database
├── handlers/      # Request Handlers
│   └── page.go    # Generic handler for non-special pages
├── markdown/      # My goldmark configuration
├── myTypes/       # Some types that are honestly only there because of circular import errors
├── static/        # Static files served directly to the browser
│   ├── assets/    # SVGs
│   ├── fonts/     # Locally downloaded fonts
│   ├── htmx/      # A copy of HTMX and its extensions
│   ├── script.js  # All of the javascript
│   └── style.css  # All styling
├── views/         # Templ templates
└── main.go        # Entry Point
```

### Development:
- If you're interested in running the code locally and modifying it, please check out the [development](DEVELOPMENT.md) file.

### Contributing:
- I am currently not really interested in getting code contributions. If you find a bug or want to request a feature, either contact me on discord (`@vovoplay`) or open an issue here on github.

### License clarification:
- This project is mainly open source so people could learn from it, but you can host it yourself provided you follow the [license](LICENSE) and the clarification below:
> If you host this website, you MUST remove all mentions of Creator Coaster and all Creator Coaster branding e.g. `--cc-brand-color` in CSS. You CANNOT impersonate the website and claim to represent Creator Coaster.