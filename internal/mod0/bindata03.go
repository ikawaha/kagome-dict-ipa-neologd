package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAdBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAd, nil
}

func dictMod0IpaNeologdDictAd() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ad", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
