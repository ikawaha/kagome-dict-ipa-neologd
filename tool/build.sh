#!/usr/bin/env bash

src="./mecab-ipadic-2.7.0-20070801"
neologd="./seed"
intermediate_dir="./dict"
intermediate_prefix="ipa-neologd-dict."
testdata_dir="../testdata"
dict="${testdata_dir}/ipa-neologd.dict"
dst_pkg=data
dst_dir="../internal"
modules=()



function build_dict() {
  rm -f ${dst_dir}/*.go ${dict}
  mkdir -p ${testdata_dir}
  go run main.go -dict ${src} -other ${neologd} -out ${dict}
}

function build_intermediate_files() {
    rm -rf ${intermediate_dir}
    mkdir -p ${intermediate_dir}
    split -a 2 -b 512k ${dict} ${intermediate_dir}/${intermediate_prefix}

    readonly max_file_num=250
    target="mod"
    target_count=0
    dst="${intermediate_dir}/${target}${target_count}"
    modules=("${modules[@]}" "${target}${target_count}")
    mkdir ${dst}
    i=0
    for f in ./dict/*.*;
    do
      i=$((i+1))
      if [ $i -eq ${max_file_num} ]; then
        i=0
        target_count=$((target_count+1))
        dst="${intermediate_dir}/${target}${target_count}"
        modules=("${modules[@]}" "${target}${target_count}")
        mkdir -p ${dst}
      fi
      mv ${f} ${dst}
    done
}

function build_packages() {
  rm -rf ${dst_dir}
  for m in "${modules[@]}";
  do
    mkdir -p ${dst_dir}/${m}
    go-bindata -o ${dst_dir}/${m}/bindata.go -nocompress -separate -pkg=${dst_pkg} ${intermediate_dir}/${m}
    echo "module github.com/ikawaha/kagome-dict-ipa-neologd/internal/${m}\n\ngo 1.15" > ${dst_dir}/${m}/go.mod

  done
  rm -rf ${intermidiate_dir}
}


# main
build_dict
build_intermediate_files
build_packages



