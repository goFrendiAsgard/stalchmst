USE db;

DROP TABLE IF EXISTS `templates`;
CREATE TABLE IF NOT EXISTS `templates` (
  `code` varchar(50) NOT NULL COMMENT 'code',
  `template` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'Go template',
  `data` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'JSON data',
  PRIMARY KEY (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


INSERT INTO `templates` (`code`, `template`, `data`) VALUES
	('cv-markdown-example', '# Name\r\n{{.Name}}\r\n# Skills\r\n{{range $skill := .Skills}}* {{$skill}}\r\n{{end}}', '{\r\n  "Name": "Emiya Shirou", \r\n  "Skills": ["Unlimited Blade Works", "Rho Aias"]\r\n}'),
	('gherkin-js', 'describe("{{.Given}}", () => {\r\n    describe("{{.When}}", () => {\r\n        it("{{.Then}}", () => {\r\n            // TODO: Write your code\r\n        }\r\n    }\r\n});', '{\r\n    "Given": "Login page",\r\n    "When": "User provide valid username, password AND click login button",\r\n    "Then": "Session created"\r\n}'),
	('pantun-receh', 'Jalan jalan ke kota {{.Kota}}\r\nJangan lupa membeli {{.Barang}}\r\nBuat apa ke kota {{.Kota}}\r\nKalau cuma membeli {{.Barang}}', '{\r\n  "Kota": "Malang",\r\n  "Barang": "Jenang"\r\n}'),
	('remark', '<!DOCTYPE html>\r\n<html>\r\n  <head>\r\n    <title>Title</title>\r\n    <meta charset="utf-8">\r\n    <style>\r\n      @import url(https://fonts.googleapis.com/css?family=Yanone+Kaffeesatz);\r\n      @import url(https://fonts.googleapis.com/css?family=Droid+Serif:400,700,400italic);\r\n      @import url(https://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700,400italic);\r\n\r\n      body { font-family: \'Droid Serif\'; }\r\n      h1, h2, h3 {\r\n        font-family: \'Yanone Kaffeesatz\';\r\n        font-weight: normal;\r\n      }\r\n      .remark-code, .remark-inline-code { font-family: \'Ubuntu Mono\'; }\r\n    </style>\r\n  </head>\r\n  <body>\r\n    <textarea id="source">\r\n{{ .Content }}\r\n    </textarea>\r\n    <script src="https://remarkjs.com/downloads/remark-latest.min.js">\r\n    </script>\r\n    <script>\r\n      var slideshow = remark.create();\r\n    </script>\r\n  </body>\r\n</html>', '{\r\n    "Content": "class: center, middle\\n\\n# Title\\n\\n---\\n\\n# Agenda\\n\\n1. Introduction\\n2. Deep-dive\\n3. ..\\n\\n---\\n\\n# Introduction"\r\n}');
