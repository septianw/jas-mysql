-- +migrate Up
INSERT INTO `user` (`uid`, `uname`, `upass`, `contact_contactid`) VALUES (1, "admin", "pass", 1);

-- +migrate Down
DELETE FROM `user` WHERE uid = 1;