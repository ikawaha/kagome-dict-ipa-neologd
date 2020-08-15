package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBtBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBt, nil
}

func dictMod0IpaNeologdDictBt() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bt", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
