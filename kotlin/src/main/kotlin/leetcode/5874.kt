package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class waysToPartition5874 {
    fun waysToPartition(nums: IntArray, k: Int): Int {
        val a = nums.map { it.toLong() }
        val sum = a.sum()
        val n = nums.size
        var max = 0
        // no change:
        if (sum % 2 == 0L) {
            val hf = sum / 2
            var s = 0L
            for (i in 0 until n - 1) {
                s += a[i]
                if (s == hf) max += 1
            }
        }

        // change:
        val cnt = MutableList(n) { 0 }
        // left
        val pre = hashMapOf<Long, Int>()
        var s = 0L
        for (i in 0 until n - 1) {
            s += a[i]
            pre[s] = pre.getOrDefault(s, 0) + 1
            val cur = sum + k - a[i + 1]
            if (cur % 2 == 0L)
                cnt[i + 1] += pre.getOrDefault(cur / 2, 0)
        }
        // right
        pre.clear()
        s = 0
        for (i in n - 1 downTo 1) {
            s += a[i]
            pre[s] = pre.getOrDefault(s, 0) + 1
            val cur = sum + k - a[i - 1]
            if (cur % 2 == 0L)
                cnt[i - 1] += pre.getOrDefault(cur / 2, 0)
        }
        return cnt.fold(max) { x, y -> maxOf(x, y) }
    }


    @Test
    fun test1() {
        assertEquals(1, waysToPartition(parseIntArray("[2,-1,2]"), 3))
    }
}
