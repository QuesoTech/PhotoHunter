all: generate

generate:
	pandoc -o index.html -f markdown+raw_html -t html --template template.html --css template.css --toc --toc-depth 2 index.md
