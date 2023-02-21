package configurations

import (

	"os"
	"strconv"
	"github.com/joho/godotenv"

	exception "backend_community_rest/exceptions"

)


/*
 *  Environment Variable Configuration
 */

type envStruct struct {

    // Database Environment
    KeyJsonBase64  string
    SpreadsheetId  int
    SpreadsheetKey string

}


func EnvironmentVariables() envStruct {

    // load environment variables
    err := godotenv.Load()
    exception.TryCatchError(err)


    // convert string to integer "SPREADSHEET_KEY" environment variable
    spreadsheet_id_conv, err := strconv.Atoi(
            os.Getenv("SPREADSHEET_ID"),
    )

    exception.TryCatchError(err)


    // environment variable into struct
    var environment = envStruct{

            // database environment
            KeyJsonBase64  : os.Getenv("KEY_JSON_BASE64"),
            SpreadsheetId  : spreadsheet_id_conv,
            SpreadsheetKey : os.Getenv("SPREADSHEET_KEY"),

    }


    return environment

}