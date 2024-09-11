test_schema_convertor:
	go test -v ./tests/convertor_test.go

test_projects_crud:
	go test -v ./tests/projects_test.go

test_schemas_crud:
	go test -v ./tests/schemas_test.go

release:
	goreleaser release --clean

release_check:
	goreleaser check
