CREATE TABLE `entries` (
  `title` varchar(255) not null default '',
  `url` varchar(255) not null default ''
);

INSERT INTO `entries` (`title`, `url`) VALUES
('ECS(Fargate)で動かすコンテナにSSMからクレデンシャル情報を渡す', 'http://khigashigashi.hatenablog.com/entry/2018/08/28/214417');