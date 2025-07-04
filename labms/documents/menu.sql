/*
SQLyog Ultimate v12.11 (64 bit)
MySQL - 8.0.40 
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `menu` (
	`menu_id` int (11),
	`label` varchar (765),
	`to_url` varchar (765),
	`icon` varchar (765),
	`parent_menu_id` int (11)
); 
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('1','Home','','',NULL);
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('2','Dashboard','/dashboard','pi pi-fw pi-home','1');
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('3','Accesses','','',NULL);
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('4','Grant Accesses','/accesses','pi pi-fw pi-eye','3');
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('5','User','','','4');
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('6','Create User','/user/create','pi pi-user','5');
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('7','Lab','','',NULL);
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('8','Create Lab','/labs','pi pi-user','7');
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('9','Create Lab Service','/labs/services','pi pi-user','7');
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('10','Patient','','',NULL);
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('11','Registration','/pat/register','pi pi-user','10');
insert into `menu` (`menu_id`, `label`, `to_url`, `icon`, `parent_menu_id`) values('12','Role Create','/role/create','pi pi-user','3');
