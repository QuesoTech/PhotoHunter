all: deploy

deploy:
	rsync -avz ./* root@104.131.64.39:/var/www/researchers/

spell:
	aspell -c index.md
