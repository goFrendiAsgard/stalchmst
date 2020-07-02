USE `db`;

DROP TABLE IF EXISTS `templates`;
CREATE TABLE IF NOT EXISTS `templates` (
  `code` varchar(50) NOT NULL COMMENT 'code',
  `template` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'Go template',
  `data` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'JSON data',
  PRIMARY KEY (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `templates` (`code`, `template`, `data`) VALUES
	('cv-markdown-example', '# Name\r\n{{.Name}}\r\n# Skills\r\n{{range $skill := .Skills}}* {{$skill}}\r\n{{end}}', '{"Name": "Emiya Shirou", "Skills": ["Unlimited Blade Works", "Rho Aias"]}');