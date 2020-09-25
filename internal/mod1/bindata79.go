package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMqBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMq, nil
}

func dictMod1IpaNeologdDictMq() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
