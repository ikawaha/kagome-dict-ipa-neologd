package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCyBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCy, nil
}

func dictMod0IpaNeologdDictCy() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cy", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
