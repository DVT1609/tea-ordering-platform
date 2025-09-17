// api-gateway sẽ luôn lắng nghe công 3000. Cổng 3000 sẽ là nơi duy nhất frontend gửi request http đến.
// Rồi api-gateway sẽ đưa request đấy đến các service khác dựa trên cổng mà các service khác nghe 
// middleware/router.go là nơi viết code để xác định request http thuộc service nào
package main