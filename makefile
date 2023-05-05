.PHONY: dev

dev: export COURSE_PLAN_POSTGRESQL_DSN=host=localhost port=5432 user=admin password=password dbname=postgres sslmode=disable
dev: export ENDPOINT=www.xcuitech.com:9002
dev: export APIHOST=https://www.xcuitech.com:9002
dev: export ACCESSKEYID=admin
dev: export ACCESSKEY=password

dev: export RESOURCES_PHOTO_PATH=res-photo
dev: export WALLPAPER_PATH=wallpaper
dev: export OPEN_AI_API_KEY=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

dev:
	@go run main.go