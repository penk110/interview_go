
### 标准雪花算法
    64位数值
```text
unused	41bits time in milliseconds
0 		0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0

			节点ID
5bits date center_id  5bits worker_id
00000				  00000

12bits sequence_id
0000 0000 0000
```

说明：
1.最高位未使用
2.41位为毫秒时间戳
3.数据中心id
4.机器id
5.循环自增ID,即可算出每毫秒内能够生成的id有 2^12 = 4096，每秒就可以生成 4096 * 1000 约为 409.6 万个


### 索尼公司开源雪花算法
    
```text
unused	39bits time in 10milliseconds，单位变成10ms
0 		0000 0000 0000 0000 0000 0000 0000 0000 0000 000

			序列ID
8bits sequence_id
0000 0000

5bits date center_id  5bits worker_id
0000 0000 0000 

16bits machine_id
0000 0000 0000
```