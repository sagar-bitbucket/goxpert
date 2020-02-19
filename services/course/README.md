# Course Service 


## Run Course-Service 

Creating a Build

    	$ go build -o service/user/bin/course-service service/user/cmd/main.go

Run user service

    	$ .service/user/bin/user-service mysqldb-host localhost -mysqldb-user rahul -mysqldb-pass password -mysqldb-addr=3306
