/* Please check that the values are the same as those described in the .env file */

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `credens_mysql` /* $MYSQL_DATABASE */ /*!40100 DEFAULT CHARACTER SET utf8_general_ci */;

USE `credens_mysql`; /* $MYSQL_DATABASE */

DROP USER 'credens'@'%'; /* $MYSQL_USER */
CREATE USER 'credens'@'%' /* $MYSQL_USER */ IDENTIFIED BY 'secret'; /* $MYSQL_PASSWORD */
GRANT ALL PRIVILEGES ON credens_mysql.* /* MYSQL_DATABASE */ TO 'credens'@'%' WITH GRANT OPTION; /* $MYSQL_USER */
