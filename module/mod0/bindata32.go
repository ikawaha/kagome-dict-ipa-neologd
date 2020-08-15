package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBgBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBg, nil
}

func dictMod0IpaNeologdDictBg() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
