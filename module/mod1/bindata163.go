package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPwBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPw, nil
}

func dictMod1IpaNeologdDictPw() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pw", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
