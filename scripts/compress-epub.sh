#!/bin/bash
echo 'running compression for each epub in dir'

shopt -s nocaseglob # ignore file name case
shopt -s globstar # search globs in 0 or more subdirectories
for i in *.epub; do
    echo 'starting epub compressing for' ${i} '...'

    rm -rf ./epub #delete the folder if it already exists

    unzip -qq "${i}" -d epub
    
    # check for jpg/jpeg and png images
    for j in **/*.{jpg,jpeg,png}; do
        imgp -x 800x800 -e -O -q 40 -w -m "${j}"
    done

    # rezip the epub
    cd epub
    zip -q -0 -X "../compress.epub" mimetype
    zip -q -rDX9 "../compress.epub" * -x mimetype
    cd ..

    echo '-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-'

    # remove our work directory
    rm -rf ./epub

    mv "${i}" "${i}.original"
    mv compress.epub "${i}"

    du -h "${i}.original"
    du -h "${i}"

    echo '-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-'
done

# remove changes to bash search interpretation
shopt -u nocaseglob
shopt -u globstar

echo '-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-'

echo 'after:'
find . -type f -name '*.epub' -exec du -ch {} + | grep total$
echo 'before:'
find . -type f -name '*.original' -exec du -ch {} + | grep total$ 

echo '-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-'
