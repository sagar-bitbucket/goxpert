package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"

	service "gitlab.com/scalent/goxpert/services/course/pkg/v1/services"

	repository "gitlab.com/scalent/goxpert/services/course/pkg/v1/repositories"

	"gitlab.com/scalent/goxpert/services/course/pkg/v1/endpoints"
	handler "gitlab.com/scalent/goxpert/services/course/pkg/v1/handlers"
	"gitlab.com/scalent/goxpert/services/course/pkg/v1/middleware"

	"gitlab.com/scalent/goxpert/database"

	"net/http"

	"github.com/gorilla/mux"
)

var (
	//HTTPPortFlag CLI flag for http port
	HTTPPortFlag = "http-port"

	//MysqlDBHostFlag CLI flag for mysqldb host
	MysqlDBHostFlag = "mysqldb-host"

	//MysqlDBUserFlag CLI flag for mysqldb username
	MysqlDBUserFlag = "mysqldb-user"

	//MysqlDBPassFlag CLI flag for mysqldb password
	MysqlDBPassFlag = "mysqldb-pass"

	//MysqlDBAddrFlag CLI flag for mysqldb port
	MysqlDBAddrFlag = "mysqldb-addr"

	//DebugPortFlag CLI flag for debug address
	DebugPortFlag = "debug-addr"

	//MysqlDBNameFlag CLI flag for database name
	MysqlDBNameFlag = "mysqldb-name"

	//HTTPPortEnvVar Environment Varriable
	HTTPPortEnvVar = "HTTP_PORT"

	//MysqlDBHostEnvVar Environment Varriable
	MysqlDBHostEnvVar = "MYSQL_DB_HOST"

	//MysqlDBUserEnvVar Environment Varriable
	MysqlDBUserEnvVar = "MYSQL_DB_USER"

	//MysqlDBPassEnvVar Environment Varriable
	MysqlDBPassEnvVar = "MYSQL_DB_PASS"

	//MysqlDBAddrEnvVar Environment Varriable
	MysqlDBAddrEnvVar = "MYSQL_DB_PORT"

	//MysqlDBNameEnvVar Environment Varriable
	MysqlDBNameEnvVar = "MYSQL_DB_NAME"
)

// Define our flags. Your service probably won't need to bind listeners for
// all supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("Goxpert", flag.ExitOnError)

//HTTPAddr http Port
var HTTPAddr = fs.String(HTTPPortFlag, "8080", "HTTP listen address defaults to 8080")

//MysqlDBHost mysqldb  hostname
var MysqlDBHost = fs.String(MysqlDBHostFlag, "", "Hostname for mysqlDB")

//MysqlDBUser mysqldb username
var MysqlDBUser = fs.String(MysqlDBUserFlag, "", "Username for mysqlDB")

//MysqlDBPass mysqldb password
var MysqlDBPass = fs.String(MysqlDBPassFlag, "", "Password for mysqlDB")

//MysqlDBAddr mysqldb port address
var MysqlDBAddr = fs.String(MysqlDBAddrFlag, "", "Port Number for mysqlDB defaults to 3306")

//MysqlDBName mysqldb port address
var MysqlDBName = fs.String(MysqlDBNameFlag, "", "mysqlDB name")

func init() {
	fs.Parse(os.Args[1:])
}

//GetEnviromentVariables Function is used to fetch enviroment varriables
func GetEnviromentVariables() {

	//get mysqlDBHost from enviroment variables
	var mysqlDBHost = os.Getenv(MysqlDBHostEnvVar)
	if len(mysqlDBHost) > 0 && (MysqlDBHost == nil || len(*MysqlDBHost) == 0) {
		MysqlDBHost = &mysqlDBHost
	}

	//get mysqlDBUser from enviroments variables
	var mysqlDBUser = os.Getenv(MysqlDBUserEnvVar)
	if len(mysqlDBUser) > 0 && (MysqlDBUser == nil || len(*MysqlDBUser) == 0) {
		MysqlDBUser = &mysqlDBUser
	}

	//get mysqlDBPass from enviroments variables
	var mysqlDBPass = os.Getenv(MysqlDBPassEnvVar)
	if len(mysqlDBPass) > 0 && (MysqlDBPass == nil || len(*MysqlDBPass) == 0) {
		MysqlDBPass = &mysqlDBPass
	}

	//get mysqlDBAddr from enviroments variables
	var mysqlDBAddr = os.Getenv(MysqlDBAddrEnvVar)
	if len(mysqlDBAddr) > 0 && (MysqlDBAddr == nil || len(*MysqlDBAddr) == 0) {
		MysqlDBAddr = &mysqlDBAddr
	}

	//get httpAddr from enviroments variables
	var httpAddr = os.Getenv(HTTPPortEnvVar)
	if len(httpAddr) > 0 && (HTTPAddr == nil || len(*HTTPAddr) == 0) {
		HTTPAddr = &httpAddr
	}

	//get mysqlDBAddr from enviroments variables
	var mysqlDBName = os.Getenv(MysqlDBNameEnvVar)
	if len(mysqlDBName) > 0 && (MysqlDBName == nil || len(*MysqlDBName) == 0) {
		MysqlDBName = &mysqlDBName
	}

}

//ValidateFlags ckecks if flags are provided
func ValidateFlags() error {
	GetEnviromentVariables()

	flagMessage := " is a requird flag"
	if MysqlDBUser == nil || len(*MysqlDBUser) == 0 {
		return errors.New(MysqlDBUserFlag + flagMessage)
	}
	if MysqlDBPass == nil || len(*MysqlDBPass) == 0 {
		return errors.New(MysqlDBPassFlag + flagMessage)
	}

	if MysqlDBHost == nil || len(*MysqlDBHost) == 0 {
		return errors.New(MysqlDBHostFlag + flagMessage)
	}

	if MysqlDBName == nil || len(*MysqlDBName) == 0 {
		return errors.New(MysqlDBNameFlag + flagMessage)
	}

	return nil
}

//CreateConnectionString Function creates a connection string for database connection
func CreateConnectionString() string {

	var connectionStr string

	connectionStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/goxpert?charset=utf8&parseTime=True&loc=Local",
		*MysqlDBUser,
		*MysqlDBPass,
		*MysqlDBHost,
		*MysqlDBAddr,
	)
	fmt.Print(connectionStr)
	return connectionStr
}

//Run Function runs all the dependancies for the service
func Run() {

	if err := ValidateFlags(); err != nil {
		fmt.Println(err)
		return
	}

	router := mux.NewRouter()

	//MysqlDB Dependancy
	conPool := database.NewMysqlConnection(CreateConnectionString())

	//Course Services dependancies
	// adminRepository := repository.NewCourseRepositoryImpl(conPool)
	// adminService := service.NewCourseServiceImpl(adminRepository)
	// adminHandler := handler.NewCourseHandlerImpl(adminService)

	//Questions Services dependancies
	quationsRepository := repository.NewQuestionsRepositoryImpl(conPool)
	quationsService := service.NewQuestionsServiceImpl(quationsRepository)
	quationsHandler := handler.NewQuestionsHandlerImpl(quationsService)
	endpoints.NewQuestionsRoute(router, quationsHandler)

	//Sections Services dependancies
	sectionsRepository := repository.NewSectionsRepositoryImpl(conPool)
	sectionsService := service.NewSectionsServiceImpl(sectionsRepository)
	sectionsHandler := handler.NewSectionsHandlerImpl(sectionsService)
	endpoints.NewSectionsRoute(router, sectionsHandler)

	//answers Services dependancies
	answersRepository := repository.NewAnswersRepositoryImpl(conPool)
	userRepository := repository.NewUserServiceRepoImpl()
	answersService := service.NewAnswersServiceImpl(answersRepository, userRepository)
	answersHandler := handler.NewAnswersHandlerImpl(answersService)
	endpoints.NewAnswersRoute(router, answersHandler)

	//Use Middleware to log requests
	router.Use(middleware.LoggingMiddleware)
	//router.Use(AuthMiddleware)
	fmt.Println(http.ListenAndServe(":"+*HTTPAddr, router))
}
