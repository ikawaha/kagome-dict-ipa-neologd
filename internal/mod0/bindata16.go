package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAqBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAq, nil
}

func dictMod0IpaNeologdDictAq() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.aq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
