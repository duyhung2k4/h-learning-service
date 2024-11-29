# Chạy Vegeta và lưu kết quả
vegeta attack -rate=10 -duration=1s -targets=targets.txt | vegeta report > report.txt

# Tạo biểu đồ (nếu cần)
# vegeta plot < report.txt > plot.html