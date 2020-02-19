CREATE TABLE `user_results` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `user_id` int(11) NOT NULL,
 `user_answers_id` int(11) NOT NULL,
 `is_correct` enum('yes','no') NOT NULL,
 `coverage` float(10,2) NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL DEFAULT NULL,
 PRIMARY KEY (`id`),
 KEY `user_id_constraint` (`user_id`),
 KEY `user_answer` (`user_answers_id`),
 CONSTRAINT `user_answer` FOREIGN KEY (`user_answers_id`) REFERENCES `user_answers` (`id`),
 CONSTRAINT `user_id_constraint` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1