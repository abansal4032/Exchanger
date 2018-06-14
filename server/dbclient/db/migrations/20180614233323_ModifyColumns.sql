-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

alter table requests modify column `owner_comment` text not null default '', modify column `requester_comment` text not null default '';
alter table entity modify column price int(11) not null default 0, modify `borrower` varchar(50) DEFAULT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

;
