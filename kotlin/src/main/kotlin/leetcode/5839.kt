package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class MinStoneSum5839 {
    fun minStoneSum(piles: IntArray, kk: Int): Int {
        val cur = PriorityQueue<Pair<Int, Int>>(compareBy { -it.first })
        for ((i, v) in piles.withIndex())
            cur.add(v to i)
        var k = kk
        while (k > 0) {
            val (n, i) = cur.poll()
            if (piles[i] != n)
                continue
            piles[i] -= piles[i] / 2
            cur.add(piles[i] to i)
            k -= 1
        }
        return piles.sum()
    }


    @Test
    fun test1() {
        assertEquals(1, minStoneSum(parseIntArray("[5,4,9]"), 2))
    }
}
