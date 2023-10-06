# Web Shopping em Go 🛒

Bem-vindo ao Web Shopping em Go, o projeto que traz um toque de diversão às compras online! 😄
    
## Visão Geral

Este projeto é uma aplicação web simples para listar e gerenciar produtos em um e-commerce fictício. Aqui, você pode adicionar, editar e excluir produtos para simular uma experiência de compra online.

## Estrutura do Projeto
Claro, vou criar um resumo em Markdown para cada uma das pastas do projeto, explicando sobre a arquitetura MVC:

### `database`
- **Descrição**: Esta pasta contém os arquivos SQL necessários para configurar e preencher o banco de dados da aplicação.
- **Conteúdo**:
    - `init.sql`: Cria a estrutura inicial do banco de dados.
    - `populate.sql`: Insere dados iniciais no banco de dados.

### `docker-compose.yml`
- **Descrição**: Este arquivo é usado para configurar o ambiente de desenvolvimento usando Docker Compose. Ele ajuda a criar um ambiente isolado e consistente.

### `go.mod` e `go.sum`
- **Descrição**: Esses arquivos são usados para gerenciar as dependências do projeto em Go, garantindo que as bibliotecas externas sejam usadas nas versões corretas.

### `main.go`
- **Descrição**: Este é o arquivo principal da aplicação em Go. Ele é responsável por iniciar o servidor e configurar as rotas da aplicação.

### `src`
- **Descrição**: Esta pasta contém a maior parte do código da aplicação e segue a estrutura MVC.

#### `config`
- **Descrição**: Contém as configurações da aplicação, incluindo a configuração de conexão com o banco de dados.
- **Arquivo**:
    - `db.go`: Lida com a configuração do banco de dados.

#### `controllers`
- **Descrição**: Contém os controladores que gerenciam as solicitações HTTP, incluindo a lógica de negócios.
- **Arquivo**:
    - `products_controller.go`:  Contém funções para listar, criar, editar e excluir produtos.

#### `models`
- **Descrição**: Define os modelos de dados que representam a estrutura dos produtos.
- **Arquivo**:
    - `products.go`: Define a estrutura de dados para os produtos.

#### `routes`
- **Descrição**: Configura as rotas da aplicação, associando-as aos controladores apropriados.
- **Arquivo**:
    - `routes.go`: Define as rotas da aplicação.

### `templates`
- **Descrição**: Contém os arquivos HTML usados para renderizar a interface do usuário da aplicação.
- **Conteúdo**:
    - `_bootstrap.html`: Parte de um template Bootstrap.
    - `_bootstrapscripts.html`: Scripts relacionados ao Bootstrap.
    - `edit.html`: Página de edição de produtos.
    - `index.html`: Página principal de listagem de produtos.
    - `_nav.html`: Componente de navegação.
    - `new-product.html`: Página de criação de novos produtos.

Essa estrutura segue o padrão MVC, onde os modelos representam os dados, os controladores gerenciam a lógica da aplicação e as visualizações (templates HTML) são responsáveis pela renderização da interface do usuário. Isso ajuda a manter a organização do código e a separação de preocupações em um projeto web.
## Recursos

- Listagem de produtos com nome, descrição, preço e quantidade.
- Adição de novos produtos.
- Edição de produtos existentes.
- Exclusão de produtos que você não quer mais (com confirmação, é claro!).

## Como Começar

1. Clone este repositório.
2. Configure o ambiente de acordo com as instruções no arquivo `docker-compose.yml`.
3. Execute a aplicação com `go run main.go`.
4. Acesse a aplicação em [http://localhost:12345](http://localhost:12345).

## Contribua

Gostou da ideia? Quer adicionar mais recursos divertidos? Ou talvez melhorar a experiência de compra? Fique à vontade para contribuir! Aceitamos contribuições em todas as formas e tamanhos.

## Licença

Este projeto é distribuído sob a licença MIT. Sinta-se à vontade para brincar com ele o quanto quiser!

Divirta-se comprando online e explorando este projeto! 🛍️🎉