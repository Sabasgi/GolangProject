/*
SQLyog Ultimate v12.11 (64 bit)
MySQL - 8.0.40 : Database - labms
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`labms` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `labms`;

/*Table structure for table `branch` */

DROP TABLE IF EXISTS `branch`;

CREATE TABLE `branch` (
  `branch_id` int NOT NULL AUTO_INCREMENT,
  `branch_name` varchar(255) DEFAULT NULL,
  `lab_id` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `address` varchar(255) DEFAULT NULL,
  `branch_code` varchar(255) NOT NULL,
  `city_id` int DEFAULT NULL,
  PRIMARY KEY (`branch_id`),
  UNIQUE KEY `branch_code` (`branch_code`),
  KEY `lab_id` (`lab_id`),
  KEY `city_id` (`city_id`),
  CONSTRAINT `branch_ibfk_1` FOREIGN KEY (`lab_id`) REFERENCES `lab` (`lab_id`),
  CONSTRAINT `branch_ibfk_2` FOREIGN KEY (`city_id`) REFERENCES `city` (`city_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `branch` */

insert  into `branch`(`branch_id`,`branch_name`,`lab_id`,`created_at`,`address`,`branch_code`,`city_id`) values (1,'Branch ABC',1,'2024-11-07 14:14:31','dhayari pune','QG986062',1);

/*Table structure for table `city` */

DROP TABLE IF EXISTS `city`;

CREATE TABLE `city` (
  `city_id` int NOT NULL AUTO_INCREMENT,
  `city_name` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `state_id` int DEFAULT NULL,
  `created_on` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`city_id`),
  KEY `state_id` (`state_id`),
  CONSTRAINT `city_ibfk_1` FOREIGN KEY (`state_id`) REFERENCES `state` (`state_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `city` */

insert  into `city`(`city_id`,`city_name`,`created_at`,`state_id`,`created_on`,`created_by`) values (1,'Pune','2024-11-07 14:12:02',1,'2024-11-07 14:12:02',NULL);

/*Table structure for table `country` */

DROP TABLE IF EXISTS `country`;

CREATE TABLE `country` (
  `country_id` int NOT NULL AUTO_INCREMENT,
  `country_name` varchar(255) NOT NULL,
  `country_code` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `gst_percentage` int DEFAULT NULL,
  PRIMARY KEY (`country_id`),
  UNIQUE KEY `country_code` (`country_code`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `country` */

insert  into `country`(`country_id`,`country_name`,`country_code`,`created_at`,`gst_percentage`) values (1,'India','','2024-11-07 14:06:05',14),(2,'USA','DL373096','2024-11-07 14:08:50',14);

/*Table structure for table `department` */

DROP TABLE IF EXISTS `department`;

CREATE TABLE `department` (
  `department_id` int NOT NULL AUTO_INCREMENT,
  `department_name` varchar(255) NOT NULL,
  `lab_id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `branch_id` int NOT NULL,
  `description` text,
  PRIMARY KEY (`department_id`),
  KEY `lab_id` (`lab_id`),
  KEY `branch_id` (`branch_id`),
  CONSTRAINT `department_ibfk_1` FOREIGN KEY (`lab_id`) REFERENCES `lab` (`lab_id`),
  CONSTRAINT `department_ibfk_2` FOREIGN KEY (`branch_id`) REFERENCES `branch` (`branch_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `department` */

insert  into `department`(`department_id`,`department_name`,`lab_id`,`created_at`,`branch_id`,`description`) values (1,'Biochemistry',1,'2024-11-07 14:17:20',1,'description of deptaerment');

/*Table structure for table `lab` */

DROP TABLE IF EXISTS `lab`;

CREATE TABLE `lab` (
  `lab_id` int NOT NULL AUTO_INCREMENT,
  `lab_name` varchar(255) NOT NULL,
  `lab_Code` varchar(255) NOT NULL,
  `created_on` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`lab_id`),
  UNIQUE KEY `lab_Code` (`lab_Code`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `lab` */

insert  into `lab`(`lab_id`,`lab_name`,`lab_Code`,`created_on`,`created_by`) values (1,'ABC Lab','LAB001','2024-11-07 14:13:17',NULL),(3,'ABC Lab','LAB00000','2025-01-14 15:57:34',NULL);

/*Table structure for table `menu` */

DROP TABLE IF EXISTS `menu`;

CREATE TABLE `menu` (
  `menu_id` int NOT NULL AUTO_INCREMENT,
  `label` varchar(255) NOT NULL,
  `to_url` varchar(255) NOT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `parent_menu_id` int DEFAULT NULL,
  PRIMARY KEY (`menu_id`),
  KEY `fk_parent_menu` (`parent_menu_id`),
  CONSTRAINT `fk_parent_menu` FOREIGN KEY (`parent_menu_id`) REFERENCES `menu` (`menu_id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `menu` */

insert  into `menu`(`menu_id`,`label`,`to_url`,`icon`,`parent_menu_id`) values (1,'User','','pi pi-user',NULL),(2,'Create User','/user/create','pi pi-user',1),(3,'Update User','/user/edit','pi pi-user',1),(4,'Laboratory','','pi pi-user',NULL),(5,'Create Laboratory','/lab/create','pi pi-user',4),(6,'Update Laboratory','/lab/edit','pi pi-user',4),(7,'Lab Service','','pi pi-user',NULL);

/*Table structure for table `patient` */

DROP TABLE IF EXISTS `patient`;

CREATE TABLE `patient` (
  `patient_id` int NOT NULL AUTO_INCREMENT,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `age` int NOT NULL,
  `gender` varchar(10) NOT NULL,
  `contact_number` varchar(15) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `address` text,
  `patient_code` varchar(50) DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `medical_history` text,
  `blood_type` varchar(3) DEFAULT NULL,
  `insurance_details` text,
  `state_id` int DEFAULT NULL,
  `country_id` int DEFAULT NULL,
  `city_id` int DEFAULT NULL,
  PRIMARY KEY (`patient_id`),
  UNIQUE KEY `patient_code` (`patient_code`),
  KEY `state_id` (`state_id`),
  KEY `country_id` (`country_id`),
  KEY `city_id` (`city_id`),
  CONSTRAINT `patient_ibfk_1` FOREIGN KEY (`state_id`) REFERENCES `state` (`state_id`),
  CONSTRAINT `patient_ibfk_2` FOREIGN KEY (`country_id`) REFERENCES `country` (`country_id`),
  CONSTRAINT `patient_ibfk_3` FOREIGN KEY (`city_id`) REFERENCES `city` (`city_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `patient` */

insert  into `patient`(`patient_id`,`first_name`,`last_name`,`age`,`gender`,`contact_number`,`email`,`address`,`patient_code`,`user_id`,`medical_history`,`blood_type`,`insurance_details`,`state_id`,`country_id`,`city_id`) values (1,'John','Doe',45,'Male','123-456-7890','john.doe@example.com','123 Main St, Springfield','QS015676',0,'No significant medical history.','O+','XYZ Insurance Co, Policy #123456',1,1,1),(2,'John','Doe',45,'Male','123-456-7890','john.doe@example.com','123 Main St, Springfield','JS606380',0,'No significant medical history.','O+','XYZ Insurance Co, Policy #123456',1,1,1),(4,'','',0,'','','','','JX879846',0,'','','',1,1,1),(5,'','',0,'','','','','DW979110',0,'','','',1,1,1),(6,'shraddha','sabasgi',16,'male','34267891927','shraddhasa@mkcl.org','pune djhayari','MT040198',0,'','','',1,1,1),(11,'shraddha','sabasgi',3,'','','','','XB101549',0,'','','',1,1,1),(12,'shraddha','sabasgi',3,'','','','','AX952507',0,'','','',1,1,1),(13,'shraddha','sabasgi',3,'','','','','MW871450',0,'','','',1,1,1),(14,'shraddha','sabasgi',3,'','','','','HC803686',0,'','','',1,1,1),(15,'shraddha','sabasgi',3,'','','','','CR114315',0,'','','',1,1,1),(16,'shraddha','sabasgi',3,'','','','','II830697',0,'','','',1,1,1);

/*Table structure for table `permission` */

DROP TABLE IF EXISTS `permission`;

CREATE TABLE `permission` (
  `permission_id` int NOT NULL AUTO_INCREMENT,
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  `allowed` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`permission_id`),
  KEY `fk_menu` (`menu_id`),
  KEY `fk_role` (`role_id`),
  CONSTRAINT `fk_menu` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`menu_id`),
  CONSTRAINT `fk_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `permission` */

insert  into `permission`(`permission_id`,`role_id`,`menu_id`,`allowed`) values (1,1,1,1),(2,1,2,1),(3,1,3,1),(4,1,4,1);

/*Table structure for table `role` */

DROP TABLE IF EXISTS `role`;

CREATE TABLE `role` (
  `role_id` int NOT NULL AUTO_INCREMENT,
  `role_name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `role` */

insert  into `role`(`role_id`,`role_name`,`description`,`created_at`) values (1,'admin','Admin is new role ','2025-01-14 17:56:30');

/*Table structure for table `service` */

DROP TABLE IF EXISTS `service`;

CREATE TABLE `service` (
  `service_id` int NOT NULL AUTO_INCREMENT,
  `department_id` int NOT NULL,
  `service_name` varchar(255) NOT NULL,
  `description` text,
  `basic_rate` decimal(10,2) NOT NULL,
  `duration_minutes` int DEFAULT NULL,
  `preparation_instructions` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`service_id`),
  KEY `department_id` (`department_id`),
  CONSTRAINT `service_ibfk_1` FOREIGN KEY (`department_id`) REFERENCES `department` (`department_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `service` */

insert  into `service`(`service_id`,`department_id`,`service_name`,`description`,`basic_rate`,`duration_minutes`,`preparation_instructions`,`created_at`) values (1,1,'dhayari pune','description of service','78.90',10,'Instructions of tets','2024-11-07 14:18:40');

/*Table structure for table `state` */

DROP TABLE IF EXISTS `state`;

CREATE TABLE `state` (
  `state_id` int NOT NULL AUTO_INCREMENT,
  `state_name` varchar(255) NOT NULL,
  `country_id` int DEFAULT NULL,
  `created_on` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`state_id`),
  KEY `country_id` (`country_id`),
  CONSTRAINT `state_ibfk_1` FOREIGN KEY (`country_id`) REFERENCES `country` (`country_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `state` */

insert  into `state`(`state_id`,`state_name`,`country_id`,`created_on`,`created_by`) values (1,'Maharashtra',1,'2024-11-07 14:09:57',NULL),(2,'Goa',1,'2024-11-07 14:09:57',NULL),(3,'Maharashtra',1,'2024-11-07 14:10:52',NULL),(4,'Goa',1,'2024-11-07 14:10:52',NULL);

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `NAME` varchar(100) DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  `PASSWORD` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `role` enum('admin','doctor','nurse','receptionist','patient','user') NOT NULL,
  `phone_number` varchar(15) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `lab_id` int NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `username` (`username`),
  KEY `lab_id` (`lab_id`),
  CONSTRAINT `user_ibfk_1` FOREIGN KEY (`lab_id`) REFERENCES `lab` (`lab_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `user` */

insert  into `user`(`user_id`,`NAME`,`email`,`PASSWORD`,`username`,`role`,`phone_number`,`created_at`,`lab_id`) values (1,'John Doe','jdoe@example.com','password123','Username','admin','1234567890','2024-11-13 17:25:59',1),(4,'John Doe','jd9oe@example.com','password123','00sername','admin','1234567890','2024-11-13 17:27:52',1);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
