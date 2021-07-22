CREATE TABLE `paymentfirst` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `amount` double NOT NULL,
  `qt` int(11) NOT NULL,
  `desc` varchar(100) NOT NULL,
  PRIMARY KEY (`id`,`qt`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;


INSERT INTO `payments`.`paymentfirst`(`amount`,`qt`,`desc`) VALUES( 1.0, 10, 'test1');
INSERT INTO `payments`.`paymentfirst`(`amount`,`qt`,`desc`) VALUES( 12.0, 2, 'test2');
