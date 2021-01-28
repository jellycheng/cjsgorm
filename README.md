# cjsgorm
```
本仓库是对gosupport以及gorm通用能力进行封装，以便更好的适应公司项目复用
gorm网站： https://gorm.io/zh_CN/
gosupport库： https://github.com/jellycheng/gosupport

```

## 引用
```
go get -u github.com/jellycheng/cjsgorm

```

## 使用
```
参考gorm_test.go文件

```

## demo sql
```
db name: db_common

DROP TABLE IF EXISTS `t_system`;
CREATE TABLE `t_system` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `system_code` varchar(32) NOT NULL DEFAULT '' COMMENT '系统Code',
  `system_name` varchar(50) NOT NULL DEFAULT '' COMMENT '系统名称',
  `app_id` varchar(32) DEFAULT NULL COMMENT 'AppID',
  `secret` varchar(32) DEFAULT NULL COMMENT '密钥',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除; 0-正常; 1-删除',
  `create_time` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `delete_time` int(10) NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `unq_sy_co` (`system_code`),
  UNIQUE KEY `unq_app_id` (`app_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='系统表';

```
