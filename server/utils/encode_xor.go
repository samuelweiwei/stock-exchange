package utils

import "strings"
import "errors"

// Base62 字符集（也可自行调整顺序）
var base62Charset = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 随机的 6 位数字
var salt uint64 = 239200

// base62Encode 将一个正整数编码为 base62 字符串
func base62Encode(num uint64) string {
	if num == 0 {
		return "0"
	}
	var encoded []byte
	base := uint64(len(base62Charset))

	for num > 0 {
		remainder := num % base
		encoded = append([]byte{base62Charset[remainder]}, encoded...)
		num = num / base
	}
	return string(encoded)
}

// base62Decode 将 base62 编码的字符串解码为整数
func base62Decode(str string) (uint64, error) {
	var result uint64
	base := uint64(len(base62Charset))

	// 先做个简单校验
	for _, c := range str {
		if !strings.ContainsRune(string(base62Charset), c) {
			return 0, errors.New("invalid base62 character")
		}
	}

	for _, c := range str {
		pos := strings.IndexRune(string(base62Charset), c)
		result = result*base + uint64(pos)
	}
	return result, nil
}

// EncryptID 通过一个盐 (salt) 对用户 ID 做简单混淆，返回短码
//   - uid: 用户ID（假设是数据库自增ID，或能唯一标识某用户的整形值）
//   - salt: 自定义盐 (必须与 DecryptID 保持一致)
//
// 返回 base62 编码的短串，例如 "35hWnv" 之类
func EncryptID(uid uint64) string {
	// 简单做个 XOR
	mixed := uid ^ salt
	return base62Encode(mixed)
}

// DecryptID 将上面加密过的短码还原回原始用户ID
func DecryptID(shortCode string) (uint64, error) {
	mixed, err := base62Decode(shortCode)
	if err != nil {
		return 0, err
	}
	// 再 XOR 回来
	uid := mixed ^ salt
	return uid, nil
}
