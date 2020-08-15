package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBhBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBh, nil
}

func dictMod0IpaNeologdDictBh() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
