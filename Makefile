# deploy project
deploy_test:
	vercel .

deploy_prod:
	vercel .
	vercel --prod


# rest api functional testing
payload = ./testing/payload.json
test_command = curl -X POST -H "Content-Type: application/json" -d @


dev_endpoint = https://backend-community-rest-alfariqyraihan-qnetics.vercel.app/api/join
prod_endpoint = https://backend-community-rest.vercel.app/api/join


dev_test:
	$(test_command)$(payload) $(dev_endpoint)


prod_test:
	$(test_command)$(payload) $(prod_endpoint)












