CREATE TABLE IF NOT EXISTS `my_books`.`books`
(
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL DEFAULT '',
    `author` VARCHAR(127) NOT NULL DEFAULT '',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    UNIQUE INDEX `name_author_idx`(`name`, `author`),
    FULLTEXT `name_fulltext_idx`(`name`) WITH PARSER ngram
) ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;

CREATE TABLE IF NOT EXISTS `my_books`.`reads`
(
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `book_id` BIGINT(20) NOT NULL,
    `source` VARCHAR(127) NOT NULL DEFAULT '',
    `language` VARCHAR(31) NOT NULL DEFAULT '',
    `finished_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `note` TEXT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
) ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;

INSERT INTO `books` (`name`, `author`)
VALUES ('Khi lỗi thuộc về các vì sao', 'John Green'),
       ('Sherlock Holmes 2', 'Arthur Conan Doyle'),
       ('Chiến binh cầu vồng', 'Khaled Hosseini'),
       ('Nhà nàng ở cạnh nhà tôi', 'Lini Thông Minh'),
       ('Ý tưởng này là của chúng mình', 'Huỳnh Vĩnh Sơn'),
       ('Giết con chim nhại', 'Harper Lee'),
       ('Hoàng tử bé', 'Antoine de Saint-Exupéry'),
       ('Nhà giả kim', 'Paulo Coelho'),
       ('Cà phê cùng Tony', 'Tony Buổi Sáng'),
       ('Ngày người thương một người thuong khác', 'Trí'),
       ('Trên đường băng', 'Tony Buổi Sáng'),
       ('Tuổi trẻ đáng giá bao nhiêu?', 'Rosie Nguyễn'),
       ('Đắc nhân tâm', 'Dale Carnegie'),
       ('Chiến thắng con quỷ trong bạn', 'Napoleon Hill'),
       ('Ai đã lấy miếng pho mát của tôi?', 'Spencer Johnson'),
       ('Sống an vui', 'Khangser Rinpoche'),
       ('Đôi tai thấu suốt thế gian', 'OOPSY'),
       ('Tôi không phải công chúa', 'Kawi'),
       ('Cảm ơn người đã rời xa tôi', 'Hà Thanh Phúc'),
       ('Không gia đình', 'Hector Malot'),
       ('Lạc lối giữa cô đơn', 'Nguyễn Minh Nhật'),
       ('Những vết thương thanh xuân', 'Nhi Thiên'),
       ('Anh ơi đừng đi', 'Hiên'),
       ('Ăn gì để anh mua?', 'Huyền Lê'),
       ('Ta có bi quan không?', 'Khải Đơn'),
       ('Bốn thoả ước', 'Don Miguel Ruiz'),
       ('Bậc thầy của tình yêu', 'Don Miguel Ruiz'),
       ('Nếu biết trăm năm là hữu hạn', 'Phạm Lữ Ân'),
       ('Bạn đắt giá bao nhiêu?', 'Vãn Tình'),
       ('Mình là cá, việc của mình là bơi', 'Takeshi Furukawa'),
       ('Ông già và biển cả', 'Ernest Hemingway'),
       ('Gói nỗi buồn lại và ném đi thật xa', 'Ngọc Hoài Nhân'),
       ('Dám bị ghét', 'Kishimi Ichiro'),
       ('Ngã tư mưa, ngã vào đâu cũng nhớ', 'Hoàng Anh Tú'),
       ('Lối sống tối giản của người Nhật', 'Sasaki Fumio'),
       ('Chuyện con mèo dạy hải âu bay', 'Luis Sepúlveda'),
       ('Ngày xưa có một con bò', 'Camilo Cruz'),
       ('80 lời mẹ gửi con gái', 'Từ Minh'),
       ('Cư xử như đàn bà, suy nghĩ như đàn ông', 'Steve Harvey'),
       ('Đời đơn giản khi ta đơn giản', 'Xuân Nguyễn'),
       ('Tôi nói gì khi nói về chạy bộ', 'Murakami Haruki'),
       ('Lịch sử Việt Nam bằng tranh - Trần Hưng Đạo', 'Trần Bạch Đằng'),
       ('Cây cam ngọt của tôi', 'Jose Mauro de Vansconcelos'),
       ('Giao tiếp bất kỳ ai', 'Jo Condrill'),
       ('Nói nhiều không bằng nói đúng', '2.1/2'),
       ('Chú bé mang Pyjama sọc', 'John Boyne'),
       ('Bước chậm lại giữa thế gian vội vã', 'Hae Min'),
       ('Không diệt không sinh đừng sợ hãi', 'Thích Nhất Hạnh'),
       ('Đi tìm lẽ sống', 'Viktor E. Frankl'),
       ('Dạy con làm giàu - Tập 1', 'Robert T. Kiyosaki'),
       ('Những đòn tâm lý trong bán hàng', 'Brian Tracy'),
       ('Thói quen nguyên tử', 'James Clear'),
       ('Tâm lý học hành vi', 'Khương Nguy')
       ;


