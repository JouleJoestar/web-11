# Сервис "Hello"

В демонстрационном решении используется СУБД Postgresql с БД sandbox, в которой имеется единственная таблица, созданная с помощью простейшего sql-кода ниже:

```
create table hello
(
    message text not null
);
```