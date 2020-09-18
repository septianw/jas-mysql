CREATE DATABASE ipointtest;
CREATE USER 'asep'@'%' IDENTIFIED BY 'dummypass';
GRANT ALL PRIVILEGES ON ipointtest.* TO 'asep'@'%';
FLUSH PRIVILEGES;
