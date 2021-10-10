package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class StockPrice5896 {
    val high = PriorityQueue<Pair<Int, Int>>(compareBy { -it.first })
    val low = PriorityQueue<Pair<Int, Int>>(compareBy { it.first })
    val pr = hashMapOf<Int, Int>()
    var max_time = 0

    init {
        pr[max_time] = 0
    }

    fun update(timestamp: Int, price: Int) {
        pr[timestamp] = price
        high.add(price to timestamp)
        low.add(price to timestamp)
        max_time = maxOf(max_time, timestamp)
    }

    fun current(): Int {
        return pr[max_time]!!
    }

    fun get_top(que: PriorityQueue<Pair<Int, Int>>): Int {
        var top = que.peek()
        while (pr[top.second] != top.first) {
            que.poll()
            top = que.peek()
        }
        return top.first
    }

    fun maximum(): Int = get_top(high)

    fun minimum(): Int = get_top(low)


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
