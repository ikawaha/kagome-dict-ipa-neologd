package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGgBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGg, nil
}

func dictMod0IpaNeologdDictGg() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
