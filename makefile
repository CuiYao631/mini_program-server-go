.PHONY: dev

dev: export COURSE_PLAN_POSTGRESQL_DSN=host=localhost port=5432 user=admin password=password dbname=postgres sslmode=disable
dev: export ENDPOINT=www.xcuitech.com:9000
dev: export APIHOST=http://www.xcuitech.com:9000
dev: export ACCESSKEYID=admin
dev: export ACCESSKEY=password

dev: export RESOURCES_PHOTO_PATH=res-photo
dev: export WALLPAPER_PATH=wallpaper

dev:
	@go run main.go