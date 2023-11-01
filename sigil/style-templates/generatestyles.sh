#!/usr/bin/env bash

# takes in the template file and replaces all variables withs their values as per the color file

colorFile=colors.sass
qtStyleFile=$(cat qt_styles_template.sass)

while IFS= read -r line
do
  if [ ! -z "$line" ] && [[ ! "$line" =~ ^\ +$ ]]; then
    IFS=':' read -ra colorParts <<< "$line"
    qtStyleFile=$(sed "s/${colorParts[0]}/${colorParts[1]}/g" <<< "$qtStyleFile")
  fi
done < "$colorFile"

echo "$qtStyleFile" > ../qt_styles.qss
