// service-order → giỏ hàng, đặt hàng, quản lý đơn hàng (MySQL + Kafka + Aerospike cache)
//server/main.go để gọi hàm kết nối với các sql và kafka từ libs/db rồi điền tham số để kết nối với các sql và kafka
//server/main.go tạo fiber lắng nghe cổng 3003 cho service-order
package main