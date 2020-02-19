CREATE TABLE `questions` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `section_id` int(11) NOT NULL,
 `sequence_number` tinyint(3) NOT NULL,
 `question_type` enum('program','mcq') NOT NULL,
 `question_title` text NOT NULL,
 `problem_statement` text NOT NULL,
 `test_cases` text NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NULL DEFAULT NULL,
 `deleted_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`),
 KEY `section_id` (`section_id`),
 CONSTRAINT `section_id` FOREIGN KEY (`section_id`) REFERENCES `sections` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1