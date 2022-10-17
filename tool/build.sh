#!/usr/bin/env bash

src="./mecab-ipadic-2.7.0-20070801"
neologd="./seed"
dest_dir="../dict"
dict_prefix="ipa-neologd-dict."
testdata_dir="../testdata"
dict="${testdata_dir}/ipa-neologd.dict"

function build_dict() {
  rm -f ${dst_dir}/*.go ${dict}
  mkdir -p ${testdata_dir}
  go run main.go -dict ${src} -other ${neologd} -out ${dict}
}

function build_intermediate_files() {
    rm -rf ${dest_dir}
    mkdir -p ${dest_dir}
    split -a 2 -b 100m ${dict} ${dest_dir}/${dict_prefix}
}

# main
build_dict
build_intermediate_files



