CREATE TABLE `users` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `uuid` varchar(100) NOT NULL,
 `email` varchar(100) NOT NULL,
 `name` varchar(100) NOT NULL,
 `password` text NOT NULL,
 `designation` varchar(100) NOT NULL,
 `emp_id` varchar(100) NOT NULL,
 `user_type` enum('admin','user') NOT NULL,
 `user_status` enum('created','invited','active','disabled','deleted') NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL DEFAULT NULL,
 PRIMARY KEY (`id`),
 UNIQUE KEY `uuid` (`uuid`),
 UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `users` (`id`, `uuid`, `email`, `name`, `password`, `designation`, `emp_id`, `user_type`, `user_status`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 'bdc5d115-a1d0-4584-be59-27c722d266e5', 'su@admin.com', 'Default Admin', '$2a$04$FddCOm6PXR7Rs524ISWcIev1iotTj76BvK1FKjzken4DlHpd5fm8S', 'Admin', 'EMP001', 'admin', 'created', CURRENT_TIMESTAMP, NULL, NULL);
INSERT INTO `users` (`id`, `uuid`, `email`, `name`, `password`, `designation`, `emp_id`, `user_type`, `user_status`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, '64dedfc7-51ee-4b16-a5e6-96e4ed540971', 'user@user.com', 'Default User', '$2a$04$C092gPZaSfkuvL57jCdhy.4fcTIzWPvT0BJWcTvAObRTJiyYNcCtO', 'SuperUser', 'EMP002', 'user', 'created', CURRENT_TIMESTAMP, NULL, NULL);