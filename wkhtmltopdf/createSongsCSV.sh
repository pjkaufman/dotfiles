#!/bin/bash

# https://stackoverflow.com/questions/5014632/how-can-i-parse-a-yaml-file-from-a-linux-shell-script
function parse_yaml {
   local prefix=$2
   local s='[[:space:]]*' w='[a-zA-Z0-9_\-]*' fs=$(echo @|tr @ '\034')
   sed -ne "s|^\($s\):|\1|" \
        -e "s|^\($s\)\($w\)$s:$s[\"']\(.*\)[\"']$s\$|\1$fs\2$fs\3|p" \
        -e "s|^\($s\)\($w\)$s:$s\(.*\)$s\$|\1$fs\2$fs\3|p"  "$1" |
   awk -F$fs '{
      indent = length($1)/2;
      vname[indent] = $2;
      for (i in vname) {if (i > indent) {delete vname[i]}}
      if (length($3) > 0) {
         vn=""; for (i=0; i<indent; i++) {vn=(vn)(vname[i])("_")}
         printf("%s%s%s=\"%s\"\n", "'$prefix'",vn, $2, $3);
      }
   }'
}

function parse_string_key {
    value="${yaml#*$1=\"}"
    value=$(echo $value | cut -f1 -d"\"")
    if [[ "$value" =~ "=" ]]; then
      echo ""
    else
      echo $value
    fi
}

function build_metadata_csv {
    authors=$(parse_string_key "authors")
    inChurch=$(parse_string_key "in-church")
    location=$(parse_string_key "location")

    metadataFound=0

    if [ "$authors" != "" ]; then 
      let metadataFound+=1
    fi

    if [ "$location" != "" ]; then 
      let metadataFound+=1
    fi

    if [ $metadataFound != 0 ]; then 
        # location
        if [ "$location" != "" ]; then
          echo "$location|"
        else 
          echo "|" 
        fi

        # authors
        if [ "$authors" != "" ]; then
          if [ "$inChurch" = "Y" ]; then 
            echo "$authors|Church"
          else 
            echo "$authors|"
          fi
        else
          echo "|" 
        fi
    else 
      echo "||"
    fi
}

echo "Converting Markdown files to csv"

echo -e "Song|Location|Author|Copyright" > ./churchSongs.csv

find ./stagingGround/*.md -maxdepth 0 -printf "%f\n" | LC_ALL=C sort -k 1.1f,1.1 > ./html/sortedFiles.txt

while read f; do
    fileName=$(basename "$f" .md)
    yaml=$(parse_yaml "./stagingGround/$f")
    csvLine="$fileName|"

    if [ -n "$yaml" ]; then
      metadata=$(build_metadata_csv $yaml)
      csvLine="$fileName|$metadata"
    fi
    
    echo -e $csvLine >> ./churchSongs.csv
done <./html/sortedFiles.txt