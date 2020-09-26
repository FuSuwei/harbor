create TABLE harbor_article(
    id int(10) UNSIGNED not null AUTO_INCREMENT,
    title VARCHAR(100) not null COMMENT '标题名称',
    summary VARCHAR(250) not null COMMENT '简介',
    content_id int(10) UNSIGNED not null COMMENT '内容的 id',
    created_on  int(10) UNSIGNED DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (id)
)ENGINE = INNODB DEFAULT CHARSET = 'utf8' comment ='文章表';

create TABLE harbor_content(
    id int(10) UNSIGNED not null AUTO_INCREMENT,
    content text COMMENT '文章内容',
    PRIMARY KEY (id)
)ENGINE = INNODB DEFAULT CHARSET = 'utf8' comment ='内容表';

create TABLE harbor_tag(
   id int(10) UNSIGNED not null AUTO_INCREMENT,
   name varchar(50) not null COMMENT '标签名称',
   PRIMARY KEY (id)
)ENGINE = INNODB DEFAULT CHARSET = 'utf8' comment ='标签表';

create TABLE harbor_categories(
   id int(10) UNSIGNED not null AUTO_INCREMENT,
   name varchar(50) not null COMMENT '种类名称',
   PRIMARY KEY (id)
)ENGINE = INNODB DEFAULT CHARSET = 'utf8' comment ='种类表';

create TABLE harbor_article_tag(
    id int(10) UNSIGNED not null AUTO_INCREMENT,
    article_id int(10) UNSIGNED,
    tag_id int(10) UNSIGNED,
    PRIMARY KEY (id)
)ENGINE = INNODB DEFAULT CHARSET = 'utf8' comment ='文章对标签表';


create TABLE harbor_article_categories(
   id int(10) UNSIGNED not null AUTO_INCREMENT,
   article_id int(10) UNSIGNED,
   categories_id int(10) UNSIGNED,
   PRIMARY KEY (id)
)ENGINE = INNODB DEFAULT CHARSET = 'utf8' comment ='文章对种类表';
