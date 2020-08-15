package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAbBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAb, nil
}

func dictMod0IpaNeologdDictAb() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ab", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
