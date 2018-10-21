# gRPC_demo

> two simple demos of gRPC

## RUN Greeter 
start the server:

    go run greeter/hello_server/main.go
    
start the client:

    go run greeter/hello_client/main.go yourNname
    
## RUN Book 

Book is a  more complicated application.

start the server:

    go run book/server/main.go
    
start the client:

    go run book/client/main.go 
    
then in client you get like this:


    2018/10/21 17:52:44 Get BookInfo of Id: 1
    2018/10/21 17:52:44 name:"Fluent Python" author:"Luciano Ramalho" page:233

    2018/10/21 17:52:44 Get Book List of Info:
    2018/10/21 17:52:44
    2018/10/21 17:52:44 Post Book List of Info:
    2018/10/21 17:52:44 post summary: msg:"success!"

    2018/10/21 17:52:44 Get Book List of Info:
    2018/10/21 17:52:44 Book Info: name:"Fluent Python" author:"Luciano Ramalho" page:2333
    2018/10/21 17:52:44 Book Info: name:"1234" author:"ahahhahahhahhah" page:1234
    2018/10/21 17:52:44

