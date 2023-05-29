-- CreateTable
CREATE TABLE `users` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(150) NOT NULL,
    `role` ENUM('USER', 'COACH') NOT NULL DEFAULT 'USER',
    `timezone` VARCHAR(191) NOT NULL,
    `created_at` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` VARCHAR(191) NULL,
    `updated_at` TIMESTAMP(0) NULL,
    `updated_by` VARCHAR(191) NULL,

    INDEX `users_role_idx`(`role`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `working_hours` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `day` SMALLINT NOT NULL,
    `start` TIME NOT NULL,
    `end` TIME NOT NULL,
    `created_at` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` VARCHAR(191) NULL,
    `updated_at` TIMESTAMP(0) NULL,
    `updated_by` VARCHAR(191) NULL,

    INDEX `working_hours_user_id_day_idx`(`user_id`, `day`),
    INDEX `working_hours_day_idx`(`day`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `appointments` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `coach_id` BIGINT UNSIGNED NOT NULL,
    `status` ENUM('SCHEDULING', 'SCHEDULED', 'CANCELED', 'RESCHEDULING', 'RESCHEDULED', 'DECLINED') NOT NULL DEFAULT 'SCHEDULING',
    `rescheduled` BOOLEAN NOT NULL DEFAULT false,
    `start_at` TIMESTAMP(0) NOT NULL,
    `end_at` TIMESTAMP(0) NOT NULL,
    `duration_minutes` INTEGER NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` VARCHAR(191) NULL,
    `updated_at` TIMESTAMP(0) NULL,
    `updated_by` VARCHAR(191) NULL,

    INDEX `appointments_user_id_coach_id_idx`(`user_id`, `coach_id`),
    INDEX `appointments_coach_id_idx`(`coach_id`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
