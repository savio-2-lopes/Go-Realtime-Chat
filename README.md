<h3 align="center">
    <img alt="Chat" title="Chat" src="./assets/1.gif" width="800px" />
</h3>
 
<p align="center"> :computer: <strong>Em progresso ...</strong> 🚧</p>

<p align="center"> 
   <img src="https://img.shields.io/badge/version-0.0.1-yellow.svg" />
  
  <a href="https://github.com/savio-2-lopes">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" target="_blank" />
  </a>
 
 <a href="https://github.com/savio-2-lopes">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-blue.svg" target="_blank" />
  </a>
</p>

<br>

## :pushpin: Índice

- [Sobre](#sobre-o-projeto)
- [Progresso](#progresso)
- [Como executar](#executar)
- [Tecnologias](#tecnologia)
- [Licença](#licenca)

<br>

<a id="sobre-o-projeto"></a>

## 💻 Sobre o projeto

:rocket: Um simples Chat Real time utilizado WebSocket, GoLag e React.js

<br>

<a id="executar"></a>

## 🚀 Como executar o projeto

### Pré-requisitos

Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:
[Git](https://git-scm.com), [Node.js](https://nodejs.org/en/), [GO](https://golang.org/) e o gerenciador de pacotes [Yarn](https://yarnpkg.com/).
Além disto é bom ter um editor para trabalhar com o código, como [VSCode](https://code.visualstudio.com/)
Para seu devido fucionamento, o Backend e o fronted devem está rodando.

<br>

#### 🧭 Rodando o backend

```bash

# Primeiramente, clone este repositório
$ git clone https://github.com/savio-2-lopes/api_node.js.git

# Entre na pasta backend
$ cd api_node/backend

# Instale a biblioteca do websocket
$ go get -u github.com/gorilla/websocket

# Rode os seguintes comandos
$ go build main.go && go run main.go

# O backend estará rodando em
$ http://localhost:3333

```

<br>

#### 🧭 Rodando o Frontend

```bash

# Entre na pasta frontend
$ cd api_node/frontend

# Instale as depedências
$ yarn

# Rode o comando
$ yarn start

```

<br>

<a id="tecnologia"></a>

## 🛠 Tecnologias

As seguintes ferramentas foram usadas na construção do projeto:

- [React.js](https://reactjs.org)
- [Go Lang](https://golang.org/)
- [websocket](github.com/gorilla/websocket)

<br>

<a id="licenca"></a>

## :memo: Licença

Este projeto está sob a licença do MIT. Veja a [página de licença](https://opensource.org/licenses/MIT) para mais detalhes.
