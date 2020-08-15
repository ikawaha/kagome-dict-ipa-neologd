package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBiBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBi, nil
}

func dictMod0IpaNeologdDictBi() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bi", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
