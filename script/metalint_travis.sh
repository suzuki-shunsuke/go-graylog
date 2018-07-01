# I can't understand the reason but it is failed to run "gometalinter ./..." travis ci
# index_set.go:1::warning: file is not goimported (goimports)
# index_set.go:1::warning: file is not gofmted with -s (gofmt)
go list ./... | sed -e "s/.*\/go-graylog/./" | xargs gometalinter
