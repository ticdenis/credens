CREATE DATABASE /*!32312 IF NOT EXISTS*/ `credens_mysql` /* $MYSQL_DATABASE */ /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `credens_mysql`; /* $MYSQL_DATABASE */

DROP USER 'credens'@'%'; /* $MYSQL_USER */
CREATE USER 'credens'@'%' /* $MYSQL_USER */ IDENTIFIED BY 'secret'; /* $MYSQL_PASSWORD */
GRANT ALL PRIVILEGES ON credens_mysql.* /* MYSQL_DATABASE */ TO 'credens'@'%' WITH GRANT OPTION; /* $MYSQL_USER */
