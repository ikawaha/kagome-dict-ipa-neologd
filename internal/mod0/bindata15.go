package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictApBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAp, nil
}

func dictMod0IpaNeologdDictAp() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictApBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ap", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
