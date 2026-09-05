package handlers

import (
	"net/http"

	"github.com/VOVOplay/creatorcoaster.com/src/views"
)

type PageHandler struct {
}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (h *PageHandler) HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	component := views.Home()
	component.Render(r.Context(), w)
}

// temporary for testing
func getTeamMemberInfo() []views.TeamMemberInfo {
	return []views.TeamMemberInfo{
		{
			DiscordUsername:    "@vovoplay",
			ProfilePictureLink: "/static/assets/discord-logo.svg",
			Role:               "Head Manager",
		},
		{
			DiscordUsername:    "@qsynx",
			ProfilePictureLink: "/static/assets/discord-logo.svg",
			Role:               "Manager",
		},
		{
			DiscordUsername:    "@vovoplay",
			ProfilePictureLink: "/static/assets/discord-logo.svg",
			Role:               "Head Manager",
		},
		{
			DiscordUsername:    "@qsynx",
			ProfilePictureLink: "/static/assets/discord-logo.svg",
			Role:               "Manager",
		},
		{
			DiscordUsername:    "@vovoplay",
			ProfilePictureLink: "/static/assets/discord-logo.svg",
			Role:               "Head Manager",
		},
		{
			DiscordUsername:    "@vovoplay",
			ProfilePictureLink: "/static/assets/discord-logo.svg",
			Role:               "Head Manager",
		},
		{
			DiscordUsername:    "@vovoplay",
			ProfilePictureLink: "/static/assets/discord-logo.svg",
			Role:               "Head Manager",
		},
		{
			DiscordUsername:    "@vovoplay",
			ProfilePictureLink: "/static/assets/discord-logo.svg",
			Role:               "Head Manager",
		},
		{
			DiscordUsername:    "@vovoplay",
			ProfilePictureLink: "/static/assets/discord-logo.svg",
			Role:               "Head Manager",
		},
	}
}

func (h *PageHandler) HandleAbout(w http.ResponseWriter, r *http.Request) {
	teamMemberInfo := getTeamMemberInfo()

	component := views.AboutUs(teamMemberInfo)
	component.Render(r.Context(), w)
}
