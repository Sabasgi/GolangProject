/*
SQLyog Ultimate v12.11 (64 bit)
MySQL - 8.0.40
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `permission` (
	`permission_id` int (11),
	`role_id` int (11),
	`menu_id` int (11),
	`allowed` tinyint (1)
);
insert into `permission`( `role_id`, `menu_id`, `allowed`) values('2','1','1');
insert into `permission` ( `role_id`, `menu_id`, `allowed`) values('2','2','1');
insert into `permission` ( `role_id`, `menu_id`, `allowed`) values('2','3','1');
insert into `permission` ( `role_id`, `menu_id`, `allowed`) values('2','4','1');
insert into `permission` ( `role_id`, `menu_id`, `allowed`) values('2','5','1');
insert into `permission` ( `role_id`, `menu_id`, `allowed`) values('2','6','1');
insert into `permission` ( `role_id`, `menu_id`, `allowed`) values('2','7','1');
insert into `permission` ( `role_id`, `menu_id`, `allowed`) values('2','8','1');
insert into `permission` ( `role_id`, `menu_id`, `allowed`) values('2','9','1');
insert into `permission` ( `role_id`, `menu_id`, `allowed`) values('2','10','1');
insert into `permission` ( `role_id`, `menu_id`, `allowed`) values('2','11','1');
insert into `permission` ( `role_id`, `menu_id`, `allowed`) values('2','12','1');
