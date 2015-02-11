all: generate

generate:
	pandoc -f markdown+raw_html -t html --template template.html --css template.css --toc --toc-depth 2 --title PhotoHunter index.md -o index.html
