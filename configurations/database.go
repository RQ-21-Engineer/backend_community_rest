package configurations

import (

        "context"
        b64 "encoding/base64"

        "golang.org/x/oauth2/google"
        "google.golang.org/api/option"
        "google.golang.org/api/sheets/v4"

        exception "backend_community_rest/exceptions"

)


func SpreadsheetDatabase() (*sheets.Service) {

    // environment variables
    env := EnvironmentVariables()

    // create api context
    ctx := context.Background()

    // get bytes from base64 encoded google service accounts key
    credBytes, err := b64.StdEncoding.DecodeString(env.KeyJsonBase64)
    exception.TryCatchError(err)

    // authenticate and get configuration
    config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
    exception.TryCatchError(err)

    // create client with config and context
    client := config.Client(ctx)

    // create new service using client
    srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
    exception.TryCatchError(err)

    return srv

}