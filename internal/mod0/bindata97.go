package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDtBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDt, nil
}

func dictMod0IpaNeologdDictDt() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dt", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
