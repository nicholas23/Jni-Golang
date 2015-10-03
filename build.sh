#/bin/bash
rm Hello.class libHello.h libHello.so
go build -buildmode=c-shared -o libHello.so main.go
javac Hello.java
java -Djava.library.path=. Hello
