CREATE TABLE IF NOT EXISTS available_port_ranges (
  id INT NOT NULL AUTO_INCREMENT,
  description VARCHAR(255) NULL,
  port_from VARCHAR(45) NULL,
  port_to VARCHAR(45) NULL,
  PRIMARY KEY (id))
ENGINE = InnoDB;
