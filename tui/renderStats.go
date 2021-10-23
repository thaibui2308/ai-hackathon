package tui

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/thaibui2308/ai-hackathon/api"
	"github.com/thaibui2308/ai-hackathon/models"
)

func RenderUserInfo(userUrl string, stats models.Stats, issue string, files []models.File) {
	user, err := api.GetUserInfo(userUrl)

	if err != nil {
		log.Fatalf("Couldn't get user details: %v", err)
	}
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	img, images := SetupImage(user.AvatarURL, user.Login)
	p := SetupProfileInfo(user)
	p1 := SetupStatsCheckInfo(stats, issue)
	pc := SetupAdditionChart(files)
	pc1 := SetupDeletionChart(files)
	render := func() {
		img.Image = images[0]

		img.Title = fmt.Sprintf(user.Login + "'s github")

		ui.Render(img, p, p1, pc, pc1)
	}
	render()
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Up>", "k":
			img.MonochromeThreshold++
		case "<Down>", "j":
			img.MonochromeThreshold--
		case "<Enter>":
			img.Monochrome = !img.Monochrome
		case "<Tab>":
			img.MonochromeInvert = !img.MonochromeInvert
		}
		render()
	}
}
