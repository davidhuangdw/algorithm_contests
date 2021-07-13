package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class MedianSlidingWindow480 {
    fun medianSlidingWindow(nums: IntArray, k: Int): DoubleArray {
        val removed = hashMapOf<Long, Int>()
        var (l, r) = 0 to 0
        val (lque, rque) = PriorityQueue<Long>(compareBy { -it }) to PriorityQueue<Long>()

        fun popRemoved(que: PriorityQueue<Long>) {
            while (que.isNotEmpty() && removed.getOrDefault(que.peek(), 0) > 0) {
                removed[que.peek()] = removed[que.peek()]!! - 1
                que.poll()
            }
        }

        val res = DoubleArray(nums.size - k + 1)
        for ((i, v) in nums.withIndex()) {
            if (lque.isEmpty() || v <= lque.peek()) {
                lque.add(v.toLong())
                l += 1
            } else {
                rque.add(v.toLong())
                r += 1
            }
            if (i - k >= 0) {
                val pre = nums[i - k].toLong()
                removed[pre] = removed.getOrDefault(pre, 0) + 1
                if (rque.isNotEmpty() && pre >= rque.peek())
                    r -= 1
                else
                    l -= 1
            }

            popRemoved(rque)
            popRemoved(lque)

            while (r > l) {
                lque.add(rque.poll())
                r -= 1
                l += 1
                popRemoved(rque)
            }
            while (l > r + 1) {
                rque.add(lque.poll())
                l -= 1
                r += 1
                popRemoved(lque)
            }

            if (i - k + 1 >= 0)
                res[i - k + 1] = if (l == r) (0.0 + lque.peek() + rque.peek()) / 2 else lque.peek().toDouble()
        }
        return res
    }

    @Test
    fun test1() {
        assertEquals(
            parseDoubleArray("[1,-1,-1,3,5,6]").toList(),
            medianSlidingWindow(parseIntArray("[1,3,-1,-3,5,3,6,7]"), 3).toList()
        )
        assertEquals(
            parseDoubleArray("[1,1,1]").toList(),
            medianSlidingWindow(parseIntArray("[1,1,1,1]"), 2).toList()
        )
        assertEquals(
            parseDoubleArray("[7,9,3,8,0,2,4,8,3,9]").toList(),
            medianSlidingWindow(parseIntArray("[7,9,3,8,0,2,4,8,3,9]]"), 1).toList()
        )
    }
}
