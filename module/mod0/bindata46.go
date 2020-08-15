package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBuBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBu, nil
}

func dictMod0IpaNeologdDictBu() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
