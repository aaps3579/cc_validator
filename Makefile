gen-api-spec:
	oapi-codegen -package card .\api\open-api-spec\open-api.yaml > .\api\card\card.gen.go
