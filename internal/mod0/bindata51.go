package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBzBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBz, nil
}

func dictMod0IpaNeologdDictBz() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
