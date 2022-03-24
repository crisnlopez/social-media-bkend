CREATE TABLE IF NOT EXISTS users (
  id int not null auto_increment,
  email varchar(60) not null,
  pass varchar(50) not null,
  user_nick varchar(10) not null,
  user_name varchar(10) not null,
  age int not null,
  PRIMARY KEY (id)
  )
