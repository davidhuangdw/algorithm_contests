package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class ThreeEqualParts {
    fun threeEqualParts(arr: IntArray): IntArray {
        val n = arr.size
        val a = (0 until n).map { arr[n - 1 - it] }

        var ones = arr.sum()
        if (ones == 0) return intArrayOf(0, n - 1)
        if (ones % 3 != 0) return intArrayOf(-1, -1)
        val part = ones / 3

        var cnt = 0
        val one_pos = IntArray(3)
        for (i in 0 until n) {
            val v = a[i]
            if (v == 1) {
                cnt += 1
                if (cnt % part == 0) one_pos[cnt / part - 1] = i
            }
        }
        if (a.slice(0..one_pos[0]) == a.slice(one_pos[1] - one_pos[0]..one_pos[1])) {
            val l = one_pos[1] - one_pos[0]
            if (a.slice(0..one_pos[0]) == a.slice(one_pos[2] - one_pos[0]..one_pos[2])) {
                return intArrayOf(n - 1 - (one_pos[2] - one_pos[0]), n - l)
            }
        }
        return intArrayOf(-1, -1)
    }


    @Test
    fun test1() {
        assertEquals(parseIntArray("[0,2]").toList(), threeEqualParts(parseIntArray("[1,1,0,0,1]")).toList())
        assertEquals(parseIntArray("[0,3]").toList(), threeEqualParts(parseIntArray("[1,0,1,0,1]")).toList())
        assertEquals(parseIntArray("[-1,-1]").toList(), threeEqualParts(parseIntArray("[1,1,0,1,1]")).toList())
        assertEquals(parseIntArray("[1,5]").toList(), threeEqualParts(parseIntArray("[1,1,0,1,1,0,1,1]")).toList())
    }
}
