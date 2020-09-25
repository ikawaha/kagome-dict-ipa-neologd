package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPqBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPq, nil
}

func dictMod1IpaNeologdDictPq() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
