package repositories


import (

    "time"
    "context"

    "google.golang.org/api/sheets/v4"

    model "backend_community_rest/models"
    exception "backend_community_rest/exceptions"
    config "backend_community_rest/configurations"

)



func AddJoinRequestRepo(RequestData model.JoinRequest) (bool) {


    // get environment variable
    env := config.EnvironmentVariables()


    /*
     * spreadsheet id and key
     * example :
     *   - https://docs.google.com/spreadsheets/d/<SPREADSHEETID>/edit#gid=<SHEETID>
     *
     */
    sheetId := env.SpreadsheetId
    spreadsheetKey := env.SpreadsheetKey


    // create api context
    ctx := context.Background()

    // database configuration
    srv := config.SpreadsheetDatabase()


    // convert sheet ID to sheet name.
    sheetid_conv, err := srv.Spreadsheets.Get(
        spreadsheetKey,

    ).Fields(
        "sheets(properties(sheetId,title))", 

    ).Do()

    exception.TryCatchError(err)

    if sheetid_conv.HTTPStatusCode != 200 {

        exception.TryCatchError(err)
        return false
    }


    sheetName := ""

    for _, v := range sheetid_conv.Sheets {

        prop := v.Properties

        if prop.SheetId == int64(sheetId) {

            sheetName = prop.Title
            break
        }

    }


    // append value to the sheet.
    row := &sheets.ValueRange {

        Values: [][]interface{}{
			
            {
                time.Now(),
                RequestData.UsernameGithub,
                "not accept",
            },
        },
    }


    // append data to sheet
    append, err := srv.Spreadsheets.Values.Append(
        spreadsheetKey,
        sheetName,
        row,

    ).ValueInputOption(
        "USER_ENTERED",

    ).InsertDataOption(
        "INSERT_ROWS",

    ).Context(ctx).Do()

    exception.TryCatchError(err)

    if append.HTTPStatusCode != 200 {

        exception.TryCatchError(err)
        return false
    }


    return true


}
