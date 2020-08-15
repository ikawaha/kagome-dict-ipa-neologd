package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAgBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAg, nil
}

func dictMod0IpaNeologdDictAg() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ag", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
