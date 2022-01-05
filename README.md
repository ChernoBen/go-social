# API !

Exemplo simples de API de usuários utilizando Gorilla Mux e MySql.
Status: ainda em desenvolvimento.

# Schema:


![Alt text](api/doc/schema.png?raw=true "Title")

# **SOBRE PACOTES**
    - *Objetivo*: 
    ``Obter um bom isolamento entre as partes``
    - *Principais*:
    ``Pacotes relacionados a estrutura do projeto`` 
        - **- ex: (Main, Router, Controllers, Modelos, Repositorios( Manipulação do banco )**
    - *Auxiliares*:
``Pacotes relacionados a utilizades do projeto`` 
   # **- ex: (Config, Banco, Autenticação, MiddleWare, Segurança, Respostas)**

        - Config: configurações de variaveis de ambiente
        - Banco: Responsavel pela conexão com banco(obs: não deve manipular o banco)
        - Autenticação: Login, criação de token e demais responsabilidades de autenticação
        - MiddleWare: Uma camada entre a requisição e a resposta e será empregado em verificar se um usuario está autenticado
        - Segurança: Responsavel por lidar com senhas como verificar senhas X hash no banco.
        - Repostas: Padrões de respostas retornadas pelo projeto

------