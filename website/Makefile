all: generate deploy

generate:
	pandoc -f markdown+raw_html -t html --template template.html --css template.css --toc --toc-depth 2 --title PhotoHunter index.md -o index.html
	pandoc -f markdown+raw_html -t html --template template.html --css template.css --toc --toc-depth 2 --title 'Meeting Notes' meetings.md -o meetings.html
	pandoc -f markdown+raw_html -t html --template template.html --css template.css --toc --toc-depth 2 --title 'Connor Greenwell' connor.md -o connor.html
	pandoc -f markdown+raw_html -t html --template template.html --css template.css --toc --toc-depth 2 --title 'Ryan Baltenberger' ryan.md -o ryan.html
	pandoc -f markdown+raw_html -t html --template template.html --css template.css --toc --toc-depth 2 --title 'J. David Smith' david.md -o david.html
	pandoc -f markdown+raw_html -t html --template template.html --css template.css --toc --toc-depth 2 --title 'Aaron Bradshaw' aaron.md -o aaron.html
	python latest_commits_to_md.py | pandoc -f markdown+raw_html -t html --template template.html --css template.css --toc --toc-depth 2 --title 'Commits' -o commits.html

deploy:
	rsync -avz * root@104.131.64.39:/var/www/

spell:
	aspell -c index.md
