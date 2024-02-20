module github.com/tasuke/go-onion

go 1.22.0

require (
	github.com/stretchr/testify v1.8.4
	github.com/tasuke/go-pkg v0.0.0-00010101000000-000000000000
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/oklog/ulid/v2 v2.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/tasuke/go-pkg => ../pkg
