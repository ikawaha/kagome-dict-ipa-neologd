package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAhBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAh, nil
}

func dictMod0IpaNeologdDictAh() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ah", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
