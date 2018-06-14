-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `enitity_attributes` (
  `entity_id` varchar(50) NOT NULL DEFAULT '',
  `attribute_key` varchar(50) NOT NULL DEFAULT '',
  `attribute_value` varchar(50) NOT NULL DEFAULT '',
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `entity_attributes_entity_entity_id` (`entity_id`),
  CONSTRAINT `entity_attributes_entity_entity_id` FOREIGN KEY (`entity_id`) REFERENCES `entity` (`entity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `entity` (
  `entity_id` varchar(50) NOT NULL DEFAULT '',
  `entity_name` varchar(50) NOT NULL DEFAULT '',
  `entity_type` varchar(50) NOT NULL DEFAULT '',
  `owner` varchar(50) NOT NULL DEFAULT '',
  `action_type` enum('SELL','SHARE') NOT NULL DEFAULT 'SHARE',
  `status` varchar(50) NOT NULL DEFAULT '',
  `price` int(11) DEFAULT NULL,
  `borrower` varchar(50) DEFAULT '',
  `location` int(11) NOT NULL,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`entity_id`),
  KEY `entity_user_user_id` (`owner`),
  KEY `entity_borrower_user_user_id` (`borrower`),
  CONSTRAINT `entity_borrower_user_user_id` FOREIGN KEY (`borrower`) REFERENCES `user` (`user_id`),
  CONSTRAINT `entity_user_user_id` FOREIGN KEY (`owner`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `requests` (
  `request_id` varchar(50) NOT NULL DEFAULT '',
  `entity_id` varchar(50) NOT NULL DEFAULT '',
  `requester` varchar(50) NOT NULL DEFAULT '',
  `intent` enum('BUY','RENT') NOT NULL DEFAULT 'RENT',
  `duration_in_days` int(11) NOT NULL,
  `status` varchar(50) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `requester_comment` text,
  `owner_comment` text,
  PRIMARY KEY (`request_id`),
  KEY `entity_entity_id` (`entity_id`),
  CONSTRAINT `entity_entity_id` FOREIGN KEY (`entity_id`) REFERENCES `entity` (`entity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `user` (
  `user_id` varchar(50) NOT NULL DEFAULT '',
  `name` varchar(50) NOT NULL DEFAULT '',
  `contact_number` varchar(10) NOT NULL DEFAULT '',
  `email` varchar(50) NOT NULL DEFAULT '',
  `location` int(11) NOT NULL,
  `credits` int(11) NOT NULL,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

;
