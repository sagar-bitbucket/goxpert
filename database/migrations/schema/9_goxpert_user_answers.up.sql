CREATE TABLE `user_answers` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `user_id` int(11) NOT NULL,
 `question_id` int(11) NOT NULL,
 `output` text NOT NULL,
 `answer` text NOT NULL,
 `is_correct` enum('true','false') NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL DEFAULT NULL,
 PRIMARY KEY (`id`),
 KEY `user_id` (`user_id`),
 KEY `question_id` (`question_id`),
 CONSTRAINT `user_answers_ibfk_1` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`),
 CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1