CREATE TABLE `Release` (
    `Id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `ReleaseKey` varchar(64) NOT NULL DEFAULT '' COMMENT '发布的Key',
    `Name` varchar(64) NOT NULL DEFAULT 'default' COMMENT '发布名字',
    `Comment` varchar(256) DEFAULT NULL COMMENT '发布说明',
    `AppId` varchar(64) NOT NULL DEFAULT 'default' COMMENT 'AppID',
    `ClusterName` varchar(500) NOT NULL DEFAULT 'default' COMMENT 'ClusterName',
    `NamespaceName` varchar(500) NOT NULL DEFAULT 'default' COMMENT 'namespaceName',
    `Configurations` longtext NOT NULL COMMENT '发布配置',
    `DeletedAt` bigint NOT NULL DEFAULT '0' COMMENT 'Delete timestamp based on milliseconds',
    `DataChange_CreatedBy` varchar(64) NOT NULL DEFAULT 'default' COMMENT '创建人邮箱前缀',
    `DataChange_CreatedTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `DataChange_LastModifiedBy` varchar(64) DEFAULT '' COMMENT '最后修改人邮箱前缀',
    `DataChange_LastTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`Id`),
    UNIQUE KEY `UK_ReleaseKey_DeletedAt` (`ReleaseKey`, `DeletedAt`),
    KEY `AppId_ClusterName_GroupName` (
        `AppId`,
        `ClusterName`(191),
        `NamespaceName`(191)
    ),
    KEY `DataChange_LastTime` (`DataChange_LastTime`)
) ENGINE = InnoDB AUTO_INCREMENT = 56 DEFAULT CHARSET = utf8mb4 COMMENT = '发布';