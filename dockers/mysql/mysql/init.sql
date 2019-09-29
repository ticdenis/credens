CREATE DATABASE /*!32312 IF NOT EXISTS*/ `credens_mysql` /*!40100 DEFAULT CHARACTER SET utf8_general_ci */;

USE `credens_mysql`;

DROP USER 'credens'@'%';
CREATE USER 'credens'@'%' IDENTIFIED BY 'secret';
GRANT ALL PRIVILEGES ON credens_mysql.* TO 'credens'@'%' WITH GRANT OPTION;
