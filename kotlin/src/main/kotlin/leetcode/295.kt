package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class MedianFinder295 {
    class MedianFinder() {
        /** initialize your data structure here. */
        val left = PriorityQueue<Int>(compareBy { -it })
        val right = PriorityQueue<Int>(compareBy { it })


        fun addNum(num: Int) {
            if (left.isEmpty() || num <= left.peek()) {
                left.add(num)
                if (left.size > right.size + 1) {
                    right.add(left.poll())
                }
            } else {
                right.add(num)
                if (right.size > left.size)
                    left.add(right.poll())
            }
        }

        fun findMedian(): Double {
            return if (left.size == right.size) (left.peek() + right.peek()) / 2.0 else left.peek().toDouble()
        }
    }
}
