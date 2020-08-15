package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBlBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBl, nil
}

func dictMod0IpaNeologdDictBl() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bl", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
