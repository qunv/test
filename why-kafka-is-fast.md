### Tại sao kafka lại nhanh

1. Low-latency IO
Thông thường để đạt được low-latency IO người ta sẽ sử dụng RAM, tuy nhiên
chi phí về RAM là lớn nếu lượng message là rất lớn, do dó Kafka sử dụng filesystem
để lưu trữ message, khái niệm này có thể được coi là 1 commit log, cấu trúc này chỉ
cho phép thêm mới vào ko thể xóa hoặc sửa đổi message, các message lưu trữ trên
local disk và sắp xếp chúng một cách tuần tự từ trái sang phải và được order theo
thời gian, do đó việc read một message sẽ bỏ qua được thời gian tìm kiếm vì read
sẽ được read message theo tuần tự các message được add vào commit log

2. Zero Copy Priciple
Do kafka sử dụng filesystem để lưu trữ message, mà việc read file vào gửi data trên
network cần nhiều context switch, và việc swicth giữa Kenel Mode và User mode 
sẽ cần đế nhiều băng thông và CPU workload, Kafka apply nguyên lý này để việc
đọc data thông qua RAM và convert trực tiếp data qua Network Interface Card Buffer,
điều này giảm thiểu được số lần switch context và giảm thiểu tối đa số lần CPU involke,
từ đó workload của CPU là nhỏ

3. Copresssion & Batching
Để giảm thiểu số lần gọi Network, Kafka batching data trước khi gửi, đồng thời
kafka nén 1 batch của các message trước khi gửi, điều này rõ ràng là có lợi cho 
băng thông mạng và và xử lý IO

4. Horizontal scalable
Việc cho phép ném vào nhiều machine hơn, cho phép kafka duy trì nhiều partition cho
1 topic, điểu đó đảm bảo việc duy trì thông lượng cao và low-latency