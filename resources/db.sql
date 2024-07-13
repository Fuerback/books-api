CREATE TABLE `books` (
	`id` VARCHAR(36) NOT NULL,
	`title` TEXT NOT NULL,
	`author` TEXT NOT NULL,
	`deleted` BOOLEAN NOT NULL DEFAULT '0',
	`pages` INT NOT NULL,
	UNIQUE KEY `id_index` (`id`) USING BTREE,
	PRIMARY KEY (`id`)
);