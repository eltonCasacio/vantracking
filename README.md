# API do projeto de monitoramento de veículos em tempo real
![Badge](https://img.shields.io/static/v1?label=go&message=1.20&color=blue&style=for-the-badge&logo=Go)
![Badge](https://img.shields.io/static/v1?label=push%20notification&message=%20&color=gray&style=for-the-badge&logo=firebase)
![Badge](https://img.shields.io/static/v1?label=aws&message=%20&color=yellow&style=for-the-badge&logo=aws)

## Outros repositorios que fazem parte desse
- https://github.com/eltonCasacio/vanmonitor
- https://github.com/eltonCasacio/vandriver

## Sobre o projeto de monitoramento
Este sistema foi desenvolvido para o monitoramento de vans escolares.
Há dois atores nesse projeto, Motorista e Monitor.

***O Motorista:***
- cria rotas
- create routes
- aprova passageiros
- envia a localização em tempo real para a API
- faz push notification para o monitor quando esta chegando ao ponto de embarque do passageiro
- visualiza se o passageiro vai e volta com a van
- visualiza se o passageiro ja embarcou

***O Monitor:***
- cadastra passageiros
- visualiza no mapa a localização do passageiro em tempo real
- informa ao motorista se o passageiro vai e volta de van
- recebe notificação quando motorista está chegando

***Some used services in this project***
- AWS to deploy API
- Firebase for push notification
- Google maps api


# Para Desenvolvedores
Para executar esse projeto será necessário criar um arquivo .env e setar seus valores. Há um arquivo de exemplo chamado .env_prod dentro da pasta cmd/server.

E também as dependencias devem ser instaladas.

### So, let`s go
1. clone esse repositorio em sua maquina
2. execute `docker-compose up`
3. execute `make migrate` para criar o banco de dados e tabelas
4. entre na pasta cmd/server e execute `go run main.go`
   
se todas as dependencias estiverem instaladas corretamente a api estará pronta para servir os clientes.

Podemos testar a api atraves dos arquivos que se encontram dentro da pasta test, lá encontramos arquivos com extensão .http <br/>
com esses arquivos conseguimos testar os endpoints. Ou se preferir pode ser usado o postman por exemplo.

Pode ser que tenha que alterar o host nas urls.
