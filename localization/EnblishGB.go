package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.EnglishGB] = map[string]string{
		// slashcommand outputs
		"#admin":             "admin",
		"#admin.Description": "Administrator only",

		"#look":             "look",
		"#look.Description": "Look something",

		"#move":             "move",
		"#move.Description": "Move review to this channel",

		"#review":             "review",
		"#review.Description": "Review user",

		// subcommand outputs
		"#admin.allow-role":                   "allow-role",
		"#admin.allow-role.Description":       "Display scores as roles. Default: False",
		"#admin.allow-role.value":             "value",
		"#admin.allow-role.value.Description": "Set value",

		"#.subject":             "subject",
		"#.subject.Description": "Select subject",

		"#look.info":                "info",
		"#look.info.Description":    "Look about information",
		"#look.reviews":             "reviews",
		"#look.reviews.Description": "Look reviews the subject has received",

		// test outputs
		"#admin.allow-role.NeedPermissions":  "Bot lacks permissions - Manage Roles",
		"#admin.allow-role.InProgress":       "Process is in progress",
		"#admin.allow-role.proc.Title":       "Update option",
		"#admin.allow-role.proc.Description": "Allow role",
		"#admin.allow-role.proc.InProgress":  "in progress",
		"#admin.allow-role.proc.Done":        "done",
		"#admin.allow-role.Keep":             "Nothing changed",

		"#look.info.IsNone":        "No reviews",
		"#look.reviews.IsNone":     "Review not exist",
		"#look.reviews.menu.Title": "Reviews to %v",
		"#look.reviews.menu.Page":  "page %v",

		"#move.IsNone": "No review on this subject",
		"#move.Move":   "Move here",

		"#review.SelfReview":    "Cannot review yourself",
		"#review.modal.Title":   "Review %v",
		"#review.lable.Score":   "score",
		"#review.lable.Title":   "title",
		"#review.lable.Content": "content",

		"$review.IsEdited": "This review has been edited",
		"$review.NoAuthor": "Deleted reviews cannot be restored because the author does not exist here",
		"$review.DM":       "New review has written",
	}
}
