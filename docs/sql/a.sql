create TABLE blog_tag(
	id int(10) UNSIGNED not null AUTO_INCREMENT,
	name VARCHAR(100) DEFAULT '' COMMENT '标签名字',
	created_on int(10) UNSIGNED DEFAULT 0 COMMENT '创建时间',
	create_by VARCHAR(100) DEFAULT '' COMMENT ' 创建人',
	modified_on int(10) unsigned DEFAULT '0' COMMENT '修改时间',
	modified_by VARCHAR(100) DEFAULT '' comment '修改人',
	state TINYINT(3) UNSIGNED DEFAULT '1' comment '状态 0为禁用、1为启用',
	PRIMARY KEY(id)
)ENGINE=INNODB DEFAULT CHARSET='utf8' comment='文章标签管理';

