
.PHONY: html pdf show normal extra

normal:
	rm -rf ./stagingGround/*
	cp ~/Notes/Obsidian/Songs/*.md ./stagingGround
extra:
	rm -rf ./stagingGround/*
	cp ~/Notes/Obsidian/Songs/*.md ./stagingGround
	cp ~/Notes/Obsidian/Songs/Extras/*.md ./stagingGround
html:
	./createSongsPdf.sh 
pdf:
	echo "Creating PDF"
	wkhtmltopdf --page-size A5 cover ./html/churchSongsCover.html toc --xsl-style-sheet ./html/updatedToc.xsl ./html/churchSongs.html churchSongsTest.pdf
show:
	evince churchSongsTest.pdf