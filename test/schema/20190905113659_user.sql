-- -----------------------------------------------------
-- Table `mydb`.`user`
-- -----------------------------------------------------
-- +migrate Up
CREATE TABLE IF NOT EXISTS `user` (
  `uid` INT NOT NULL AUTO_INCREMENT,
  `uname` VARCHAR(225) NOT NULL,
  `upass` VARCHAR(225) NOT NULL,
  `contact_contactid` INT NOT NULL,
  PRIMARY KEY (`uid`))
ENGINE = InnoDB;

-- +migrate Down
DROP TABLE IF EXISTS `user` ;