package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGhBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGh, nil
}

func dictMod0IpaNeologdDictGh() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
