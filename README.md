Link postman: https://web.postman.co/workspace/My-Workspace~e0b341c2-89ea-4314-8c58-868567ce6167/documentation/35041036-d6fd8c95-8d98-4057-8d18-4dc2a74cbf8f (invite indisponível por agora)

# Trabalho final LTP2 - Catálogo

Este é um sistema de catálogo que tem como intuito informar o usuário de todos os filmes que ele deveria assistir para conseguir acompanhar a história inteira de franquias de super heróis.
Como exemplo inicial, franquias marvel serão usadas.
Para entender a história inteira de [filme] o usuário deverá assistir: [filmes anteriores]

Além disso, ele também poderá avaliar e salvar seus filmes em listas com denominações personalizadas.
> Lista 1: Filmes que quero assistir

> Lista 2: Filmes favoritos

## Url base
`http://www.catalogo.com`

## Autenticação
[placeholder]

## Estrutura de dados
Segue um exemplo de como será enviado os dados do filme por JSON.


```JSON
{
"id": "string",
"nome": "string",
"descricao": "string",
"numero": "number",
"categoria": "string"
}
```

Para códigos de status:


```JSON
{
"error": "Requisito inválido.",
"message": "Motivo específicado."
}
```

**401** Unauthorized
**404** Not found
**500** Internal Server Error
**511** Network Authentication Required

## Endpoints
Seguem as endpoints de gerenciamento:

**Home com todos os filmes**
> GET/catalogo

Retorna na página inicial todos os filmes que estão no banco de dados no site.

**Receber filme pesquisado**
> POST/catalogo?filme

Recebe, em uma pequena aba de busca, o filme que o usuário que está procurando no banco de dados.

**Mostrar filmes que o antecedem**
> GET/catalogo?=f 

Ao encontrar o filme, o sistema irá mostrar todos os filmes que devem ser assistidos antes desse para acompanhar a narrativa.

**Criar listas**
> POST/catalogo/listas 

Função que cria listas de filmes.

**Colocar filme na lista**
> PUT/catalogo/listas

Função que adiciona os filmes na lista criada pelo usuário.

**Deletar filme da lista**
> DELETE/catalogo/listas 

Função que deleta o filme da lista desejada.

**Mandar avaliação**
> POST/catalogo/f# 

Função que envia ao sistema a avaliação pessoal (em estrelas) de um usuário.

**Mostrar avaliação do usuário**
> GET/catalogo/f# 

Função que irá devolver ao usuário sua avaliação pessoal, sempre que o mesmo abrir a página do filme.

**Atualizar avaliação**
> PATCH/catalogo/f# 

Função que irá modificar a avaliação do usuário, se desejado pelo mesmo.

**Deletar avaliação**
> DELETE/catalogo/f# 

Função que irá deletar um filme de uma lista desejada.

