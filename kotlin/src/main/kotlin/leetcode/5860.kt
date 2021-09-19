package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FindOriginalArray5860 {
    fun findOriginalArray(changed: IntArray): IntArray {
        val count = hashMapOf<Int, Int>()
        val ret = mutableListOf<Int>()
        for (v in changed) {
            count[v] = count.getOrDefault(v, 0) + 1
        }
        if (0 in count) {
            if (count[0]!! % 2 != 0)
                return intArrayOf()
            for (i in 0 until count[0]!! / 2)
                ret.add(0)
            count[0] = 0
        }
        for (v in count.keys.sorted()) {
            val c = count[v]!!
            if (c == 0) continue
            if (count.getOrDefault(v * 2, 0) < c)
                return intArrayOf()
            count[v * 2] = count[v * 2]!! - c       // bug: miss the case that count[v*2] > count[v]
            for (i in 0 until c)
                ret.add(v)
        }
        return ret.toIntArray()
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
