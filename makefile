.PHONY: dev

dev: export COURSE_PLAN_POSTGRESQL_DSN=host=localhost port=5432 user=admin password=password dbname=postgres sslmode=disable
dev: export ENDPOINT=tencent.xcuitech.com:1688
dev: export ACCESSKEYID=admin
dev: export ACCESSKEY=password

dev:
	@go run main.go