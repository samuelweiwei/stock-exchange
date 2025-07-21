alter table `frontend_users` add user_type tinyint(1) unsigned default 1 comment '用户类型 1:普通用户 2: 试玩用户';
alter table `platform_report_summary` add column  `invite_user_count` int not null default 0 comment '邀请用户数量' after `yesterday_active_user_count`;
alter table `platform_report_daily` add column  `invite_user_count` int not null default 0 comment '邀请用户数量' after `active_user_count`;

// frontned_users 表增加 login_time 字段, 记录登录毫秒时间戳
alter table `frontend_users` add column `last_login_time` bigint unsigned not null default 0 comment '最近登录时间' after `last_login_ip`;

alter table platform_report_summary add column  effective_invite_user_count bigint not null default 0 comment '有效邀请用户数' after invite_user_count;

alter table platform_report_daily add total_actual_withdraw decimal(20,2) not null default 0 comment '实际总提现' after total_withdraw;
alter table platform_report_summary add total_actual_withdraw decimal(20,2) not null default 0 comment '实际总提现' after total_withdraw;