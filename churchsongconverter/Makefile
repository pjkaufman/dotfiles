
.PHONY: html pdf show normal extra csv

normal:
	rm -rf ./stagingGround/*
	cp ~/Notes/Obsidian/Songs/*.md ./stagingGround
extra:
	rm -rf ./stagingGround/*
	cp ~/Notes/Obsidian/Songs/*.md ./stagingGround
	cp ~/Notes/Obsidian/Songs/Extras/*.md ./stagingGround
dark-html:
	@$(eval type ?= 'Abridged')
	@cat ./html/dark.css > ./html/styles.css && cat ./html/common.css >> ./html/styles.css
	@song-converter create-html -d ./stagingGround/ -c ~/Notes/Obsidian/Songs/Cover/churchSongsCover.md -o ./html/churchSongs.html -v $(type)
html:
	@$(eval type ?= 'Abridged')
	@cat ./html/common.css > ./html/styles.css
	@song-converter create-html -d ./stagingGround/ -c ~/Notes/Obsidian/Songs/Cover/churchSongsCover.md -o ./html/churchSongs.html -v $(type)
pdf:
	@echo "Creating PDF"
	weasyprint ./html/churchSongs.html churchSongs.pdf -s ./html/styles.css
odt:
	# Still needs a lot of work to get it ready. May need to rework this to make it use an odt template
	@pandoc -t odt -o churchSongs.odt --css ./html/styles.css ./html/churchSongs.html
show:
	evince churchSongs.pdf
csv:
	song-converter create-csv -d ./stagingGround/ -o churchSongs.csv
all:
	@echo "Creating Abridged Songs"
	@make normal
	@make html && make pdf
	@mv churchSongs.pdf churchSongsAbridged.pdf
	@make dark-html && make pdf
	@mv churchSongs.pdf churchSongsAbridgedDark.pdf
	@echo "Creating Unabridged Songs"
	@make extra
	@make html type=Unabridged && make pdf
	@mv churchSongs.pdf churchSongsUnabridged.pdf
	@make dark-html type=Unabridged && make pdf
	@mv churchSongs.pdf churchSongsUnabridgedDark.pdf
	@echo "Done"
