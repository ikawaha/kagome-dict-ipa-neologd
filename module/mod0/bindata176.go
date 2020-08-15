package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGuBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGu, nil
}

func dictMod0IpaNeologdDictGu() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
