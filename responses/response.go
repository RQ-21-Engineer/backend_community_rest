package responses


func AddJoinResponse(http_code_message int) (int, string, string) {


	// default value (if http response code 200)
	status        := "success"
	message       := "Uraa... Permintaan bergabung telah dikirimkan."


	if http_code_message != 200 {

		status        = "failed"
		message       = "yahh... Permintaan bergabung tidak dapat terkirim."

	}

	return http_code_message, status, message

	
}