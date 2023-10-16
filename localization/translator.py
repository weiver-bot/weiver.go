locale = {
    "Danish": 'da',
    "German": 'de',
    "SpanishES": 'es',
    "French": 'fr',
    "Croatian": 'hr',
    "Italian": 'it',
    "Lithuanian": 'lt',
    "Hungarian": 'hu',
    "Dutch": 'nl',
    "Norwegian": 'no',
    "Polish": 'pl',
    "PortugueseBR": 'pt',
    "Romanian": 'ro',
    "Finnish": 'fi',
    "Swedish": 'sv',
    "Vietnamese": 'vi',
    "Turkish": 'tr',
    "Czech": 'cs',
    "Greek": 'el',
    "Bulgarian": 'bg',
    "Russian": 'ru',
    "Ukrainian": 'uk',
    "Hindi": 'hi',
    "Thai": 'th',
    "ChineseCN": 'zh-cn',
    "Japanese": 'ja',
    "ChineseTW": 'zh-tw',
}

commands = {
    "#admin":                   "admin",
    "#admin.allow-role":                   "allow role",
    "#look":             "보기",
    "#move":             "옮기기",
    "#review":             "리뷰",
    "#admin.allow-role.value":             "값",
    "#.subject":             "대상",
    "#look.info":                    "정보",
    "#look.review-list":             "리뷰 목록",
}

text = {
    "#admin.allow-role.Description": "관리자 전용 - 기본값:False",

    "#look.Description": "무언가 보기",

    "#move.Description": "이 채널로 리뷰 옮기기",

    "#review.Description": "유저 리뷰하기",

    "#admin.allow-role.Description":       "점수를 역할로 표시합니다. 기본값: False",
    "#admin.allow-role.value.Description": "값 설정",

    "#.subject.Description": "대상 선택",

    "#look.info.Description":        "유저 정보 보기",
    "#look.review-list.Description": "유저가 받은 리뷰 목록 보기",

    "#allow-role.NeedPermissions":  "봇의 권한이 부족합니다 - 역할 관리",
    "#allow-role.InProgress":       "작업 진행 중",
    "#allow-role.proc.Title":       "옵션 수정",
    "#allow-role.proc.Description": "역할허용",
    "#allow-role.proc.InProgress":  "진행 중",
    "#allow-role.proc.Done":        "완료",
    "#allow-role.Keep":             "바뀐 설정이 없습니다",

    "#look.info.IsNone":            "리뷰 없음",
    "#look.review-list.IsNone":     "리뷰가 존재하지 않습니다",
    "#look.review-list.menu.Title": "%s에 대한 리뷰",
    "#look.review-list.menu.Page":  "%d페이지",

    "#move.IsNone": "대상에게 작성한 리뷰가 없습니다",
    "#move.Move":   "이곳으로 옮기기",

    "#review.SelfReview":    "자신을 리뷰할 수 없습니다",
    "#review.modal.Title":   "%s 리뷰하기",
    "#review.lable.Score":   "점수",
    "#review.lable.Title":   "제목",
    "#review.lable.Content": "내용",

    "$review.IsEdited": "해당 리뷰는 수정되었습니다",
    "$review.NoAuthor": "이곳에 작성자가 존재하지 않아 삭제된 리뷰를 복구할 수 없습니다",
    "$review.DM":       "새로운 리뷰가 작성되었습니다",
}

from google.cloud import translate
client = translate.TranslationServiceClient()
parent = "projects/weiver-localization/locations/global"

def get_translation(text, dest) -> str:
    response = client.translate_text(
        request={
            "parent": parent,
            "contents": [text],
            "mime_type": "text/plain",
            "target_language_code": dest
        }
    )
    return response.translations[0].translated_text

for key in locale.keys():
    print(f"Try {key}...")
    result = {}
    for v in commands.keys():
        t: str = get_translation(commands[v], dest=locale[key]).lower().replace(' ', '-')
        result[v] = t

    for v in text.keys():
        t = get_translation(text[v], dest=locale[key])
        result[v] = t
    
    with open(f'./{key}.go', 'w') as f:
        t = """package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.%s] = map[string]string{"""%key

        for v in result.keys():
            t += '\n\t\t"%s": "%s",'%(v, result[v])

        t += """
    }
}
""" 
        f.write(t)