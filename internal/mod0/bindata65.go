package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCnBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCn, nil
}

func dictMod0IpaNeologdDictCn() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
