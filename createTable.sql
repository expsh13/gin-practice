create table if not exists todo (
  todo_id integer unsigned auto_increment primary key,
  title varchar(100) not null,
  created_at datetime
);