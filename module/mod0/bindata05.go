package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAfBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAf, nil
}

func dictMod0IpaNeologdDictAf() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.af", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
