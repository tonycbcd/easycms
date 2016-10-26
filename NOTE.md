# For Develop Note.

## The DB for MYSQL

CREATE DATABASE `easycms` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
GRANT ALL ON social.* to 'easycms'@'localhost' IDENTIFIED BY "easycms";
flush privileges;
