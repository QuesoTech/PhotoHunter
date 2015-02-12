all: generate deploy

generate:
	pandoc -f markdown+raw_html -t html --template template.html --css template.css --toc --toc-depth 2 --title PhotoHunter index.md -o index.html

deploy:
	rsync -avz * root@104.131.64.39:/var/www/

spell:
	aspell -c index.md
