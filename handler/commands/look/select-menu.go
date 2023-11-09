package look

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	"github.com/y2hO0ol23/weiver/localization"
	"github.com/y2hO0ol23/weiver/utils/builder"
)

func SelectMenu(reviews []db.ReviewModel, locale discordgo.Locale, subjectName string, pageNow int, pageCount int) *builder.SelectMenuStructure {
	// init for building select menu
	pageBack := (pageNow+pageCount-2)%pageCount + 1
	pageNext := pageNow%pageCount + 1

	// set title
	selectMenu := builder.SelectMenu().
		SetCustomID("reviews").
		SetPlaceholder(fmt.Sprintf(localization.Load(locale, "#look.reviews.menu.Title")+" (%v/%v)", subjectName, pageNow, pageCount))

	// if page is not only one
	if pageCount > 1 {
		selectMenu.AddOptions(
			builder.SelectMenuOption().
				SetLabel("▲").
				SetValue(fmt.Sprintf("page/back:%v", pageBack)).
				SetDescription(fmt.Sprintf(localization.Load(locale, "#look.reviews.menu.Page"), pageBack)),
		)
	}

	// add reviews on data
	for i := (pageNow - 1) * pageRow; i < pageNow*pageRow; i++ {
		if i >= len(reviews) {
			break
		}

		review := reviews[i]
		selectMenu.AddOptions(
			builder.SelectMenuOption().
				SetLabel(fmt.Sprintf("%v 〔%v%v〕", review.Title, "★★★★★"[:review.Score*3], "☆☆☆☆☆"[review.Score*3:])).
				SetDescription(fmt.Sprintf("👍 %v", review.LikeTotal)).
				SetValue(fmt.Sprintf("review:%v#%v", review.ID, review.TimeStamp.Unix())),
		)
	}

	// if page is not only one
	if pageCount > 1 {
		selectMenu.AddOptions(
			builder.SelectMenuOption().
				SetLabel("▼").
				SetValue(fmt.Sprintf("page/next:%v", pageNext)).
				SetDescription(fmt.Sprintf(localization.Load(locale, "#look.reviews.menu.Page"), pageNext)),
		)
	}

	return selectMenu
}
