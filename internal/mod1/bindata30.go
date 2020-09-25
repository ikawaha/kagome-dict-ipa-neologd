package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKtBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKt, nil
}

func dictMod1IpaNeologdDictKt() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kt", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
