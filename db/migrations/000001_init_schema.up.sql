CREATE TABLE `notifications` (
  `uuid` varchar(255) PRIMARY KEY,
  `user_uuid` varchar(255),
  `message` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `notification_events_queue` (
  `uuid` varchar(255) PRIMARY KEY,
  `notification_uuid` varchar(255) NOT NULL,
  `follower_uuid` varchar(255) NOT NULL,
  `state_uuid` varchar(255) NOT NULL,
  `attempts` int
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `notification_states` (
  `uuid` varchar(255) PRIMARY KEY,
  `state` enum('pending','delivered','failed') NOT NULL,
  `message` varchar(255),
  `requested_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `completed_at` datetime
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `users` (
  `uuid` varchar(255) PRIMARY KEY,
  `name` varchar(255),
  `url` varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `followers` (
  `uuid` varchar(255) PRIMARY KEY,
  `user_uuid` varchar(255),
  `follower_uuid` varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `followers` ADD FOREIGN KEY (`user_uuid`) REFERENCES `users` (`uuid`);

ALTER TABLE `followers` ADD FOREIGN KEY (`follower_uuid`) REFERENCES `users` (`uuid`);

ALTER TABLE `notification_events_queue` ADD FOREIGN KEY (`notification_uuid`) REFERENCES `notifications` (`uuid`);

ALTER TABLE `notification_events_queue` ADD FOREIGN KEY (`state_uuid`) REFERENCES `notification_states` (`uuid`);

ALTER TABLE `notification_events_queue` ADD FOREIGN KEY (`follower_uuid`) REFERENCES `followers` (`uuid`);

ALTER TABLE `notifications` ADD FOREIGN KEY (`user_uuid`) REFERENCES `users` (`uuid`);
