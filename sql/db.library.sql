-- --------------------------------------------------------
--
-- Database: library
-- Please see on http://sqlfiddle.com/#!9/eb08b21/1
--

-- --------------------------------------------------------
--
-- Table structure for table authors
--
CREATE TABLE authors (
  id int(10) UNSIGNED NOT NULL,
  first_name varchar(255) NOT NULL,
  middle_name varchar(255) DEFAULT NULL,
  last_name varchar(255) NOT NULL,
  summary varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------
--
-- Table structure for table books
--
CREATE TABLE books (
  id int(10) UNSIGNED NOT NULL,
  title varchar(255) NOT NULL,
  publication DATE DEFAULT NULL,
  summary varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------
--
-- Table structure for table members
--
CREATE TABLE members (
  id int(10) UNSIGNED NOT NULL,
  full_name varchar(255) NOT NULL,
  summary varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------
--
-- Table structure for table book_author
--
CREATE TABLE book_author (
  book_id int(10) UNSIGNED NOT NULL,
  author_id int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------
--
-- Table structure for table book_member
--
CREATE TABLE book_member (
  id         INT(10) UNSIGNED NOT NULL,
  book_id    INT(10) UNSIGNED NOT NULL,
  member_id  INT(10) UNSIGNED NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------
--
-- Primaries for tables
--

--
ALTER TABLE authors
  ADD PRIMARY KEY (id),
  MODIFY id int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
ALTER TABLE books
  ADD PRIMARY KEY (id),
  MODIFY id int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
ALTER TABLE members
  ADD PRIMARY KEY (id),
  MODIFY id int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
ALTER TABLE book_member
  ADD PRIMARY KEY (id),
  MODIFY id int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

-- --------------------------------------------------------
--
-- Constraints for tables
--

--
ALTER TABLE book_author
  ADD PRIMARY KEY (book_id, author_id),
  ADD CONSTRAINT book_author_author_id_foreign FOREIGN KEY (author_id) REFERENCES authors (id),
  ADD CONSTRAINT book_author_book_id_foreign FOREIGN KEY (book_id) REFERENCES books (id);

--
ALTER TABLE book_member
  ADD INDEX (book_id),
  ADD INDEX (member_id),
  ADD CONSTRAINT book_member_book_id_foreign FOREIGN KEY (book_id) REFERENCES books (id),
  ADD CONSTRAINT book_member_member_id_foreign FOREIGN KEY (member_id) REFERENCES members (id);

-- --------------------------------------------------------
--
-- Dumping data for table authors
--
INSERT INTO authors (id, first_name, middle_name, last_name, summary) VALUES
  (1, 'Иван', 'Петрович', 'Писарев', 'Публицист'),
  (2, 'Мария', 'Павловна', 'Фомина', 'Поэтесса'),
  (3, 'Фёдор', 'Васильевич', 'Известный', 'Писатель');

-- --------------------------------------------------------
--
-- Dumping data for table books
--
INSERT INTO books (id, title, publication, summary) VALUES
  (1, 'Ясновидец', '2016-05-20', 'Мистика'),
  (2, 'Игра', '2016-09-01', 'Детектив'),
  (3, 'Фата Моргана', '2016-09-20', 'Триллер'),
  (4, 'Атомный след', '2016-09-25', 'Фентези'),
  (5, 'История ужасов', '2016-10-10', 'Научно-популярный рассказ'),
  (6, 'Мигдаль', '2016-10-30', 'Роман');

-- --------------------------------------------------------
--
-- Dumping data for table authors
--
INSERT INTO members (id, full_name) VALUES
  (1, 'Иван Петрович'),
  (2, 'Мария Павловна'),
  (3, 'Фёдор Васильевич');

-- --------------------------------------------------------
--
-- Dumping data for table book_author
--
INSERT INTO book_author (book_id, author_id) VALUES
  (1, 2),
  (1, 3),
  (1, 1),
  (2, 3),
  (3, 2),
  (3, 1),
  (4, 1),
  (5, 2),
  (6, 2),
  (6, 3);

-- --------------------------------------------------------
--
-- Dumping data for table book_member
--
INSERT INTO book_member (book_id, member_id, created_at, updated_at) VALUES
  (3, 2, TIMESTAMP('2017-10-10 09:45:00'), TIMESTAMP('2017-11-10 10:35:00')),
  (1, 2, TIMESTAMP('2017-11-10 10:10:00'), TIMESTAMP('2017-11-21 18:15:00')),
  (1, 3, TIMESTAMP('2017-11-24 15:05:00'), TIMESTAMP('2017-11-29 16:05:00')),
  (4, 1, TIMESTAMP('2017-11-28 16:00:00'), null),
  (3, 3, TIMESTAMP('2017-11-30 11:11:00'), null),
  (5, 1, TIMESTAMP('2017-12-01 18:50:00'), null),
  (1, 1, TIMESTAMP('2017-12-03 10:20:00'), null),
  (6, 3, TIMESTAMP('2017-12-03 15:15:00'), null);

-- --------------------------------------------------------
--
-- Select books which have three co-authors
--
SELECT
  b.title book_title,
  COUNT(ba.author_id) co_authors
FROM
  book_author ba
INNER JOIN books b ON b.id = ba.book_id
GROUP BY b.id
HAVING co_authors = 3;

-- --------------------------------------------------------
--
-- Select rented books with publication date in September 2016
--
SELECT
  b.title book_title,
  b.publication book_publication,
  COUNT(bm.member_id) rented_count
FROM
  book_member bm
INNER JOIN books b ON b.id = bm.book_id
WHERE
  b.publication BETWEEN TIMESTAMP('2016-09-01 00:00:00') AND TIMESTAMP('2016-09-30 23:59:59')
GROUP BY
  b.id
ORDER BY
  b.publication;

-- --------------------------------------------------------
--
-- Select members which have been rented two or more books
--
SELECT
  m.full_name member_full_name,
  COUNT(bm.book_id) rented_books
FROM
  members m
INNER JOIN book_member bm ON m.id = bm.member_id
WHERE
  bm.updated_at IS NULL
GROUP BY
  m.id
HAVING rented_books >= 2;

-- --------------------------------------------------------
--
-- Select books which have been rented in December 2017
--
SELECT
  m.full_name member_full_name,
  b.title book_title,
  bm.created_at rented_at
FROM
  books b,
  members m,
  book_member bm
WHERE
    b.id = bm.book_id
  AND
    m.id = bm.member_id
  AND
    bm.created_at BETWEEN TIMESTAMP('2017-12-01 00:00:00') AND TIMESTAMP('2017-12-31 23:59:59')
ORDER BY
  m.full_name,
  b.title;

-- --------------------------------------------------------
--
-- Select books which ever have no rents
--
SELECT
  b.id,
  b.title book_title
FROM
  books b
LEFT JOIN book_member bm ON b.id = bm.book_id
WHERE
  bm.book_id IS NULL;
