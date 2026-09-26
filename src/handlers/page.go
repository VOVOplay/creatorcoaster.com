package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/VOVOplay/creatorcoaster.com/src/database"
	"github.com/VOVOplay/creatorcoaster.com/src/myTypes"
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

	// for caching
	w.Header().Set("Cache-Control", "public, max-age=60")

	component := views.Home()
	component.Render(r.Context(), w)
}

func (h *PageHandler) HandleAbout(w http.ResponseWriter, r *http.Request) {
	staffList := database.GetStaffMemberList()
	sortedStaffList := sortStaffList(staffList)

	// for caching
	w.Header().Set("Cache-Control", "public, max-age=60")

	component := views.AboutUs(sortedStaffList)
	component.Render(r.Context(), w)
}

func (h *PageHandler) HandleAboutWebsite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=60")

	component := views.AboutWebsite()
	component.Render(r.Context(), w)
}

func (h *PageHandler) HandlePrivacy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=60")

	component := views.Privacy()
	component.Render(r.Context(), w)
}

func loadTestFile() string {
	filePath := filepath.Join("./config/static-markdown-content/EXAMPLE-WIKI.md")
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	data_str := string(data)

	return data_str
}

func sortStaffList(staffList []myTypes.StaffMember) []myTypes.StaffMember {
	correctOrder := []string{
		"owner",
		"head_manager",
		"manager",
		"head_admin",
		"admin",
		"head_moderator",
		"moderator",
		"support",
	}

	var sortedStaffList []myTypes.StaffMember

	for _, position := range correctOrder {
		for _, staffMember := range staffList {
			if staffMember.Position == position {
				sortedStaffList = append(sortedStaffList, staffMember)
			}
		}
	}

	return sortedStaffList
}
