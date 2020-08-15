package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictByBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBy, nil
}

func dictMod0IpaNeologdDictBy() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictByBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.by", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
