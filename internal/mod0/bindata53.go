package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCbBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCb, nil
}

func dictMod0IpaNeologdDictCb() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
