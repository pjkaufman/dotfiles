
.PHONY: html pdf show normal extra csv

normal:
	rm -rf ./stagingGround/*
	cp ~/Notes/Obsidian/Songs/*.md ./stagingGround
extra:
	rm -rf ./stagingGround/*
	cp ~/Notes/Obsidian/Songs/*.md ./stagingGround
	cp ~/Notes/Obsidian/Songs/Extras/*.md ./stagingGround
dark-html:
	@cat ./html/dark.css > ./html/styles.css && cat ./html/common.css >> ./html/styles.css
	@song-converter create-html -d ./stagingGround/ -c ~/Notes/Obsidian/Songs/Cover/churchSongsCover.md -o ./html/churchSongs.html
html:
	@cat ./html/common.css > ./html/styles.css
	@song-converter create-html -d ./stagingGround/ -c ~/Notes/Obsidian/Songs/Cover/churchSongsCover.md -o ./html/churchSongs.html
pdf:
	@echo "Creating PDF"
	weasyprint ./html/churchSongs.html churchSongs.pdf -s ./html/styles.css
show:
	evince churchSongs.pdf
csv:
	song-converter create-csv -d ./stagingGround/ -o churchSongs.csv
