# About this website
Hey, welcome to this little section where I (`@vovoplay`) talk a bit about this website!

## Why?
Why even make a website? Most discord servers definitely don't have one.

At some point, me and a few other staff members realized how horribly decentralized all of the information about Creator Coaster was. We kind of solved this for the staff team with a knowledgebase forum channel in discord, but for everyone else, I wanted a better solution.

The idea was for a website where all of the information could be stored, and have a system integrated with the bot for sharing that information. For example, parts of the wiki articles could be marked as a "variable" the bot would have access to. This would allow me to have messages in discord, that would contain some of the information (e. g. shorter version of an article) directly from the wiki. And vice-versa, I could use variables directly from discord for the wiki articles. A good example of this would be a way to link to discord channels directly from the wiki.

## Ideology
From the start, I wanted to make a **simple** website, more reminiscent of the older web. That's why it might look a little boring to some people, but I honestly couldn't love it more. A surprisingly big inspiration ended up being the [discord.py documentation](https://discordpy.readthedocs.io/) website, which (in my opinion) really strikes the balance between being usable and nice-looking. Another inspiration was [localhost8000.com](https://localhost8000.com/). I like the way that looks, and I even stole their fonts.

## Technologies
- `Go`: The main language used on the backend for, well, everything
- `Templ`: A Go templating library/framework. This lets me easily build dynamic HTML pages with Go-like syntax
- `HTMX`: Some parts of this website use HTMX, this is mostly to avoid unnecessary JavaScript
- `HTML` + `CSS` + `JavaScript`: The only notable thing to mention here is that I only use JavaScript for tiny client-side interactive elements, not for communicating with the backend.

## How I came to choose those technologies
Coming into this, I knew I wanted to avoid JavaScript wherever possible. I am truly a JS hater. Thankfully for me, there are many JavaScript-less ways to build a website. I went with Go for the backend because I wanted to learn a new language and because `@eulmdev` has a lot of experience in Go, so he convinced me it's the best option.

I chose Templ because it just seemed like a much nicer way to make templates in Go than the standard template package. I really love how Templ (almost) fully supports normal Go syntax, for example making multiple divs can be as simple as:
```
for i := range 5 {
    <div>{ i }</div>
}
```

As for HTMX, that was an obvious choice for letting me avoid more JavaScript. So far the user-facing part of the website doesn't use HTMX for much more than preloading pages, but even just that was really easy to implement with HTMX.

## Open-Source
Why is this website even open-source? I mean the bot isn't right?

Well, originally I wanted to write my own markdown-inspired format for the wiki/blog articles, but in the end I gave up on using that on this website. So the original idea was that people could learn from that parser. But even though the parser isn't used here, I still decided to keep the website open-source. I'm not sure if anyone will get much value from the code, but you never know. If you want to do something with the code, I fully encourage you to do so as long as it follows the [AGPL-3.0 license](https://github.com/VOVOplay/creatorcoaster.com/blob/main/LICENSE).

If you're wondering why the bot isn't open source, it's because of a few things. First off the repo itself has some old credentials that shouldn't get leaked. A lot of script-kiddies also love to completely copy Creator Coaster, and I don't want them to straight up steal hundreds of hours of work I put into the bot. 

## Thank-yous
Thank you to the whole CC Staff Team for motivating me to keep going with the website, `@eulmdev` for helping me with Go, and `@phi0ta` for motivating me to do the hard stuff I wanted to avoid.