package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class StoneGameIX5892 {
    fun stoneGameIX(stones: IntArray): Boolean {
        val n = stones.size
        val cnt = MutableList(3) { 0 }
        for (v in stones)
            cnt[v % 3] += 1
        for (x in 1..2)
            if (cnt[x] > 0) {
                val y = 3 - x
                val t = minOf(cnt[x] - 1, cnt[y])
                var cur = 1 + cnt[0] + t * 2
                if (cnt[x] - 1 > t)
                    cur += 1
                if (cur < n && cur % 2 == 1)
                    return true
            }
        return false
    }


    @Test
    fun test1() {
        assertEquals(false, stoneGameIX(parseIntArray("[2]")))
        assertEquals(true, stoneGameIX(parseIntArray("[2,1]")))
        assertEquals(false, stoneGameIX(parseIntArray("[5,1,2,4,3]")))
    }
}
