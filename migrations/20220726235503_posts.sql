-- +goose Up
-- +goose StatementBegin
CREATE TABLE `posts` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `title` varchar(255) DEFAULT NULL,
    `summary` varchar(255) DEFAULT NULL,
    `detail` text,
    `file_reference` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_posts_deleted_at` (`deleted_at`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `posts`;
-- +goose StatementEnd
