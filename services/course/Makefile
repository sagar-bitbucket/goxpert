 

GOCMD=go
MYSQL_PASS=--mysqldb-pass
MYSQL_HOST=--mysqldb-host
MYSQL_USER=--mysqldb-user
MYSQL_PORT=--mysqldb-addr
MYSQL_DB_NAME=--mysqldb-name
GOBUILD=$(GOCMD) build


BINARY_NAME=course-service
SERVICE_NAME=admin
BUILD_PATH=./bin/$(BINARY_NAME)

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) cmd/main.go
   
run:
	$(BUILD_PATH) $(MYSQL_HOST) localhost $(MYSQL_USER) rahul $(MYSQL_PASS) password $(MYSQL_PORT)=3306 $(MYSQL_DB_NAME)=goxpert
            
   