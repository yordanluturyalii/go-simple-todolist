create table todolists
(
    id   bigint auto_increment not null primary key,
    task varchar(255)          not null
) engine = InnoDB;

insert into todolists (TASK)
values ("task 1"),
       ("task 2"),
       ("task 3");