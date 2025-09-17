// Chứa hàm kết nối với mysql mà nhiều service dùng chung
// Lưu ý đây là chổ để tạo hàm thôi còn khi main.go của service nào import file này về thì sẽ 
// gọi hàm connect ở đây và điền thêm tham số để phù hợp với service đấy  
package main