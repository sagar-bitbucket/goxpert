CREATE TABLE `question_options` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `question_id` int(11) NOT NULL,
 `is_correct` enum('yes','no') NOT NULL,
 `options` text NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NULL DEFAULT NULL,
 `deleted_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`),
 KEY `question_id` (`question_id`),
 CONSTRAINT `question_id` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1