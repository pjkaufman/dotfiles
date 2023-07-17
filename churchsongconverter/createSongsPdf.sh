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

function build_metadata_div {
    melody=$(parse_string_key "melody")
    key=$(parse_string_key "key")
    authors=$(parse_string_key "authors")
    inChurch=$(parse_string_key "in-church")
    verse=$(parse_string_key "verse")
    location=$(parse_string_key "location")

    metadataFound=0
    row1=0
    row2=0
    if [ "$melody" != "" ]; then 
      let metadataFound+=1
      let row2+=1
    fi

    if [ "$verse" != "" ]; then 
      let metadataFound+=1
      let row2+=1
    fi

    if [ "$authors" != "" ]; then 
      let metadataFound+=1
      let row1+=1
    fi

    if [ "$key" != "" ]; then 
      let metadataFound+=1
      let row1+=1
    fi

    if [ "$location" != "" ]; then 
      let metadataFound+=1
      let row1+=1
    fi

    if [ $metadataFound != 0 ]; then 
      width=$(expr 100 / $metadataFound)
      echo "<div>"

      if [ $row1 != 0 ]; then
        if [ $row2 != 0 ]; then
          echo "<div class=\"metadata row-padding\">"
        else
          echo "<div class=\"metadata\">"
        fi

        # authors
        echo "<div><div class=\"author\">"
        if [ "$authors" != "" ]; then
          if [ "$inChurch" = "Y" ]; then 
            echo "<b>$authors</b>"
          else 
            echo "$authors"
          fi
        else
          echo "&nbsp;&nbsp;&nbsp;&nbsp;" 
        fi
        echo "</div></div>"

        # song key
        echo "<div><div class=\"key\">"
        if [ "$key" != "" ]; then
          echo "<b>$key</b>"
        else 
          echo "&nbsp;&nbsp;&nbsp;&nbsp;" 
        fi
        echo "</div></div>"

        # book location
        echo "<div><div class=\"location\">"
        if [ "$location" != "" ]; then
          echo "$location"
        else 
          echo "&nbsp;&nbsp;&nbsp;&nbsp;" 
        fi
        echo "</div></div>"
        
        echo "</div>"
      fi

      if [ $row2 == 1 ] && [ "$melody" != "" ]; then
        echo "<div class=\"metadata\">"
        echo "<div><div class=\"melody-75\">"
        echo "<b>$melody</b>"
        echo "</div></div></div>"
      elif [ $row2 != 0 ]; then
        echo "<div class=\"metadata\">"

        # melody
        echo "<div><div class=\"melody\">"
        if [ "$melody" != "" ]; then
          echo "<b>$melody</b>"
        else 
        echo "&nbsp;&nbsp;&nbsp;&nbsp;"
        fi
        echo "</div></div>"
      
        # verse reference
        echo "<div><div class=\"verse\">"
        if [ "$verse" != "" ]; then
          echo "$verse"
        else
          echo "&nbsp;&nbsp;&nbsp;&nbsp;" 
        fi
        echo "</div></div>"

        echo "</div>"
      fi

      echo "</div><br/><br/>"
    else 
      echo ""
    fi
}

echo "Converting Markdown files to html"

echo -e "$(cat "./html/styles.html")" > ./html/churchSongs.html

pandoc ~/Notes/Obsidian/Songs/Cover/churchSongsCover.md -o ./html/churchSongsCover.html

echo -e "<div style=\"text-align: center\">\n$(cat ./html/churchSongsCover.html)" > ./html/churchSongsCover.html
echo -e "</div>\n" >> ./html/churchSongsCover.html

find ./stagingGround/*.md -maxdepth 0 -printf "%f\n" | LC_ALL=C sort -k 1.1f,1.1 > ./html/sortedFiles.txt

while read f; do
    fileName=$(basename "$f" .md)
    yaml=$(parse_yaml "./stagingGround/$f")
    pandoc "./stagingGround/$f" -o "./html/build/$fileName.html"

    if [ -n "$yaml" ]; then
      metadata=$(build_metadata_div $yaml)
      sed -i "/<\/h1>/a ${metadata@Q}" "./html/build/$fileName.html"
      sed -i "s/\$'</</" "./html/build/$fileName.html" 
      sed -i "s/>'/>/" "./html/build/$fileName.html" 
    fi
    
    sed -i -r 's/^(<h1.*)\((.*)\)<(.*)/\1<span class="other-title">\(\2\)<\/span><\3/' "./html/build/$fileName.html"
    sed -i "/<hr \/>/{N;d;}" "./html/build/$fileName.html"
    sed -i "/''/d" "./html/build/$fileName.html"
    echo -e "<div class=\"keep-together\">\n$(cat "./html/build/$fileName.html")" > "./html/build/$fileName.html"
    echo -e "</div>\n<br/>" >> "./html/build/$fileName.html"
    echo -e "$(cat "./html/build/$fileName.html")" >> ./html/churchSongs.html
done <./html/sortedFiles.txt