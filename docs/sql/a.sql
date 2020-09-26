create TABLE harbor_article
(
    uuid       CHAR(32),
    title      VARCHAR(100) COMMENT '标题名称',
    summary    VARCHAR(250)     not null COMMENT '简介',
    content_id char(32) COMMENT '内容的 id',
    created_on int(10) UNSIGNED DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (uuid),
    INDEX(title)
) ENGINE = INNODB
  DEFAULT CHARSET = 'utf8' comment ='文章表';

create TABLE harbor_content
(
    uuid    CHAR(32),
    content text COMMENT '文章内容',
    PRIMARY KEY (uuid)
) ENGINE = INNODB
  DEFAULT CHARSET = 'utf8' comment ='内容表';

create TABLE harbor_tag
(
    id   int(10) UNSIGNED not null AUTO_INCREMENT,
    name varchar(50)      not null COMMENT '标签名称',
    PRIMARY KEY (id)
) ENGINE = INNODB
  DEFAULT CHARSET = 'utf8' comment ='标签表';

create TABLE harbor_category
(
    id   int(10) UNSIGNED not null AUTO_INCREMENT,
    name varchar(50)      not null COMMENT '种类名称',
    PRIMARY KEY (id)
) ENGINE = INNODB
  DEFAULT CHARSET = 'utf8' comment ='种类表';

create TABLE harbor_article_tag
(
    id         int(10) UNSIGNED not null AUTO_INCREMENT,
    article_id CHAR(32)         not null,
    tag_id     int(10) UNSIGNED,
    PRIMARY KEY (id),
    INDEX(article_id)
) ENGINE = INNODB
  DEFAULT CHARSET = 'utf8' comment ='文章对标签表';


create TABLE harbor_article_categories
(
    id          int(10) UNSIGNED not null AUTO_INCREMENT,
    article_id  CHAR(32)         NOT null,
    category_id int(10) UNSIGNED,
    PRIMARY KEY (id),
    INDEX(article_id)
) ENGINE = INNODB
  DEFAULT CHARSET = 'utf8' comment ='文章对种类表';

create TABLE harbor_article_content
(
    id          int(10) UNSIGNED not null AUTO_INCREMENT,
    article_id  CHAR(32)         NOT null ,
    content_id CHAR(32)         NOT null,
    PRIMARY KEY (id),
    UNIQUE(article_id)
) ENGINE = INNODB
  DEFAULT CHARSET = 'utf8' comment ='文章对内容表';

