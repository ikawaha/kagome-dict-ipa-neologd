package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBaBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBa, nil
}

func dictMod0IpaNeologdDictBa() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ba", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
