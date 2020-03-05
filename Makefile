PROJECT_PATH=${GOPATH}/src/github.com/tribunadigital/dataloaden

regenerate_local:
	- go build ${PROJECT_PATH}
	- go generate ${PROJECT_PATH}/...
	- rm dataloaden