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
			DiscordUsername:    "@kdesa",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Owner",
		},
		{
			DiscordUsername:    "@vovoplay",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Head Manager",
		},
		{
			DiscordUsername:    "@qsynx",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Manager",
		},
		{
			DiscordUsername:    "@lgx_",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Head Admin",
		},
		{
			DiscordUsername:    "@fiona.ktk",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Admin",
		},
		{
			DiscordUsername:    "@furbated",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Admin",
		},
		{
			DiscordUsername:    "@sevtube",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Moderator",
		},
		{
			DiscordUsername:    "@hazelstyx",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Moderator",
		},
		{
			DiscordUsername:    "@reaper_mc.",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Moderator",
		},
		{
			DiscordUsername:    "@random_dudy",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Moderator",
		},
		{
			DiscordUsername:    "@fyzter123",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Moderator",
		},
		{
			DiscordUsername:    "@eulmdev",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Support",
		},
		{
			DiscordUsername:    "@claucefx",
			ProfilePictureLink: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROpm841O9MktJCr_u6slU_C-XDufbsj0GzH2VTWsjs3A&s",
			Role:               "Support",
		},
	}
}

func (h *PageHandler) HandleAbout(w http.ResponseWriter, r *http.Request) {
	teamMemberInfo := getTeamMemberInfo()

	component := views.AboutUs(teamMemberInfo)
	component.Render(r.Context(), w)
}
