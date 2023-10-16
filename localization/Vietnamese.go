package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Vietnamese] = map[string]string{
		"#admin":                              "quản-trị-viên",
		"#admin.allow-role":                   "cho-phép-vai-trò",
		"#look":                               "nhìn",
		"#move":                               "di-chuyển",
		"#review":                             "ôn-tập",
		"#admin.allow-role.value":             "giá-trị",
		"#.subject":                           "mục-tiêu",
		"#look.info":                          "thông-tin",
		"#look.review-list":                   "danh-sách-đánh-giá",
		"#admin.allow-role.Description":       "Hiển thị điểm số theo vai trò. Mặc định: Sai",
		"#look.Description":                   "nhìn thấy một cái gì đó",
		"#move.Description":                   "Chuyển bài đánh giá sang kênh này",
		"#review.Description":                 "Đánh giá của người dùng",
		"#admin.allow-role.value.Description": "đặt giá trị",
		"#.subject.Description":               "Chọn mục tiêu",
		"#look.info.Description":              "Xem thông tin người dùng",
		"#look.review-list.Description":       "Xem danh sách đánh giá mà người dùng nhận được",
		"#allow-role.NeedPermissions":         "Bot thiếu quyền - Quản lý vai trò",
		"#allow-role.InProgress":              "công việc đang tiến triển",
		"#allow-role.proc.Title":              "Sửa đổi tùy chọn",
		"#allow-role.proc.Description":        "Quyền vai trò",
		"#allow-role.proc.InProgress":         "Đang tiến hành",
		"#allow-role.proc.Done":               "hoàn thành",
		"#allow-role.Keep":                    "Không có cài đặt nào được thay đổi",
		"#look.info.IsNone":                   "Không có bài đánh giá nào",
		"#look.review-list.IsNone":            "Không có đánh giá nào",
		"#look.review-list.menu.Title":        "Đánh giá cho %s",
		"#look.review-list.menu.Page":         "trang %d",
		"#move.IsNone":                        "Không có đánh giá nào được viết cho mục tiêu",
		"#move.Move":                          "chuyển đến đây",
		"#review.SelfReview":                  "Bạn không thể xem lại chính mình",
		"#review.modal.Title":                 "Xem lại %s",
		"#review.lable.Score":                 "điểm",
		"#review.lable.Title":                 "tiêu đề",
		"#review.lable.Content":               "chi tiết",
		"$review.IsEdited":                    "Đánh giá này đã được chỉnh sửa",
		"$review.NoAuthor":                    "Không thể khôi phục đánh giá đã xóa vì tác giả không tồn tại ở đây",
		"$review.DM":                          "Một đánh giá mới đã được viết",
	}
}
