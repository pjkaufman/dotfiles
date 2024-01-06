# Church Song Converter

This is a program takes markdown files and converts them to html. The html files are then converted to a pdf file.

## How It Works

There is a make file that lets you copy markdown files to the `stagingGround` folder for either the regular (`make normal`)
or extra (`make extra`) markdown file songs.

Then you can either create the dark or regular style html via `make dark-html` or `make html`.
The generated html ist stored in `html`.

Lastly you can create the pdf via `make pdf` which uses `weasyprint`.

## Csv

There is also a csv file that can be created which holds metadata for the different songs that exist.
