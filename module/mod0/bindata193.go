package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHlBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHl, nil
}

func dictMod0IpaNeologdDictHl() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hl", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
