DROP TABLE IF EXISTS `authors`;
CREATE TABLE `authors` (
	`a_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`a_name` varchar(255) NOT NULL,
	PRIMARY KEY(`a_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `authors` VALUES (1, 'Amina'),(2, 'Carmen');

DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
	`p_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`p_title` varchar(255) NOT NULL,
	`p_author` bigint(20) unsigned NOT NULL,
	PRIMARY KEY(p_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `posts` VALUES (1, 'L\'Aventure',1),(2, 'Rien de nouveau',1),(3, 'Espace Vide',2);
