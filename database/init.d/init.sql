CREATE Table IF NOT EXISTS `goods-square`.`products` (
    `id` INT AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `type` VARCHAR(10) NOT NULL,

    PRIMARY KEY(`id`)
)ENGINE = InnoDB;