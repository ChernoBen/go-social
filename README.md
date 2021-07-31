# API !

O nome **mux** significa “**multiplexador** de solicitação HTTP“. Como o http. ServeMux padrão, o **mux**. Router combina solicitações recebidas em uma lista de rotas registradas e chama um manipulador para a rota que corresponde à URL ou a outras condições. 
**Go** é uma linguagem de código aberto que permite com que você construa códigos simples, confiáveis e eficiente. Criada pelo Google em 2007 por Robert Grisemer, Rob Pike e Ken Thompson, Go é uma linguagem compilada e estaticamente tipada.


**SOBRE PACOTES**
*Objetivo*: ``obter um bom isolamento entre as partes``
*Principais*:
``Pacotes relacionados a estrutura do projeto`` **- ex: (Main,Router,Controllers,Modelos,Repositorios(Manipulação do banco)**
*Auxiliares*:
``Pacotes relacionados a utilizades do projeto`` **- ex: (Config,Banco,Autenticação,MiddleWare,Segurança,Respostas)***
- Config: configurações de variaveis de ambiente
- Banco: Responsavel pela conexão com banco(obs: não deve manipular o banco)
- Autenticação: Login, criação de token e demais responsabilidades de autenticação
- MiddleWare: Uma camada entre a requisição e a resposta e será empregado em verificar se um usuario está autenticado
- Segurança: Responsavel por lidar com senhas como verificar senhas X hash no banco.
- Repostas: Padrões de respostas retornadas pelo projeto


------
# Desenvolvimento
- go mod init nome-projeto
- criação do inicializador da api main.go
