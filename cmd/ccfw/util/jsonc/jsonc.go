package jsonc

import (
	"encoding/json"
)

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(toJSON(data), v)
}

// 改行文字のサポート状況
//   - ❌CR（\r）
//   - ✅LF（\n）
//   - ✅CRLF（\r\n）
//
// コメントのサポート状況
//   - ✅シングルライン（// ...）
//   - ❌マルチライン（/* ... */）
//
//nolint:cyclop,gocognit
func toJSON(data []byte) []byte {
	result := make([]byte, 0)
	escaped := false
	inString := false

	// resultにdata[startPos:x]をappendしたら、startPosを更新する...の繰り返し。
	startPos := 0
	for i := 0; i < len(data); i++ {
		c := data[i]

		// 直前がエスケープ文字だった場合、現在の文字を読み飛ばし、フラグをリセットする。
		if escaped {
			escaped = false

			continue
		}

		switch c {
		case '"':
			inString = !inString
		case '\n':
			result = append(result, data[startPos:i+1]...)
			startPos = i + 1
		case '\\': // エスケープ文字だった場合、直後の文字を読み飛ばすためのフラグを立てる。
			if inString {
				escaped = true
			}
		case ',': // ケツカンマを削除する。
			if !inString {
				// カンマの直後（空白文字と改行文字を除く）が `]` または `}` だった場合、ケツカンマとみなす。
				j := i + 1
				for j < len(data) && (data[j] == ' ' || data[j] == '\t' || data[j] == '\n' || data[j] == '\r') {
					j++
				}
				if j < len(data) && (data[j] == ']' || data[j] == '}') {
					result = append(result, data[startPos:i]...)
					startPos = i + 1
				}
			}
		case '/': // コメントを削除する。
			if !inString && i+1 < len(data) && data[i+1] == '/' {
				// コメントの前の空白文字を削除する。
				// 現在地点（コメント開始地点）から空白文字が出現しなくなるまで遡る。
				endPos := i
				for endPos > startPos && (data[endPos-1] == ' ' || data[endPos-1] == '\t') {
					endPos--
				}
				result = append(result, data[startPos:endPos]...)

				// コメント自体を削除する。
				// 現在地点（コメント開始地点）から改行文字が出現するまで読み飛ばす。
				for i < len(data) && data[i] != '\n' {
					i++
				}
				if i < len(data) {
					result = append(result, '\n') // 改行文字はLFに統一
				}
				startPos = i + 1
			}
		}
	}

	// 末尾の処理
	if startPos < len(data) {
		result = append(result, data[startPos:]...)
	}

	return result
}
