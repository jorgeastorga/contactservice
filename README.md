# Contact Service

## Branches

### Branch: *master*
*This version of the microservice attempts requires a postgres database to connect to.

### Branch: *nodb*
*This version of the microservice only fakes the response to the API calls. 
It is only used as a means for me to learn how to deploy a service in AWS 
without having to deal with the complexity of running a database.* 



## Dependencies :
1. Go programming language installed
2. Go path configured
3. Postgres installed (running on port 5432)
4. Postman to test API calls

## Postman Collections:
1. AWS Postman API Calls: https://documenter.getpostman.com/view/1598740/collection/7LjDkFS
2. Local Postman API Calls: https://documenter.getpostman.com/view/1598740/collection/7LjDkFV


## Getting started with the project
1. Install Go
2. Setup your Go Workspace (e.g. setup $GOPATH)
3. Run: go get github.com/jorgeastorga/contactservice
4. Navigate to: dev/go-workspace/src/github.com/jorgeastorga/contactservice
5. Run: go install
6. Run: contactservice

Note: I won't go into the details of steps #1 and #2 as those are platform-dependent and can be explained elsewhere in more detail.


