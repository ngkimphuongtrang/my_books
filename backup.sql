-- MySQL dump 10.13  Distrib 8.0.31, for Linux (aarch64)
--
-- Host: localhost    Database: my_books
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `books` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `author` varchar(127) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_author_idx` (`name`,`author`),
  FULLTEXT KEY `name_fulltext_idx` (`name`) /*!50100 WITH PARSER `ngram` */ 
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (1,'Khi lỗi thuộc về các vì sao','John Green','2024-03-24 11:41:19','2024-03-24 11:41:19'),(2,'Sherlock Holmes 2','Arthur Conan Doyle','2024-03-24 11:41:19','2024-03-24 11:41:19'),(3,'Chiến binh cầu vồng','Khaled Hosseini','2024-03-24 11:41:19','2024-03-24 11:41:19'),(4,'Nhà nàng ở cạnh nhà tôi','Lini Thông Minh','2024-03-24 11:41:19','2024-03-24 11:41:19'),(5,'Ý tưởng này là của chúng mình','Huỳnh Vĩnh Sơn','2024-03-24 11:41:19','2024-03-24 11:41:19'),(6,'Giết con chim nhại','Harper Lee','2024-03-24 11:41:19','2024-03-24 11:41:19'),(7,'Hoàng tử bé','Antoine de Saint-Exupéry','2024-03-24 11:41:19','2024-03-24 11:41:19'),(8,'Nhà giả kim','Paulo Coelho','2024-03-24 11:41:19','2024-03-24 11:41:19'),(9,'Cà phê cùng Tony','Tony Buổi Sáng','2024-03-24 11:41:19','2024-03-24 11:41:19'),(10,'Ngày người thương một người thuong khác','Trí','2024-03-24 11:41:19','2024-03-24 11:41:19'),(11,'Trên đường băng','Tony Buổi Sáng','2024-03-24 11:41:19','2024-03-24 11:41:19'),(12,'Tuổi trẻ đáng giá bao nhiêu?','Rosie Nguyễn','2024-03-24 11:41:19','2024-03-24 11:41:19'),(13,'Đắc nhân tâm','Dale Carnegie','2024-03-24 11:41:19','2024-03-24 11:41:19'),(14,'Chiến thắng con quỷ trong bạn','Napoleon Hill','2024-03-24 11:41:19','2024-03-24 11:41:19'),(15,'Ai đã lấy miếng pho mát của tôi?','Spencer Johnson','2024-03-24 11:41:19','2024-03-24 11:41:19'),(16,'Sống an vui','Khangser Rinpoche','2024-03-24 11:41:19','2024-03-24 11:41:19'),(17,'Đôi tai thấu suốt thế gian','OOPSY','2024-03-24 11:41:19','2024-03-24 11:41:19'),(18,'Tôi không phải công chúa','Kawi','2024-03-24 11:41:19','2024-03-24 11:41:19'),(19,'Cảm ơn người đã rời xa tôi','Hà Thanh Phúc','2024-03-24 11:41:19','2024-03-24 11:41:19'),(20,'Không gia đình','Hector Malot','2024-03-24 11:41:19','2024-03-24 11:41:19'),(21,'Lạc lối giữa cô đơn','Nguyễn Minh Nhật','2024-03-24 11:41:19','2024-03-24 11:41:19'),(22,'Những vết thương thanh xuân','Nhi Thiên','2024-03-24 11:41:19','2024-03-24 11:41:19'),(23,'Anh ơi đừng đi','Hiên','2024-03-24 11:41:19','2024-03-24 11:41:19'),(24,'Ăn gì để anh mua?','Huyền Lê','2024-03-24 11:41:19','2024-03-24 11:41:19'),(25,'Ta có bi quan không?','Khải Đơn','2024-03-24 11:41:19','2024-03-24 11:41:19'),(26,'Bốn thoả ước','Don Miguel Ruiz','2024-03-24 11:41:19','2024-03-24 11:41:19'),(27,'Bậc thầy của tình yêu','Don Miguel Ruiz','2024-03-24 11:41:19','2024-03-24 11:41:19'),(28,'Nếu biết trăm năm là hữu hạn','Phạm Lữ Ân','2024-03-24 11:41:19','2024-03-24 11:41:19'),(29,'Bạn đắt giá bao nhiêu?','Vãn Tình','2024-03-24 11:41:19','2024-03-24 11:41:19'),(30,'Mình là cá, việc của mình là bơi','Takeshi Furukawa','2024-03-24 11:41:19','2024-03-24 11:41:19'),(31,'Ông già và biển cả','Ernest Hemingway','2024-03-24 11:41:19','2024-03-24 11:41:19'),(32,'Gói nỗi buồn lại và ném đi thật xa','Ngọc Hoài Nhân','2024-03-24 11:41:19','2024-03-24 11:41:19'),(33,'Dám bị ghét','Kishimi Ichiro','2024-03-24 11:41:19','2024-03-24 11:41:19'),(34,'Ngã tư mưa, ngã vào đâu cũng nhớ','Hoàng Anh Tú','2024-03-24 11:41:19','2024-03-24 11:41:19'),(35,'Lối sống tối giản của người Nhật','Sasaki Fumio','2024-03-24 11:41:19','2024-03-24 11:41:19'),(36,'Chuyện con mèo dạy hải âu bay','Luis Sepúlveda','2024-03-24 11:41:19','2024-03-24 11:41:19'),(37,'Ngày xưa có một con bò','Camilo Cruz','2024-03-24 11:41:19','2024-03-24 11:41:19'),(38,'80 lời mẹ gửi con gái','Từ Minh','2024-03-24 11:41:19','2024-03-24 11:41:19'),(39,'Cư xử như đàn bà, suy nghĩ như đàn ông','Steve Harvey','2024-03-24 11:41:19','2024-03-24 11:41:19'),(40,'Đời đơn giản khi ta đơn giản','Xuân Nguyễn','2024-03-24 11:41:19','2024-03-24 11:41:19'),(41,'Tôi nói gì khi nói về chạy bộ','Murakami Haruki','2024-03-24 11:41:19','2024-03-24 11:41:19'),(42,'Lịch sử Việt Nam bằng tranh - Trần Hưng Đạo','Trần Bạch Đằng','2024-03-24 11:41:19','2024-03-24 11:41:19'),(43,'Cây cam ngọt của tôi','Jose Mauro de Vansconcelos','2024-03-24 11:41:19','2024-03-24 11:41:19'),(44,'Giao tiếp bất kỳ ai','Jo Condrill','2024-03-24 11:41:19','2024-03-24 11:41:19'),(45,'Nói nhiều không bằng nói đúng','2.1/2','2024-03-24 11:41:19','2024-03-24 11:41:19'),(46,'Chú bé mang Pyjama sọc','John Boyne','2024-03-24 11:41:19','2024-03-24 11:41:19'),(47,'Bước chậm lại giữa thế gian vội vã','Hae Min','2024-03-24 11:41:19','2024-03-24 11:41:19'),(48,'Không diệt không sinh đừng sợ hãi','Thích Nhất Hạnh','2024-03-24 11:41:19','2024-03-24 11:41:19'),(49,'Đi tìm lẽ sống','Viktor E. Frankl','2024-03-24 11:41:19','2024-03-24 11:41:19'),(50,'Dạy con làm giàu - Tập 1','Robert T. Kiyosaki','2024-03-24 11:41:19','2024-03-24 11:41:19'),(51,'Những đòn tâm lý trong bán hàng','Brian Tracy','2024-03-24 11:41:19','2024-03-24 11:41:19'),(52,'Thói quen nguyên tử','James Clear','2024-03-24 11:41:19','2024-03-24 11:41:19'),(53,'Tâm lý học hành vi','Khương Nguy','2024-03-24 11:41:19','2024-03-24 11:41:19'),(54,'duplicated_name','X','2024-03-24 18:41:19','2024-03-24 18:41:19'),(55,'abc_0','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(56,'abc_1','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(57,'abc_2','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(58,'Atomic Habits','Harper Lee','2024-03-24 18:41:19','2024-03-24 18:41:19'),(59,'abc_3','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(60,'Giết con chim nhại','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(61,'abc_4','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(62,'Chiến binh cầu vồng','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(63,'Who moved my xeese?','X','2024-03-24 18:41:19','2024-03-24 18:41:19'),(64,'abc_5','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(65,'Who moved my cheese?','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(66,'abc_6','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(67,'Đi tìm lẽ sống','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(68,'abc_7','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(69,'abc_8','','2024-03-24 18:41:19','2024-03-24 18:41:19'),(70,'abc_9','','2024-03-24 18:41:19','2024-03-24 18:41:19');
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reads`
--

DROP TABLE IF EXISTS `reads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reads` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `book_id` bigint NOT NULL,
  `source` varchar(127) NOT NULL DEFAULT '',
  `language` varchar(31) NOT NULL DEFAULT '',
  `finished_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `note` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reads`
--

LOCK TABLES `reads` WRITE;
/*!40000 ALTER TABLE `reads` DISABLE KEYS */;
/*!40000 ALTER TABLE `reads` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES (1,0);
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-03-24 13:59:14
