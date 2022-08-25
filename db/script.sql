create database milhasdb

create table milhas(
    id_milhas SERIAL primary key,
    id_cartao varchar,
    cpf varchar,
    valor_compra float
)