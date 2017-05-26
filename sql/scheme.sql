drop database elis_test;
create database elis_test;
use elis_test;

CREATE TABLE admin (
  login VARCHAR(10),
  password INT,

  PRIMARY KEY (login)
);

CREATE TABLE image (
  id INT NOT NULL AUTO_INCREMENT,
  url VARCHAR(200) NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE portfolio_image (
  id INT NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE post (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(50),
  content TEXT,
  permalink VARCHAR(100) NOT NULL,
  created_at DATE NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE tag (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(40) NOT NULL UNIQUE,

  PRIMARY KEY (id)
);

CREATE TABLE post_tag (
  post_id INT NOT NULL,
  tag_id INT NOT NULL,

  PRIMARY KEY (post_id, tag_id)
);

CREATE TABLE liker (
  id INT NOT NULL AUTO_INCREMENT,
  post_id INT NOT NULL,
  ip VARCHAR(15),

  PRIMARY KEY (id)
);

CREATE TABLE message (
  id INT NOT NULL AUTO_INCREMENT,
  phone VARCHAR(20),
  email VARCHAR(30),
  content TEXT,
  name VARCHAR(20),
  theme VARCHAR(20),
  created_at DATE,

  PRIMARY KEY (id)
);

ALTER TABLE post
ADD CONSTRAINT un_post UNIQUE (permalink);

ALTER TABLE admin
ADD CONSTRAINT un_admin UNIQUE (login); 

ALTER TABLE post_tag 
ADD CONSTRAINT fk_post FOREIGN KEY (post_id) REFERENCES post (id);

ALTER TABLE post_tag
ADD CONSTRAINT fk_tag FOREIGN KEY (tag_id) REFERENCES tag (id);

ALTER TABLE liker
ADD CONSTRAINT fk_liker FOREIGN KEY (post_id) REFERENCES post (id);
