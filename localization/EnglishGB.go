package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Korean] = map[string]string{
		// slashcommand outputs
		"#allow-role":             "AllowRole",
		"#allow-role.Description": "For Admin - defulat:false",

		"#look":             "Look",
		"#look.Description": "Look something",

		"#move-review":             "move-review",
		"#move-review.Description": "Move review to this channel",

		"#review":             "review",
		"#review.Description": "Review user",

		// subcommand outputs
		"#allow-role.value":             "value",
		"#allow-role.value.Description": "Set value",

		"#.subject":             "subject",
		"#.subject.Description": "Select subject",

		"#look.info":                    "info",
		"#look.info.Description":        "Look about info",
		"#look.review-list":             "review-list",
		"#look.review-list.Description": "Look about review-list of user",

		// test outputs
		"#allow-role.InProgress":       "Process is in progress",
		"#allow-role.proc.Title":       "Update option",
		"#allow-role.proc.Description": "AllowRole",
		"#allow-role.proc.InProgress":  "in progress",
		"#allow-role.proc.Done":        "done",
		"#allow-role.Keep":             "Nothing changed",

		"#look.info.IsNone":            "No reviews",
		"#look.review-list.IsNone":     "Review not exist",
		"#look.review-list.menu.Title": "Reviews of %s",
		"#look.review-list.menu.Page":  "page %d",

		"#move-review.IsNone": "No review on this subject",
		"#move-review.Move":   "Move on",

		"#review.SelfReview":    "Can not review yourself",
		"#review.modal.Title":   "Review %s",
		"#review.lable.Score":   "score",
		"#review.lable.Title":   "title",
		"#review.lable.Content": "content",

		"$review.IsEdited": "This review has been edited",
		"$review.NoAuthor": "Author of this deleted review does not exist on here",
		"$review.DM":       "New review has written",
	}
}
