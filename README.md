# Teste Brand Monitor

__Objetivo__:  
Uma marca (a sua escolha) precisa identificar quais concorrentes estão utilizando seus termos de marca em anúncios patrocinados do Google. Para  isso, é necessário desenvolver uma aplicação web com uma interface simples (se responsiva melhor ainda) onde o usuário possa:

1. Cadastrar um ou mais termos de marca.
2. Fornecer um e-mail corporativo.
3. Clicar em "Descobrir".

Ao finalizar, uma mensagem de agradecimento deve ser exibida, informando que o diagnóstico será processado e enviado por e-mail. Dentro de um prazo máximo de 5 minutos, o usuário deve receber uma lista com os domínios dos principais concorrentes identificados.

__Requisitos__:
1. Consulta ao Google via API: A aplicação deve realizar consultas ao Google usando uma API.
2. Backend em Golang: O backend deve ser desenvolvido utilizando a linguagem Golang.
3. Banco de Dados MongoDB: Utilizar MongoDB para armazenar os resultados das buscas.
4. Frontend: O frontend pode ser desenvolvido com ou sem frameworks, conforme sua preferência.
5. Design do Frontend: Não se preocupe com a aparência do frontend, um visual básico é suficiente, porém responsivo ganhará pontos
6. Deploy: A aplicação deve estar hospedada em um servidor acessível e o código deve estar disponível em um repositório público no GitHub.
7. Apresentação: Após a entrega do repositório, haverá uma sessão de revisão (Sprint Review) para demonstrar a funcionalidade entregue e em funcionamento.

__Prazo__: __17 de julho de 2024__.

__Desafio Extra__ (_Opcional_):
Para um desafio adicional, considere implementar as seguintes funcionalidades:

1. Busca Simulada em Múltiplas Cidades: Agende buscas no Google em 10 cidades diferentes e agrupe os resultados por cidade no e-mail enviado.
2. Simulação de Diferentes Dispositivos: Simule buscas em diferentes tipos de dispositivos.
3. Processamento em Filas: Avalie a necessidade de utilizar filas de processamento. Justifique sua escolha explicando se você usaria ou não e o porquê.
4. Testes unitários e Clean code

## Roadmap

__TODOs__:

_Backend_

- [x] Configurar GIN framework
- [x] Criar package `utils`
- [x] Criar package `types`
- [x] Criar packages `services`
- [x] Criar rota de teste `/healthz` [GET]
- [x] Criar `Makefile`
- [x] Criar `dockerfile`

_Frontend_

- [x] Configurar React framework
- [x] Instalar dependências
- [x] Adicionar componentes
- [x] Adicionar páginas
- [x] Criar `dockerfile`

__Database__

- [x] Integrar banco de dados (mongodb)