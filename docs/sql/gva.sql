DROP TABLE IF EXISTS `casbin_rule`;

CREATE TABLE `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `exa_customers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_name` varchar(191) DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint(20) unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint(20) unsigned DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `exa_file_chunks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `exa_file_id` bigint(20) unsigned DEFAULT NULL,
  `file_chunk_number` bigint(20) DEFAULT NULL,
  `file_chunk_path` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `exa_files` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(191) DEFAULT NULL,
  `file_md5` varchar(191) DEFAULT NULL,
  `file_path` varchar(191) DEFAULT NULL,
  `chunk_total` bigint(20) DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_files_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `gva_announcements_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(191) DEFAULT NULL COMMENT '公告标题',
  `content` text COMMENT '公告内容',
  `user_id` bigint(20) DEFAULT NULL COMMENT '发布者',
  `attachments` json DEFAULT NULL COMMENT '相关附件',
  PRIMARY KEY (`id`),
  KEY `idx_gva_announcements_info_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `jwt_blacklists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;------------------------------------------------

CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint(20) unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `uni_sys_authorities_authority_id` (`authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_authority_btns` (
  `authority_id` bigint(20) unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `sys_auto_code_histories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `table_name` varchar(191) DEFAULT NULL COMMENT '表名',
  `package` varchar(191) DEFAULT NULL COMMENT '模块名/插件名',
  `request` text COMMENT '前端传入的结构化信息',
  `struct_name` varchar(191) DEFAULT NULL COMMENT '结构体名称',
  `business_db` varchar(191) DEFAULT NULL COMMENT '业务库',
  `description` varchar(191) DEFAULT NULL COMMENT 'Struct中文名称',
  `templates` text COMMENT '模板信息',
  `Injections` text COMMENT '注入路径',
  `flag` bigint(20) DEFAULT NULL COMMENT '[0:创建,1:回滚]',
  `api_ids` varchar(191) DEFAULT NULL COMMENT 'api表注册内容',
  `menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单ID',
  `export_template_id` bigint(20) unsigned DEFAULT NULL COMMENT '导出模板ID',
  `package_id` bigint(20) unsigned DEFAULT NULL COMMENT '包ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_auto_code_packages` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `desc` varchar(191) DEFAULT NULL COMMENT '描述',
  `label` varchar(191) DEFAULT NULL COMMENT '展示名',
  `template` varchar(191) DEFAULT NULL COMMENT '模版',
  `package_name` varchar(191) DEFAULT NULL COMMENT '包名',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_packages_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `sys_base_menu_btns` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL,
  `type` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `sys_base_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint(20) unsigned DEFAULT NULL,
  `parent_id` bigint(20) unsigned DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# Dump of table sys_data_authority_id
# ------------------------------------------------------------

CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_dictionaries` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_dictionary_details` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint(20) unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_export_template_condition` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) DEFAULT NULL COMMENT '模板标识',
  `from` varchar(191) DEFAULT NULL COMMENT '条件取的key',
  `column` varchar(191) DEFAULT NULL COMMENT '作为查询条件的字段',
  `operator` varchar(191) DEFAULT NULL COMMENT '操作符',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_template_condition_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `sys_export_template_join` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) DEFAULT NULL COMMENT '模板标识',
  `joins` varchar(191) DEFAULT NULL COMMENT '关联',
  `table` varchar(191) DEFAULT NULL COMMENT '关联表',
  `on` varchar(191) DEFAULT NULL COMMENT '关联条件',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_template_join_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `sys_export_templates` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `db_name` varchar(191) DEFAULT NULL COMMENT '数据库名称',
  `name` varchar(191) DEFAULT NULL COMMENT '模板名称',
  `table_name` varchar(191) DEFAULT NULL COMMENT '表名称',
  `template_id` varchar(191) DEFAULT NULL COMMENT '模板标识',
  `template_info` text,
  `limit` bigint(20) DEFAULT NULL COMMENT '导出限制',
  `order` varchar(191) DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_templates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_ignore_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) DEFAULT NULL COMMENT 'api路径',
  `method` varchar(191) DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_ignore_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_operation_records` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) DEFAULT NULL COMMENT '请求路径',
  `status` bigint(20) DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(20) DEFAULT NULL COMMENT '延迟',
  `agent` text COMMENT '代理',
  `error_message` varchar(191) DEFAULT NULL COMMENT '错误信息',
  `body` text COMMENT '请求Body',
  `resp` text COMMENT '响应Body',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_params` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '参数名称',
  `key` varchar(191) DEFAULT NULL COMMENT '参数键',
  `value` varchar(191) DEFAULT NULL COMMENT '参数值',
  `desc` varchar(191) DEFAULT NULL COMMENT '参数说明',
  PRIMARY KEY (`id`),
  KEY `idx_sys_params_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sys_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `authority_id` bigint(20) unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint(20) DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  `origin_setting` text COMMENT '配置',
  `user_type` tinyint NOT NULL DEFAULT '0' COMMENT '用户类型：1-前台用户，0-后台用户',
  `parent_id` bigint(20) unsigned DEFAULT NULL COMMENT '直接上级用户ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_sys_users_uuid` (`uuid`),
  KEY `idx_sys_users_username` (`username`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `sys_users`
    ADD COLUMN `user_type` tinyint NOT NULL DEFAULT '0' COMMENT '用户类型：1-前台用户，0-后台用户',
    ADD COLUMN `parent_id` bigint(20) unsigned DEFAULT NULL COMMENT '直接上级用户ID',
    ADD INDEX `idx_parent_id` (`parent_id`);

-- 删除合约表
DROP TABLE IF EXISTS contract_account;
DROP TABLE IF EXISTS contract_order;
DROP TABLE IF EXISTS contract_position;
DROP TABLE IF EXISTS contract_entrust;
DROP TABLE IF EXISTS contract_leverage;

-- 创建合约账户表
CREATE TABLE contract_account (
    id BIGINT ( 20 ) UNSIGNED  PRIMARY KEY COMMENT '合约账户唯一标识',
    user_id BIGINT ( 20 ) UNSIGNED COMMENT '关联用户表的用户 ID',
    total_margin DECIMAL ( 10, 2 ) COMMENT '总保证金',
    available_margin DECIMAL ( 10, 2 ) COMMENT '可用保证金',
    frozen_margin DECIMAL ( 10, 2 ) COMMENT '冻结保证金',
    used_margin DECIMAL ( 10, 2 ) COMMENT '已用保证金',
    realized_profit_loss DECIMAL ( 10, 2 ) COMMENT '已实现盈亏',
    account_status INT UNSIGNED COMMENT '账号状态',
    created_at DATETIME ( 3 ) COMMENT '创建时间',
    updated_at DATETIME ( 3 ) COMMENT '更新时间',
    deleted_at DATETIME ( 3 ) COMMENT '删除时间',
    created_by BIGINT ( 20 ) UNSIGNED COMMENT '创建人ID',
    updated_by BIGINT ( 20 ) UNSIGNED COMMENT '更新人ID',
    deleted_by BIGINT ( 20 ) UNSIGNED COMMENT '删除人ID'
);

-- 创建订单表
CREATE TABLE contract_order (
    id BIGINT ( 20 ) UNSIGNED  PRIMARY KEY COMMENT '订单唯一标识',
    order_number VARCHAR(50) COMMENT '订单号',
    order_time DATETIME ( 3 ) COMMENT '订单时间',
    user_id BIGINT ( 20 ) UNSIGNED COMMENT '关联用户表的用户 ID',
    stock_id BIGINT ( 20 ) UNSIGNED COMMENT '股票id',
    stock_name VARCHAR(255) COMMENT '股票名称',
    order_type INT UNSIGNED COMMENT '下单类型',
    open_price float COMMENT '开仓价格',
    close_price float COMMENT '平仓价格',
    operation_type INT UNSIGNED COMMENT '操作类型',
    quantity float COMMENT '数量',
    order_status INT UNSIGNED COMMENT '订单状态',
    fee DECIMAL(10, 2) COMMENT '手续费',
    realized_profit_loss DECIMAL(10, 2) COMMENT '已实现盈亏',
    created_at DATETIME ( 3 ) COMMENT '创建时间',
    updated_at DATETIME ( 3 ) COMMENT '更新时间',
    deleted_at DATETIME ( 3 ) COMMENT '删除时间',
    created_by BIGINT ( 20 ) UNSIGNED COMMENT '创建人ID',
    updated_by BIGINT ( 20 ) UNSIGNED COMMENT '更新人ID',
    deleted_by BIGINT ( 20 ) UNSIGNED COMMENT '删除人ID'
);

-- 创建持仓表
CREATE TABLE contract_position (
    id BIGINT ( 20 ) UNSIGNED  PRIMARY KEY COMMENT '持仓唯一标识',
    user_id BIGINT ( 20 ) UNSIGNED COMMENT '所在用户表的用户 ID',
    stock_id BIGINT ( 20 ) UNSIGNED COMMENT '股票id',
    stock_name VARCHAR ( 255 ) COMMENT '股票名称',
    position_time DATETIME ( 3 ) COMMENT '持仓时间',
    quantity float COMMENT '持仓数量',
    open_price float COMMENT '开仓价格',
    leverage_ratio INT UNSIGNED COMMENT '杠杆倍数',
    margin DECIMAL ( 10, 2 ) COMMENT '保证金',
    position_amount float COMMENT '持仓金额',
    force_close_price float COMMENT '强平价格',
    position_type INT UNSIGNED COMMENT '持仓类型',
    position_status INT UNSIGNED COMMENT '持仓状态',
    created_at DATETIME ( 3 ) COMMENT '创建时间',
    updated_at DATETIME ( 3 ) COMMENT '更新时间',
    deleted_at DATETIME ( 3 ) COMMENT '删除时间',
    created_by BIGINT ( 20 ) UNSIGNED COMMENT '创建人ID',
    updated_by BIGINT ( 20 ) UNSIGNED COMMENT '更新人ID',
    deleted_by BIGINT ( 20 ) UNSIGNED COMMENT '删除人ID'
);

-- 创建委托交易表
CREATE TABLE contract_entrust (
    id BIGINT ( 20 ) UNSIGNED  PRIMARY KEY COMMENT '委托交易唯一标识',
    user_id BIGINT ( 20 ) UNSIGNED COMMENT '关联用户表的用户 ID',
    position_id BIGINT ( 20 ) UNSIGNED COMMENT '关联持仓表的持仓 ID',
    order_id BIGINT ( 20 ) UNSIGNED COMMENT '关联订单表的订单 ID',
    stock_id BIGINT ( 20 ) UNSIGNED COMMENT '股票id',
    stock_name VARCHAR(255) COMMENT '股票名称',
    trigger_type INT UNSIGNED COMMENT '触发类型',
    trigger_price float COMMENT '触发价格',
    margin DECIMAL ( 10, 2 ) COMMENT '保证金',
    operation_type INT UNSIGNED COMMENT '操作类型',
    quantity float COMMENT '数量',
    entrust_status INT UNSIGNED COMMENT '委托状态',
    created_at DATETIME ( 3 ) COMMENT '创建时间',
    updated_at DATETIME ( 3 ) COMMENT '更新时间',
    deleted_at DATETIME ( 3 ) COMMENT '删除时间',
    created_by BIGINT ( 20 ) UNSIGNED COMMENT '创建人ID',
    updated_by BIGINT ( 20 ) UNSIGNED COMMENT '更新人ID',
    deleted_by BIGINT ( 20 ) UNSIGNED COMMENT '删除人ID'
);

-- 创建合约杠杆表
CREATE TABLE contract_leverage (
    id BIGINT ( 20 ) UNSIGNED  PRIMARY KEY COMMENT '合约杠杆唯一标识',
    user_id BIGINT ( 20 ) UNSIGNED COMMENT '关联用户表的用户 ID',
    stock_id BIGINT ( 20 ) UNSIGNED COMMENT '股票id',
    stock_name VARCHAR(255) COMMENT '股票名称',
    leverage_ratio INT UNSIGNED COMMENT '杠杆倍数',
    created_at DATETIME ( 3 ) COMMENT '创建时间',
    updated_at DATETIME ( 3 ) COMMENT '更新时间',
    deleted_at DATETIME ( 3 ) COMMENT '删除时间',
    created_by BIGINT ( 20 ) UNSIGNED COMMENT '创建人ID',
    updated_by BIGINT ( 20 ) UNSIGNED COMMENT '更新人ID',
    deleted_by BIGINT ( 20 ) UNSIGNED COMMENT '删除人ID'
);


-- 划转/转出合约账户
-- CREATE TABLE fund_transfer (
--     transfer_id INT AUTO_INCREMENT PRIMARY KEY COMMENT '资金划转唯一标识',
--     user_id INT COMMENT '关联用户表的用户 ID',
--     from_account_type ENUM('cash_account', 'margin_account') COMMENT '转出账户类型',
--     to_account_type ENUM('cash_account', 'margin_account') COMMENT '转入账户类型',
--     transfer_amount DECIMAL(10, 2) COMMENT '划转金额',
--     transfer_time DATETIME(3) COMMENT '划转时间',
--     transfer_status ENUM('submitted', 'processing', 'success', 'failure') COMMENT '划转状态',
--     description VARCHAR(255) COMMENT '划转描述（可选，用于记录划转原因等相关信息）',
--     FOREIGN KEY (user_id) REFERENCES user_table(user_id)
-- );

-- 跟单相关开始
CREATE TABLE IF NOT EXISTS `advisor`
(
    `id`                    BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `nick_name`             VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '导师昵称',
    `duty`                  VARCHAR(50)      NULL COMMENT '导师职务',
    `intro`                 VARCHAR(100)     NULL COMMENT '导师介绍',
    `exp`                   TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '从业时间,单位年',
    `avatar_url`            VARCHAR(100)     NULL COMMENT '头像地址',
    `seven_day_return`      DECIMAL(10, 3)   NOT NULL DEFAULT 0 COMMENT '7日盈亏',
    `seven_day_return_rate` DECIMAL(10, 3)   NOT NULL DEFAULT 0 COMMENT '7日收益率%',
    `active_status`         TINYINT          NOT NULL DEFAULT 1 COMMENT '启用状态： 0-关闭；1-启用',
    `created_at`            DATETIME(3)      NULL COMMENT '创建时间',
    `updated_at`            DATETIME(3)      NULL COMMENT '更新时间',
    `deleted_at`            DATETIME(3)      NULL COMMENT '删除时间'
) ENGINE InnoDB
  DEFAULT CHARACTER SET utf8mb4 COMMENT '导师信息表';

CREATE TABLE IF NOT EXISTS `advisor_prod`
(
    `id`              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `advisor_id`      BIGINT UNSIGNED  NOT NULL COMMENT '导师ID',
    `product_name`    VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '产品名称',
    `amount_step`     DECIMAL(10, 2)   NOT NULL DEFAULT 0 COMMENT '步长',
    `max_amount`      DECIMAL(10, 2)   NOT NULL DEFAULT 0 COMMENT '最大金额',
    `min_amount`      DECIMAL(10, 2)   NOT NULL DEFAULT 0 COMMENT '最小金额',
    `follow_period`   INT UNSIGNED              DEFAULT 0 COMMENT '跟卖周期，单位天',
    `commission_rate` DECIMAL(10, 2) UNSIGNED   DEFAULT 0 COMMENT '佣金比例，单位%',
    `auto_renew`      TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否自动续期: 0-关闭;1-开启',
    `active_status`   TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '启用状态：0-关闭;1-开启',
    `created_at`      DATETIME(3)      NULL COMMENT '创建时间',
    `updated_at`      DATETIME(3)      NULL COMMENT '更新时间',
    `deleted_at`      DATETIME(3)      NULL COMMENT '删除时间'
) ENGINE InnoDB
  DEFAULT CHARACTER SET utf8mb4 COMMENT '导师产品表';

CREATE TABLE IF NOT EXISTS `advisor_prod_coupon`
(
    `id`              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `advisor_prod_id` BIGINT UNSIGNED NOT NULL COMMENT '导师产品ID',
    `coupon_id`       BIGINT UNSIGNED NOT NULL COMMENT '优惠券ID',
    `created_at`      DATETIME(3)     NULL COMMENT '创建时间',
    `updated_at`      DATETIME(3)     NULL COMMENT '更新时间',
    `deleted_at`      DATETIME(3)     NULL COMMENT '删除时间'
) ENGINE InnoDB
  DEFAULT CHARACTER SET utf8mb4 COMMENT '导师产品可用优惠券表';

CREATE TABLE IF NOT EXISTS `advisor_stock_order`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `advisor_id` BIGINT UNSIGNED  NOT NULL COMMENT '导师ID',
    `stock_id`   BIGINT UNSIGNED  NOT NULL COMMENT '股票ID',
    `buy_price`  DECIMAL(20, 8)  NULL COMMENT '买入价格',
    `buy_time`   DATE             NULL COMMENT '买入时间',
    `sell_price` DECIMAL(20, 8)  NULL COMMENT '卖出价格',
    `sell_time`  DATE             NULL COMMENT '卖出时间',
    `status`     TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '持仓状态: 1-持仓中;5-已完结',
    `created_at` DATETIME(3)      NULL COMMENT '创建时间',
    `updated_at` DATETIME(3)      NULL COMMENT '更新时间',
    `deleted_at` DATETIME(3)      NULL COMMENT '删除时间'
) ENGINE InnoDB
  DEFAULT CHARACTER SET utf8mb4 COMMENT '导师开单记录表';

CREATE TABLE IF NOT EXISTS `user_follow_order`
(
    `id`                       BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `follow_order_no`          VARCHAR(20)      NOT NULL COMMENT '跟单号',
    `user_id`                  BIGINT UNSIGNED  NOT NULL COMMENT '用户ID',
    `advisor_prod_id`          BIGINT UNSIGNED  NOT NULL COMMENT '导师产品ID',
    `auto_renew`               TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否自动续期: 0-否;1-是',
    `period_start`             DATE             NULL COMMENT '周期开始时间',
    `period_end`               DATE             NULL COMMENT '周期结束时间',
    `advisor_commission_rate`  DECIMAL(10, 2)   NULL COMMENT '导师佣金比例(冗余快照)',
    `platform_commission_rate` DECIMAL(10, 2)   NULL COMMENT '平台佣金比例(冗余快照)',
    `follow_order_status`      TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户跟单状态: 0-审核中;3-已驳回;5-跟单中;8-已撤回;10-已结束',
    `stock_status`             TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '持仓状态： 0-未开单；1-持仓中；3-已过期',
    `follow_amount`            DECIMAL(10, 2)   NOT NULL COMMENT '跟单金额',
    `follow_available_amount`  DECIMAL(10, 2)   NOT NULL COMMENT '跟单可用金额',
    `retrievable_amount`       DECIMAL(10, 2)   NOT NULL DEFAULT 0 COMMENT '可提取金额',
    `coupon_record_id`         BIGINT           NULL COMMENT '使用优惠券ID',
    `apply_time`               DATETIME(3)      NOT NULL COMMENT '申请时间',
    `end_time`                 DATETIME(3)      NULL COMMENT '结束时间',
    `created_at`               DATETIME(3)      NULL COMMENT '创建时间',
    `updated_at`               DATETIME(3)      NULL COMMENT '更新时间',
    `deleted_at`               DATETIME(3)      NULL COMMENT '删除时间'
) ENGINE InnoDB
  DEFAULT CHARACTER SET utf8mb4 COMMENT '用户跟单订单记录';

CREATE TABLE IF NOT EXISTS `user_follow_stock_detail`
(
    `id`                  BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `user_id`             BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `stock_order_id`      BIGINT UNSIGNED NOT NULL COMMENT '导师开单记录ID',
    `follow_order_id`     BIGINT UNSIGNED NOT NULL COMMENT '用户跟单记录ID',
    `stock_id`            BIGINT UNSIGNED NOT NULL COMMENT '股票ID',
    `stock_num`           DECIMAL(20, 8)  NOT NULL COMMENT '股票数量',
    `buy_price`           DECIMAL(20, 8)  NULL COMMENT '买入价格',
    `buy_time`            DATETIME(3)     NULL COMMENT '买入时间',
    `sell_price`          DECIMAL(20, 8)  NULL COMMENT '卖出价格',
    `sell_time`           DATETIME(3)     NULL COMMENT '卖出时间',
    `advisor_commission`  DECIMAL(10, 2)  NULL COMMENT '导师佣金',
    `platform_commission` DECIMAL(10, 2)  NULL COMMENT '平台佣金',
    `created_at`          DATETIME(3)     NULL COMMENT '创建时间',
    `updated_at`          DATETIME(3)     NULL COMMENT '更新时间',
    `deleted_at`          DATETIME(3)     NULL COMMENT '删除时间'
) ENGINE InnoDB
  DEFAULT CHARACTER SET utf8mb4 COMMENT '用户跟单股票详情表';

CREATE TABLE IF NOT EXISTS `user_follow_append_order`
(
    `id`              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `append_order_no` VARCHAR(20)     NOT NULL COMMENT '追单号',
    `user_id`         BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `follow_order_id` BIGINT UNSIGNED NOT NULL COMMENT '用户跟单订单ID',
    `append_amount`   DECIMAL(10, 2)  NOT NULL COMMENT '追单金额',
    `created_at`      DATETIME(3)     NULL COMMENT '创建时间',
    `updated_at`      DATETIME(3)     NULL COMMENT '更新时间',
    `deleted_at`      DATETIME(3)     NULL COMMENT '删除时间'
) ENGINE InnoDB
  DEFAULT CHARACTER SET utf8mb4 COMMENT '用户追单记录表';
  
create table if not exists `user_profit_share`
(
    `id`              bigint unsigned auto_increment primary key comment '主键',
    `from_user_id`    bigint unsigned not null comment '来源用户ID',
    `to_user_id`      bigint unsigned not null comment '目标用户ID',
    `follow_order_id` bigint unsigned not null comment '分成来源跟单ID',
    `amount`          decimal(20, 2)  not null default 0 comment '分成金额',
    `share_rate`      decimal(10, 2)  not null default 0 comment '分成比例',
    `created_at`      datetime(3)     null comment '创建时间',
    `updated_at`      datetime(3)     null comment '更新时间',
    `deleted_at`      datetime(3)     null comment '删除时间'
) ENGINE = InnoDB
  DEFAULT CHARACTER SET 'utf8mb4' COMMENT '用户分成流水记录表';

CREATE TABLE IF NOT EXISTS `system_order_msg`
(
    `id`          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '记录ID',
    `type`        TINYINT UNSIGNED NOT NULL COMMENT '消息类型： 0-新审核订单；1-新追加订单',
    `order_id`    BIGINT           NOT NULL COMMENT '关联订单ID',
    `user_id`     BIGINT UNSIGNED  NOT NULL COMMENT '用户ID',
    `read_status` TINYINT UNSIGNED NOT NULL COMMENT '已读状态: 0-未读；1-已读',
    `created_at`  DATETIME(3)      NULL COMMENT '创建时间',
    `updated_at`  DATETIME(3)      NULL COMMENT '更新时间',
    `deleted_at`  DATETIME(3)      NULL COMMENT '删除时间'
) ENGINE InnoDB
  DEFAULT CHARACTER SET utf8mb4 COMMENT '系统订单站内信表';

CREATE TABLE IF NOT EXISTS `system_config`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '记录ID',
    `config`     json        NOT NULL COMMENT '配置信息',
    `created_at` DATETIME(3) NULL COMMENT '创建时间',
    `updated_at` DATETIME(3) NULL COMMENT '更新时间',
    `deleted_at` DATETIME(3) NULL COMMENT '删除时间'
) ENGINE InnoDB
  DEFAULT CHARACTER SET utf8mb4 COMMENT '系统配置';

create table if not exists `stock_exchange_holiday`
(
    `id`         bigint unsigned auto_increment primary key comment '记录ID',
    `name`       varchar(50) not null default '' comment '节日名称',
    `exchange`   varchar(50) not null default '' comment '交易所名称',
    `date`       date        not null comment '节日日期',
    `status`     varchar(50) comment '开市状态',
    `created_at` datetime(3) null comment '创建时间',
    `updated_at` datetime(3) null comment '更新时间',
    `deleted_at` datetime(3) null comment '删除时间'
) ENGINE = InnoDB
  default character set utf8mb4 comment '交易所节假日信息表';

create table if not exists `platform_report_summary`
(
    id                          bigint unsigned auto_increment primary key comment '主键',
    root_user_id                bigint unsigned not null default 0 comment '跟用户ID',
    total_balance               decimal(20, 2)  not null default 0 comment '盘内结余',
    total_recharge              decimal(20, 2)  not null default 0 comment '历史总充值',
    total_withdraw              decimal(20, 2)  not null default 0 comment '历史总提现',
    total_user_count            int             not null default 0 comment '用户总数',
    total_follow_user_count     int             not null default 0 comment '跟买用户总数',
    platform_profit             decimal(20, 2)  not null default 0 comment '平台总盈亏',
    today_active_user_count     int             not null default 0 comment '今日活跃用户',
    yesterday_active_user_count int             not null default 0 comment '昨日活跃用户',
    created_at                  datetime(3)     null comment '创建时间',
    updated_at                  datetime(3)     null comment '更新时间',
    deleted_at                  datetime(3)     null comment '删除时间'
) ENGINE InnoDB
  DEFAULT CHARACTER SET utf8mb4 COMMENT '平台固定报表';

create table if not exists `platform_report_daily`
(
    id                        bigint unsigned auto_increment primary key comment '主键',
    root_user_id              bigint unsigned not null default 0 comment '跟用户ID',
    total_recharge            decimal(20, 2)  not null default 0 comment '总充值',
    recharge_user_count       int             not null default 0 comment '充值用户数',
    total_withdraw            decimal(20, 2)  not null default 0 comment '总提现',
    withdraw_user_count       int             not null default 0 comment '提现用户数',
    total_new_user_count      int             not null default 0 comment '新增用户总数',
    active_user_count         int             not null default 0 comment '活跃用户数',
    first_recharge_user_count int             not null default 0 comment '首冲用户总数',
    first_recharge_amount     decimal(20, 2)  not null default 0 comment '首冲总金额',
    follow_profit             decimal(20, 2)  not null default 0 comment '跟单总盈利',
    follow_amount             decimal(20, 2)  not null default 0 comment '跟单总金额',
    follow_user_count         int             not null default 0 comment '跟单用户数',
    report_date               date            not null comment '报告日期',
    created_at                datetime(3)     null comment '创建时间',
    updated_at                datetime(3)     null comment '更新时间',
    deleted_at                datetime(3)     null comment '删除时间'
) ENGINE InnoDB
  DEFAULT CHARACTER SET utf8mb4 COMMENT '平台日报表';
-- 跟单相关结束

CREATE TABLE symbols (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    symbol VARCHAR(10) NOT NULL UNIQUE COMMENT '股票代码，通常是股票的唯一标识',
    corporation VARCHAR(255) COMMENT '公司名称',
    industry VARCHAR(255) COMMENT '行业',
    exchange VARCHAR(50) COMMENT '所属交易所',
    market_cap BIGINT COMMENT '市值',
    ipo_date DATE COMMENT 'IPO 上市日期',
    current_price DECIMAL(10, 2) COMMENT '最新价格',
    average_volume BIGINT COMMENT '日均交易量',
    change_ratio DECIMAL(10, 2) COMMENT '日涨跌幅度',
    pe_ratio DECIMAL(10, 2) COMMENT '市盈率',
    created_at DATETIME(3) COMMENT '创建时间',
    updated_at DATETIME(3) COMMENT '更新时间'
);

CREATE TABLE symbols_custom (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    user_id BIGINT COMMENT '用户ID',
    symbol_id BIGINT COMMENT '股票ID',
    created_at DATETIME(3) COMMENT '创建时间',
    updated_at DATETIME(3) COMMENT '更新时间'
);

CREATE TABLE symbols_hot (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    symbol_id BIGINT COMMENT '股票ID',
    created_at DATETIME(3) COMMENT '创建时间',
    updated_at DATETIME(3) COMMENT '更新时间'
);

