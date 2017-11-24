-- --------------------------------------------------------
--
-- Database: emails
--

-- --------------------------------------------------------
--
-- Table structure for table users
--
CREATE TABLE users (
  id int(10) UNSIGNED NOT NULL,
  email varchar(255) NOT NULL,
  full_name varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------
--
-- Table structure for table letters
--
CREATE TABLE letters (
  id int(10) UNSIGNED NOT NULL,
  sender_email varchar(255) NOT NULL,
  recipient_email varchar(255) NOT NULL,
  subject varchar(255) DEFAULT NULL,
  message text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------
--
-- Table structure for table folders
--
CREATE TABLE folders (
  id int(10) UNSIGNED NOT NULL,
  name varchar(255) NOT NULL,
  summary varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------
--
-- Table structure for table user_folder
--
CREATE TABLE user_folder (
  user_id int(10) UNSIGNED NOT NULL,
  folder_id int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------
--
-- Table structure for table letter_folder
--
CREATE TABLE letter_folder (
  letter_id int(10) UNSIGNED NOT NULL,
  folder_id int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- --------------------------------------------------------
--
ALTER TABLE users
  ADD PRIMARY KEY (id),
  ADD UNIQUE (email),
  MODIFY id int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

-- --------------------------------------------------------
--
ALTER TABLE folders
  ADD PRIMARY KEY (id),
  MODIFY id int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

-- --------------------------------------------------------
--
ALTER TABLE letters
  ADD PRIMARY KEY (id),
  ADD INDEX (sender_email),
  ADD INDEX (recipient_email),
  MODIFY id int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

-- --------------------------------------------------------
--
ALTER TABLE user_folder
  ADD PRIMARY KEY (user_id, folder_id),
  ADD INDEX (folder_id),
  ADD CONSTRAINT user_folder_user_id_foreign FOREIGN KEY (user_id) REFERENCES users (id),
  ADD CONSTRAINT user_folder_folder_id_foreign FOREIGN KEY (folder_id) REFERENCES folders (id);

-- --------------------------------------------------------
--
ALTER TABLE letter_folder
  ADD PRIMARY KEY (letter_id, folder_id),
  ADD INDEX (folder_id),
  ADD CONSTRAINT letter_folder_letter_id_foreign FOREIGN KEY (letter_id) REFERENCES letters (id),
  ADD CONSTRAINT letter_folder_folder_id_foreign FOREIGN KEY (folder_id) REFERENCES folders (id);

-- --------------------------------------------------------
--
SELECT
  u.id,
  u.email,
  u.full_name
FROM
  users u
INNER JOIN letters il ON il.sender_email = u.email
LEFT  JOIN letters ll ON ll.recipient_email = u.email
WHERE
  ll.recipient_email IS NULL
