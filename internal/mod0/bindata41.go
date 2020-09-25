package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBpBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBp, nil
}

func dictMod0IpaNeologdDictBp() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
