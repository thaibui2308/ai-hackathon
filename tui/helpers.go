package tui

import (
	"image"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gizak/termui/v3/widgets"
	"github.com/thaibui2308/ai-hackathon/models"
)

const (
	modified = "(fg:yellow)"
	deleted  = "(fg:red)"
	added    = "(fg:green)"
)

func BuildUserInfo(user models.User) string {
	var baseProfile string
	joinedAt := strings.Split(user.CreatedAt.String(), " ")

	if user.Name == "" {
		baseProfile = "[" + user.Login + ":]" + "(fg:red)" + "\n\n" +
			"[Followers:](fg:yellow) " + strconv.Itoa(user.Followers) + "\n" +
			"[Following:](fg:yellow) " + strconv.Itoa(user.Following) + "\n" +
			"[Joined:](fg:yellow) " + joinedAt[0] + "\n"
	} else {
		baseProfile = "[Username:](fg:red)" + user.Name + "\n\n" +
			"[Followers:](fg:yellow) " + strconv.Itoa(user.Followers) + "\n" +
			"[Following:](fg:yellow) " + strconv.Itoa(user.Following) + "\n" +
			"[Joined:](fg:yellow) " + joinedAt[0] + "\n"
	}

	if user.Location != "" {
		baseProfile += "[Location:](fg:yellow) " + user.Location + "\n"
	}

	if user.Email != "" {
		baseProfile += "[Email:](fg:yellow) " + user.Email + "\n"
	}

	if user.TwitterUsername != "" {
		baseProfile += "[Twitter:](fg:yellow) @" + user.TwitterUsername + "\n"
	}

	return baseProfile
}

func SetupImage(profileImg string, login string) (*widgets.Image, []image.Image) {
	var images []image.Image

	resp, err := http.Get(profileImg)
	if err != nil {
		log.Fatalf("failed to fetch image: %v", err)
	}

	image, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatalf("failed to decode fetched image: %v", err)
	}

	images = append(images, image)

	img := widgets.NewImage(nil)
	img.SetRect(0, 0, 30, 14)
	img.Title = login + "'s GitHub"

	return img, images
}

func BuildStatChecked(summary models.Stats, issues string) string {
	basedStatsChecked := "\n[Issue:](fg:red) " + issues + "\n\n"

	basedStatsChecked += "[Total changes:](fg:yellow) " + strconv.Itoa(summary.Total) + "\n"
	basedStatsChecked += "[Additions:](fg:yellow) " + strconv.Itoa(summary.Additions) + "\n"
	basedStatsChecked += "[Deletions:](fg:yellow) " + strconv.Itoa(summary.Deletions) + "\n"

	return basedStatsChecked
}

func BuildFileChangedInfo(files []models.File) string {
	var baseFileChanged string

	for _, file := range files {
		baseFileChanged += "[" + file.Filename + ":](fg:cyan)" + " [" + file.Status + "]" + FileStatusChecked(file.Status) + "\n"
	}
	return baseFileChanged
}

func FileStatusChecked(status string) string {
	if status == "modified" {
		return modified
	} else if status == "added" {
		return added
	} else if status == "removed" {
		return deleted
	}
	return ""
}

func BuildCommitVerificationInfo(vInfo models.Verification) string {
	var verification string
	verification += "[Verified:](fg:green)" + strconv.FormatBool(vInfo.Verified) + "\n"
	verification += "[Reason:](fg:green)" + vInfo.Reason + "\n"
	return verification
}
