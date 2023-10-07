
.PHONY: html pdf show normal extra csv

normal:
	rm -rf ./stagingGround/*
	cp ~/Notes/Obsidian/Songs/*.md ./stagingGround
extra:
	rm -rf ./stagingGround/*
	cp ~/Notes/Obsidian/Songs/*.md ./stagingGround
	cp ~/Notes/Obsidian/Songs/Extras/*.md ./stagingGround
html:
	song-converter create-cover -f ~/Notes/Obsidian/Songs/Cover/churchSongsCover.md -o ./html/cover.html
	song-converter md-to-html -d ./stagingGround/ -s ./html/styles.html -o ./html/songs.html 
pdf:
	echo "Creating PDF"
	wkhtmltopdf --page-size A5 cover ./html/cover.html toc --xsl-style-sheet ./html/updatedToc.xsl ./html/songs.html churchSongsTest.pdf
show:
	evince churchSongsTest.pdf
csv:
	song-converter create-csv -d ./stagingGround/ -o churchSongs.csv