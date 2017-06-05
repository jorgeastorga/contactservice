FROM golang

ADD . /go/src/github.com/jorgeastorga/contactservice
WORKDIR /go/src/github.com/sausheong/contactservice

#Fetch contact service go dependencies
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
RUN go get github.com/lib/pq

#Install the service
RUN go install github.com/jorgeastorga/contactservice

#Run the file
ENTRYPOINT /go/bin/contactservice

EXPOSE 8080
