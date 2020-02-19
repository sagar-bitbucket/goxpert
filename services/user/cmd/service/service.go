package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"gitlab.com/scalent/goxpert/email"

	"gitlab.com/scalent/goxpert/services/user/pkg/v1/endpoints"
	"gitlab.com/scalent/goxpert/services/user/pkg/v1/middleware"

	handler "gitlab.com/scalent/goxpert/services/user/pkg/v1/handlers"

	service "gitlab.com/scalent/goxpert/services/user/pkg/v1/services"

	repository "gitlab.com/scalent/goxpert/services/user/pkg/v1/repositories"

	"gitlab.com/scalent/goxpert/database"

	"net/http"

	"github.com/gorilla/mux"
)

var (
	//HTTPPortFlag cli flag name for http port
	HTTPPortFlag = "http-port"

	//MysqlDBHostFlag cli flag name for mysqldb host
	MysqlDBHostFlag = "mysqldb-host"

	//MysqlDBUserFlag cli flag name for mysqldb username
	MysqlDBUserFlag = "mysqldb-user"

	//MysqlDBPassFlag cli flag name for mysqldb password
	MysqlDBPassFlag = "mysqldb-pass"

	//MysqlDBAddrFlag cli flag name for mysqldb port
	MysqlDBAddrFlag = "mysqldb-addr"

	//DebugPortFlag cli flag name for debug address
	DebugPortFlag = "debug-addr"

	//MysqlDBNameFlag cli flag name for database name
	MysqlDBNameFlag = "mysqldb-name"

	EmailSecretFlag = "email-secret"

	//HTTPPortEnvVar **
	HTTPPortEnvVar = "HTTP_PORT"

	//MysqlDBHostEnvVar **
	MysqlDBHostEnvVar = "MYSQL_DB_HOST"

	//MysqlDBUserEnvVar **
	MysqlDBUserEnvVar = "MYSQL_DB_USER"

	//MysqlDBPassEnvVar **
	MysqlDBPassEnvVar = "MYSQL_DB_PASS"

	//MysqlDBAddrEnvVar **
	MysqlDBAddrEnvVar = "MYSQL_DB_PORT"

	//MysqlDBNameEnvVar **
	MysqlDBNameEnvVar = "MYSQL_DB_NAME"

	//EmailSecretEnvVar
	EmailSecretEnvVar = "EMAIL_SECRET"
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
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

var EmailSecret = fs.String(EmailSecretFlag, "", "email secret")

func init() {
	flag.Parse()
}

//GetEnviromentVariables **
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

	//get mysqlDBAddr from enviroments variables
	var emailSecret = os.Getenv(EmailSecretEnvVar)
	if len(emailSecret) > 0 && (EmailSecret == nil || len(*EmailSecret) == 0) {
		EmailSecret = &emailSecret
	}
}

//ValidateFlags ckecks the flags and update
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

	if EmailSecret == nil || len(*EmailSecret) == 0 {
		return errors.New(EmailSecretFlag + flagMessage)
	}

	return nil
}

//CreateConnectionString cc
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

//sql.Open("mysql", "root:password@/test?multiStatements=true")
//Run **

//CreateMigrationConnectionString cc
func CreateMigrationConnectionString() string {

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
func Run() {

	if err := ValidateFlags(); err != nil {
		fmt.Println(err)
		return
	}

	router := mux.NewRouter()

	//MysqlDB Dependancy
	conPool := database.NewMysqlConnection(CreateConnectionString())

	//Email Dependancy
	email := email.NewSendInBlue(*EmailSecret)

	//User Services dependancies
	userRepository := repository.NewUserRepositoryImpl(conPool)
	sectionServiceRepo := repository.NewSectionServiceRepoImpl()
	userService := service.NewUserServiceImpl(userRepository, email, sectionServiceRepo)
	userHandler := handler.NewUserHandlerImpl(userService)

	endpoints.NewUserRoute(router, userHandler)

	router.Use(middleware.LoggingMiddleware)
	//router.Use(AuthMiddleware)
	fmt.Println(http.ListenAndServe(":"+*HTTPAddr, router))
}
