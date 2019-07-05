package method

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"regexp"
)

func RegexFind(pat, str string) []string {
	var patRegex = regexp.MustCompile(pat)
	strs := patRegex.FindStringSubmatch(str)
	return strs
}

func GetMD5(obj interface{}) (string, error) {
	byteObj, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	md5Inst := md5.New()
	md5Inst.Write(byteObj)
	result := hex.EncodeToString(md5Inst.Sum(nil))

	return result, nil
}

func IsTheSame(refer, compare interface{}) bool {
	byteRefer, err := json.Marshal(refer)
	if err != nil {
		return false
	}

	byteCompare, err := json.Marshal(compare)
	if err != nil {
		return false
	}

	md5Refer := md5.New()
	md5Refer.Write(byteRefer)
	referHash := hex.EncodeToString(md5Refer.Sum(nil))

	md5Compare := md5.New()
	md5Compare.Write(byteCompare)
	compareHash := hex.EncodeToString(md5Compare.Sum(nil))

	return referHash == compareHash
}
