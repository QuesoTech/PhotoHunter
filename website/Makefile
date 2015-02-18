DEPLOY_DIR = deployable
STATIC_DIR = static
TEMPLATE_DIR = template
MARKDOWN_DIR = md

PANDOC_OPTIONS = -f markdown+raw_html -t html --template $(TEMPLATE_DIR)/template.html --css $(STATIC_DIR)/style.css --toc --toc-depth 2

all: generate deploy

generate: commits
	pandoc $(PANDOC_OPTIONS) --title PhotoHunter $(MARKDOWN_DIR)/index.md -o $(DEPLOY_DIR)/index.html
	pandoc $(PANDOC_OPTIONS) --title 'Meeting Notes' $(MARKDOWN_DIR)/meetings.md -o $(DEPLOY_DIR)/meetings.html
	pandoc $(PANDOC_OPTIONS) --title 'Connor Greenwell' $(MARKDOWN_DIR)/connor.md -o $(DEPLOY_DIR)/connor.html
	pandoc $(PANDOC_OPTIONS) --title 'Ryan Baltenberger' $(MARKDOWN_DIR)/ryan.md -o $(DEPLOY_DIR)/ryan.html
	pandoc $(PANDOC_OPTIONS) --title 'J. David Smith' $(MARKDOWN_DIR)/david.md -o $(DEPLOY_DIR)/david.html
	pandoc $(PANDOC_OPTIONS) --title 'Aaron Bradshaw' $(MARKDOWN_DIR)/aaron.md -o $(DEPLOY_DIR)/aaron.html
	pandoc $(PANDOC_OPTIONS) --title 'Recent Commits' $(MARKDOWN_DIR)/commits.md -o $(DEPLOY_DIR)/commits.html

commits:
	python latest_commits_to_md.py > $(MARKDOWN_DIR)/commits.md

deploy:
	rsync -avz $(DEPLOY_DIR) root@104.131.64.39:/var/www/

spell:
	aspell -c index.md

clean:
	rm $(DEPLOY_DIR)/*.html
