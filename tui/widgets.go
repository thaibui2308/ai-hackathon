package tui

import (
	"fmt"
	"math"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thaibui2308/ai-hackathon/models"
)

func SetupProfileInfo(user models.User) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Text = BuildUserInfo(user)
	p.SetRect(0, 21, 30, 14)

	return p
}

func SetupStatsCheckInfo(stats models.Stats, issue string) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Title = "Pull Request's Summary"
	p.SetRect(35, 0, 70, 12)
	p.Text = BuildStatChecked(stats, issue)

	return p
}

func SetupAdditionChart(files []models.File) *widgets.PieChart {
	additions := make([]float64, 0)

	for _, file := range files {
		additions = append(additions, float64(file.Additions))

	}

	pc := widgets.NewPieChart()
	pc.Title = "Additions"
	pc.SetRect(35, 28, 70, 13)

	addNum := len(additions)
	if addNum == 0 {
		pc.Title = "Additions changed by commit (no files changed)"
	} else {
		pc.Data = additions[:addNum]
	}

	pc.AngleOffset = .15 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f"+" %s", v, files[i].Filename)
	}

	return pc
}

func SetupDeletionChart(files []models.File) *widgets.PieChart {
	deletion := make([]float64, 0)

	for _, file := range files {
		deletion = append(deletion, float64(file.Deletions))

	}

	pc := widgets.NewPieChart()
	pc.Title = "Deletions"
	pc.SetRect(75, 28, 110, 13)
	pc.Colors = []termui.Color{4, 5, 6}
	delNum := len(deletion)
	if delNum == 0 {
		pc.Title = "Deletions changed by commit (no files changed)"
	} else {
		pc.Data = deletion[:delNum]
	}

	pc.AngleOffset = .15 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f"+" %s", v, files[i].Filename)
	}

	return pc
}

func SetupFilesChangedInfo(files []models.File) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Title = "Files changed info"
	p.Text = BuildFileChangedInfo(files)
	p.SetRect(75, 0, 110, 12)

	return p
}

func SetupVerification(vInfo models.Verification) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.Border = true
	p.Title = "Verification info"
	p.Text = BuildCommitVerificationInfo(vInfo)
	p.SetRect(115, 0, 150, 12)

	return p
}

func SetupOverallChart(files []models.File) *widgets.PieChart {
	overalls := make([]float64, 0)

	for _, file := range files {
		overalls = append(overalls, float64(file.Deletions+file.Additions))

	}

	pc := widgets.NewPieChart()
	pc.Title = "Total Changes"
	pc.SetRect(115, 28, 150, 13)
	pc.Colors = []termui.Color{2, 4, 6, 5}
	delNum := len(overalls)
	if delNum == 0 {
		pc.Title = "Total changed by commit (no files changed)"
	} else {
		pc.Data = overalls[:delNum]
	}

	pc.AngleOffset = .15 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f"+" %s", v, files[i].Filename)
	}

	return pc
}
