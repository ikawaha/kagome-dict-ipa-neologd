package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBqBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBq, nil
}

func dictMod0IpaNeologdDictBq() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
