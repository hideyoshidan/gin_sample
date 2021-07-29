create table if not exists samples(
    id int auto_increment primary key,
    text varchar(255),
    created_at datetime default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
)