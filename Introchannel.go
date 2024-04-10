package main

/* func main() {
// channel 是一种在协程之间通信的机制 通过它们可以安全的传递数据
/*
	1. channel 是一种引用类型 作为参数传递时传递的是引用
	2. 默认情况下 创建的channel 是无缓存的 只有当接收者准备好接收数据时发送者才能发送数据
	3. 通过make创建channel 可以指定缓冲区的大小
	4. channel 是类型相关的 一个channel只能传递一种类型的数据
	5. channel 支持多路复用 可以使用select语句监听多个channel 它允许多个信号或数据流共享一个通信信道或媒介
	6. channel 可以用于控制gorouter 例如关闭一个gorouter
*/

/*
	1. 一般而言channel 是被初始化为一个FIFO的队列 当发送数据时候放在队尾
	从队头取出数据 当channel为空时接收者会阻塞 当channel满时发送者会阻塞
	2. 所有有了一种新的channel 即ring buffer 环形缓冲区
	3. 环形缓冲区是一种特殊的缓冲区 他的长度是固定的 当缓冲区满时新的数据会覆盖旧的数据
	4. 环形缓冲区的实现是通过一个数组和两个指针来实现的
	5. 通过环形缓冲区可以实现一个无锁的队列
	需要注意的是
	1. 避免了动态的内存分配
	2. 提高了缓存命中率
	并且大小要是2的幂次方
*/
//}