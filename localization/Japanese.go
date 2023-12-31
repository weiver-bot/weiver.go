package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Japanese] = map[string]string{
		"#admin":                              "管理者",
		"#admin.allow-role":                   "ロールを許可する",
		"#look":                               "見る",
		"#move":                               "移動する",
		"#review":                             "レビュー",
		"#admin.allow-role.value":             "値",
		"#.subject":                           "対象",
		"#look.info":                          "情報",
		"#look.reviews":                       "レビュー",
		"#admin.allow-role.Description":       "スコアを役割として表示します。既定値: False",
		"#look.Description":                   "何かを見る",
		"#move.Description":                   "このチャンネルにレビューを移動する",
		"#review.Description":                 "ユーザーレビュー",
		"#admin.allow-role.value.Description": "値の設定",
		"#.subject.Description":               "ターゲットを選択",
		"#look.info.Description":              "ユーザー情報を見る",
		"#look.reviews.Description":           "対象のレビューを見る",
		"#admin.allow-role.NeedPermissions":   "ボットの権限がない - 役割の管理",
		"#admin.allow-role.InProgress":        "作業中",
		"#admin.allow-role.proc.Title":        "オプションの変更",
		"#admin.allow-role.proc.Description":  "役割を許可",
		"#admin.allow-role.proc.InProgress":   "進行中",
		"#admin.allow-role.proc.Done":         "完了",
		"#admin.allow-role.Keep":              "変更設定はありません",
		"#look.info.IsNone":                   "レビューなし",
		"#look.reviews.IsNone":                "レビューは存在しません",
		"#look.reviews.menu.Title":            "%vのレビュー",
		"#look.reviews.menu.Page":             "%vページ",
		"#move.IsNone":                        "ターゲットに投稿したレビューはありません",
		"#move.Move":                          "ここに移動",
		"#review.SelfReview":                  "自分をレビューできません",
		"#review.modal.Title":                 "%vレビュー",
		"#review.lable.Score":                 "スコア",
		"#review.lable.Title":                 "タイトル",
		"#review.lable.Content":               "内容",
		"$review.IsEdited":                    "このレビューは修正されました",
		"$review.NoAuthor":                    "ここに作成者が存在しないため、削除されたレビューを復元できません",
		"$review.DM":                          "新しいレビューが作成されました",
	}
}
