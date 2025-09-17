// api-gateway sẽ luôn lắng nghe công 3000. Cổng 3000 sẽ là nơi duy nhất frontend gửi request http đến.
// Rồi api-gateway sẽ đưa request đấy đến các service khác dựa trên cổng mà các service khác nghe 
// gateway/main.go là nơi viết code lắng nghe cổng 3000
package main