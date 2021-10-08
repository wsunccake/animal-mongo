# animal-mongo

## command

```bash
# init
go mod init animal-mongo	    # generate go.mod
go mod tiny        	            # remove unused module
go get 				            # download all package

# build
go build			            # build binary

# test
go test -v                      # only animal-mongo folder test case
go test animal-mongo/utils -v   # animal-mongo/utils folder test case
go test ./... -v                # all test case
```
