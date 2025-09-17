// api-gateway sẽ luôn lắng nghe công 3000. Cổng 3000 sẽ là nơi duy nhất frontend gửi request http đến.
// Rồi api-gateway sẽ đưa request đấy đến các service khác dựa trên cổng mà các service khác nghe 
// proxy/proxy.go là nơi viết code để gửi request http đến service thuộc về request đó
package main