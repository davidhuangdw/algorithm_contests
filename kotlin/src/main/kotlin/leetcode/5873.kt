package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaxConsecutiveAnswers5873 {
    fun maxConsecutiveAnswers(answerKey: String, k: Int): Int {
        val a = answerKey.map { if (it == 'T') 1 else 0 }
        val n = a.size
        val pre = List(n + 1) { MutableList(2) { 0 } }
        for ((i, x) in a.withIndex())
            for (t in 0..1)
                pre[i + 1][t] = if (x == t) pre[i][t] + 1 else 0
        fun max_count(x: Int): Int {
            val y = x xor 1
            val last = ArrayDeque<Int>()
            var max = 0
            for ((i, v) in a.withIndex()) {
                if (v == x)
                    last.add(i)
                if (last.size > k)
                    last.removeFirst()
                max = maxOf(max, if (last.isEmpty()) i + 1 else (i - last.first() + 1 + pre[last.first()][y]))
            }
            return max
        }
        return maxOf(max_count(0), max_count(1))
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
