CREATE TABLE `sections` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `course_id` int(11) NOT NULL,
 `sequence_number` tinyint(3) NOT NULL,
 `name` varchar(100) NOT NULL,
 `study_material` text NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL DEFAULT NULL,
 PRIMARY KEY (`id`),
 KEY `course_id` (`course_id`),
 CONSTRAINT `course_id` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1