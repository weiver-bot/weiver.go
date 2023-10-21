package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Turkish] = map[string]string{
		"#admin":                              "yönetici",
		"#admin.allow-role":                   "role-izin-ver",
		"#look":                               "bakmak",
		"#move":                               "taşınmak",
		"#review":                             "gözden-geçirmek",
		"#admin.allow-role.value":             "değer",
		"#.subject":                           "hedef",
		"#look.info":                          "bilgi",
		"#look.reviews":                       "yorumlar",
		"#admin.allow-role.Description":       "Puanları roller olarak görüntüleyin. Varsayılan: Yanlış",
		"#look.Description":                   "bir şeyler gör",
		"#move.Description":                   "İncelemeyi bu kanala taşı",
		"#review.Description":                 "Kullanıcı Değerlendirmesi",
		"#admin.allow-role.value.Description": "değeri belirle",
		"#.subject.Description":               "Hedef seç",
		"#look.info.Description":              "Kullanıcı bilgilerini görüntüle",
		"#look.reviews.Description":           "Konunun aldığı değerlendirmeleri görün",
		"#admin.allow-role.NeedPermissions":   "Botun izinleri yok - Rolleri Yönet",
		"#admin.allow-role.InProgress":        "devam eden çalışma",
		"#admin.allow-role.proc.Title":        "Seçenekleri Değiştir",
		"#admin.allow-role.proc.Description":  "Rollere izin ver",
		"#admin.allow-role.proc.InProgress":   "devam ediyor",
		"#admin.allow-role.proc.Done":         "tamamlamak",
		"#admin.allow-role.Keep":              "Hiçbir ayar değişmedi",
		"#look.info.IsNone":                   "Yorum yok",
		"#look.reviews.IsNone":                "Hiç yorum yok",
		"#look.reviews.menu.Title":            "%v için incelemeler",
		"#look.reviews.menu.Page":             "%v sayfa",
		"#move.IsNone":                        "Hedef için yazılmış bir inceleme yok",
		"#move.Move":                          "Buraya taşın",
		"#review.SelfReview":                  "Kendinizi gözden geçiremezsiniz",
		"#review.modal.Title":                 "%v'yi inceleyin",
		"#review.lable.Score":                 "Gol",
		"#review.lable.Title":                 "başlık",
		"#review.lable.Content":               "detay",
		"$review.IsEdited":                    "Bu inceleme düzenlendi",
		"$review.NoAuthor":                    "Silinen yorumlar geri yüklenemez çünkü yazar burada mevcut değildir",
		"$review.DM":                          "Yeni bir inceleme yazıldı",
	}
}
