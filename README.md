# go-hcaptcha-solver
This sample export api to solve hcaptcha

# How run
go run cmd/api.go

# How request api
curl --request POST \
  --url http://localhost:8080/token \
  --header 'Content-Type: application/json' \
  --data '{
	"url": "https://servicos.receita.fazenda.gov.br/servicos/cpf/consultasituacao/ConsultaPublica.asp",
	"api_key": "af4fc5a3-1ac5-4e6d-819d-324d412a5e9d"
}'