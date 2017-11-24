-- --------------------------------------------------------
--
-- Database: test
-- Please see on http://sqlfiddle.com/#!9/891eeb/1
--

-- --------------------------------------------------------
--
-- Table structure for table `list`
--
CREATE TABLE list (
  id int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------
--
-- Dumping data for table `test`
--
INSERT INTO list (id) VALUES
  (1),
  (2),
  (3),
  (6),
  (8),
  (9),
  (12),
  (13),
  (14),
  (20),
  (25);

-- --------------------------------------------------------
--
-- Show two columns as FROM - TILL diapason which will be content missing ids into the scope
--
SELECT
  s1.id 'FROM',
  s2.id 'TILL'
FROM
  (
    SELECT
      (@rn1:=@rn1+1) rn,
      mt.id
    FROM
      list mt,
      (SELECT @rn1:=0) r
    WHERE
      mt.id != (SELECT MAX(id) FROM list)
      AND
      (SELECT 1 FROM list AS st WHERE st.id = mt.id+1) IS NULL
  ) s1,
  (
    SELECT
      (@rn2:=@rn2+1) rn,
      mt.id
    FROM
      list mt,
      (SELECT @rn2:=0) r
    WHERE
      mt.id != (SELECT MIN(id) FROM list)
      AND
      (SELECT 1 FROM list AS st WHERE st.id = mt.id-1) IS NULL
  ) s2
WHERE s1.rn = s2.rn
