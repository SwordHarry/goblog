create database
if
    not exists blog_service default character
    set utf8mb4 default collate utf8mb4_general_ci;

create table `blog_tag` (
    `id` int(10) unsigned not null auto_increment,
    `name` varchar(100) default '' comment '标签名称',
    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_by` varchar(100) default '' comment '修改人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `deleted_on` int(10) unsigned default  '0' comment  '删除时间',
    `is_del` tinyint(3) unsigned default  '0' comment '是否删除 0未删除、1已删除',
    `state` tinyint(3) default '1' comment '状态 0禁用，1启用',
    primary key (`id`)
) engine=InnoDB default charset=utf8mb4 comment='标签管理';

create table `blog_article` (
    `id` int(10) unsigned not null auto_increment,
    `title` varchar(100) default '' comment '文章标题',
    `desc` varchar(255) default '' comment '文章简述',
    `cover_image_url` varchar(255) default '' comment '封面图片地址',
    `content` longtext comment '文章内容',
    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_by` varchar(100) default '' comment '修改人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `deleted_on` int(10) unsigned default  '0' comment  '删除时间',
    `is_del` tinyint(3) unsigned default  '0' comment '是否删除 0未删除、1已删除',
    `state` tinyint(3) default '1' comment '状态 0禁用，1启用',
    primary key (`id`)
) engine=InnoDB default charset=utf8mb4 comment='文章管理';

create table `blog_article_tag` (
    `id` int(10) unsigned not null auto_increment,
    `article_id` int(11) not null comment '文章id',
    `tag_id` int(10) unsigned not null default '0' comment '标签id',
    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_by` varchar(100) default '' comment '修改人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `deleted_on` int(10) unsigned default  '0' comment  '删除时间',
    `is_del` tinyint(3) unsigned default  '0' comment '是否删除 0未删除、1已删除',
    `state` tinyint(3) default '1' comment '状态 0禁用，1启用',
    primary key (`id`)
) engine=InnoDB default charset=utf8mb4 comment='文章标签关联';

-- 鉴权认证表
create table `blog_auth` (
    `id` int(10) unsigned not null auto_increment,
    `app_key` varchar(20) default '' comment 'key',
    `app_secret` varchar(50) default '' comment 'secret',
    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_by` varchar(100) default '' comment '修改人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `deleted_on` int(10) unsigned default  '0' comment  '删除时间',
    `is_del` tinyint(3) unsigned default  '0' comment '是否删除 0未删除、1已删除',
    primary key(`id`) using btree
) engine=InnoDB default charset=utf8mb4 comment='认证管理';

insert into `blog_service`.`blog_auth`
(`id`,`app_key`,`app_secret`,`created_on`,`created_by`,`modified_on`,`modified_by`,`deleted_on`,`is_del`)
values (1,'hzj','go-blog',0,'hzj',0,'',0,0)
