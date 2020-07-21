-- MySQL dump 10.13  Distrib 5.7.19, for macos10.12 (x86_64)
--
-- Host: localhost    Database: admin
-- ------------------------------------------------------
-- Server version	5.7.19

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `title` varchar(100) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '2',
  `sort` int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group`
--

LOCK TABLES `group` WRITE;
/*!40000 ALTER TABLE `group` DISABLE KEYS */;
INSERT INTO `group` VALUES (1,'APP','System',2,1);
/*!40000 ALTER TABLE `group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `node`
--

DROP TABLE IF EXISTS `node`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `node` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL DEFAULT '',
  `name` varchar(100) NOT NULL DEFAULT '',
  `level` int(11) NOT NULL DEFAULT '1',
  `pid` bigint(20) NOT NULL DEFAULT '0',
  `remark` varchar(200) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '2',
  `group_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `node`
--

LOCK TABLES `node` WRITE;
/*!40000 ALTER TABLE `node` DISABLE KEYS */;
INSERT INTO `node` VALUES (1,'RBAC','rbac',1,0,'',2,1),(2,'Node','node/index',2,1,'',2,1),(3,'node list','index',3,2,'',2,1),(4,'add or edit','AddAndEdit',3,2,'',2,1),(5,'del node','DelNode',3,2,'',2,1),(6,'User','user/index',2,1,'',2,1),(7,'user list','Index',3,6,'',2,1),(8,'add user','AddUser',3,6,'',2,1),(9,'update user','UpdateUser',3,6,'',2,1),(10,'del user','DelUser',3,6,'',2,1),(11,'Group','group/index',2,1,'',2,1),(12,'group list','index',3,11,'',2,1),(13,'add group','AddGroup',3,11,'',2,1),(14,'update group','UpdateGroup',3,11,'',2,1),(15,'del group','DelGroup',3,11,'',2,1),(16,'Role','role/index',2,1,'',2,1),(17,'role list','index',3,16,'',2,1),(18,'add or edit','AddAndEdit',3,16,'',2,1),(19,'del role','DelRole',3,16,'',2,1),(20,'get roles','Getlist',3,16,'',2,1),(21,'show access','AccessToNode',3,16,'',2,1),(22,'add accsee','AddAccess',3,16,'',2,1),(23,'show role to userlist','RoleToUserList',3,16,'',2,1),(24,'add role to user','AddRoleToUser',3,16,'',2,1),(25,'SecKill','seckill',1,0,'',2,1),(26,'Product','product/index',2,25,'',2,1),(27,'product list','index',3,26,'',2,1),(29,'add product','AddProduct',3,26,'',2,1),(30,'update product','UpdateProduct',3,26,'',2,1),(31,'del product','DelProduct',3,26,'',2,1),(32,'Activity','activity/index',2,25,'',2,1),(33,'activity list','index',3,32,'',2,1),(34,'add activity','AddActivity',3,32,'',2,1),(35,'update activity','UpdateActivity',3,32,'',2,1),(36,'del activity','DelActivity',3,32,'',2,1);
/*!40000 ALTER TABLE `node` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `node_roles`
--

DROP TABLE IF EXISTS `node_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `node_roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `node_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `node_roles`
--

LOCK TABLES `node_roles` WRITE;
/*!40000 ALTER TABLE `node_roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `node_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL DEFAULT '',
  `name` varchar(100) NOT NULL DEFAULT '',
  `remark` varchar(200) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '2',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'Admin role','Admin','I\'m a admin role',2);
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sec_kill_activity`
--

DROP TABLE IF EXISTS `sec_kill_activity`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sec_kill_activity` (
  `activity_id` int(11) NOT NULL AUTO_INCREMENT,
  `activity_name` varchar(60) NOT NULL DEFAULT '',
  `product_id` int(11) NOT NULL DEFAULT '0',
  `start_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `total` int(11) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '2',
  `second_limit` int(11) NOT NULL DEFAULT '0',
  `everyone_limit` int(11) NOT NULL DEFAULT '0',
  `buy_rate` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`activity_id`),
  UNIQUE KEY `activity_name` (`activity_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sec_kill_activity`
--

LOCK TABLES `sec_kill_activity` WRITE;
/*!40000 ALTER TABLE `sec_kill_activity` DISABLE KEYS */;
INSERT INTO `sec_kill_activity` VALUES (1,'iPhone X秒杀',1,'2018-05-16 10:25:00','2018-05-20 16:00:00',100,3,5,1,50);
/*!40000 ALTER TABLE `sec_kill_activity` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sec_kill_lucky_user`
--

DROP TABLE IF EXISTS `sec_kill_lucky_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sec_kill_lucky_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `activity_id` int(11) NOT NULL DEFAULT '0',
  `user_id` int(11) NOT NULL DEFAULT '0',
  `token` varchar(60) NOT NULL DEFAULT '',
  `add_time` datetime NOT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `token` (`token`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sec_kill_lucky_user`
--

LOCK TABLES `sec_kill_lucky_user` WRITE;
/*!40000 ALTER TABLE `sec_kill_lucky_user` DISABLE KEYS */;
INSERT INTO `sec_kill_lucky_user` VALUES (1,1,1,'180239c501f38eb631b97966d2cf1255','2018-05-20 11:17:00',1),(2,1,1,'048bdbfac17b237f74a04f61f2e93b3a','2018-05-20 11:17:03',1),(3,1,1,'3cb3de1c1cef0e6a28f7758e850f8dc8','2018-05-20 11:21:38',1),(4,1,1,'4c7acc6ae4922c6e267a01aa514024ea','2018-05-20 11:27:48',1),(5,1,1,'07bae0ac7bcc02ac42a2462ce42b159b','2018-05-20 11:34:51',1),(6,1,1,'75bec81e726d84ca8f872866340adfc2','2018-05-20 11:38:33',1),(7,1,1,'4c69ca360eddd4af0443b5cac87792ab','2018-05-20 11:43:25',1),(8,1,1,'a4a12f1f98da098b2da1fb254c62beaf','2018-05-20 11:45:30',1),(9,1,1,'e30a7bf299e2a13dc86da00e214bf960','2018-05-20 11:47:26',1),(10,1,1,'dbf399f8b13c62df7253507427fcf826','2018-05-20 11:54:17',1),(11,1,1,'3014569eea107d164b5e1892082d2105','2018-05-20 11:57:52',1),(12,1,1,'5283b0752952ca366eb5ddf35bf3f417','2018-05-20 11:59:06',1),(13,1,1,'f09cfdc8758b09122c1a8e2d396fa148','2018-05-20 12:02:28',1),(14,1,1,'5153d7fa0facbd9f0b53e6251638921f','2018-05-20 12:03:49',1),(15,1,1,'4615514b1e504a540a4a2765db25c63c','2018-05-20 13:27:18',1);
/*!40000 ALTER TABLE `sec_kill_lucky_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sec_kill_product`
--

DROP TABLE IF EXISTS `sec_kill_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sec_kill_product` (
  `product_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `product_name` varchar(60) NOT NULL DEFAULT '''',
  `total` int(11) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '2',
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sec_kill_product`
--

LOCK TABLES `sec_kill_product` WRITE;
/*!40000 ALTER TABLE `sec_kill_product` DISABLE KEYS */;
INSERT INTO `sec_kill_product` VALUES (1,'iPhone X',100,2),(2,'iPhone 8 Plus',200,2);
/*!40000 ALTER TABLE `sec_kill_product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sec_kill_user`
--

DROP TABLE IF EXISTS `sec_kill_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sec_kill_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) NOT NULL DEFAULT '',
  `user_pwd` varchar(60) NOT NULL DEFAULT '',
  `user_email` varchar(60) NOT NULL DEFAULT '',
  `user_mobile` bigint(20) unsigned NOT NULL DEFAULT '0',
  `status` int(10) unsigned NOT NULL DEFAULT '1',
  `is_login` int(10) unsigned NOT NULL DEFAULT '1',
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_name` (`user_name`),
  UNIQUE KEY `user_pwd` (`user_pwd`),
  UNIQUE KEY `user_email` (`user_email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sec_kill_user`
--

LOCK TABLES `sec_kill_user` WRITE;
/*!40000 ALTER TABLE `sec_kill_user` DISABLE KEYS */;
INSERT INTO `sec_kill_user` VALUES (1,'user','29fe7d4261683524a453faea8abe2295','user@qq.com',18088889999,1,1,'2018-05-15 07:26:45');
/*!40000 ALTER TABLE `sec_kill_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `nickname` varchar(32) NOT NULL DEFAULT '',
  `email` varchar(32) NOT NULL DEFAULT '',
  `remark` varchar(200) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '2',
  `lastlogintime` datetime DEFAULT NULL,
  `createtime` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `nickname` (`nickname`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','21232f297a57a5a743894a0e4a801fc3','ClownFish','osgochina@gmail.com','I\'m admin',2,NULL,'2018-05-08 05:33:04');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_roles`
--

DROP TABLE IF EXISTS `user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_roles`
--

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-05-21 16:04:45
