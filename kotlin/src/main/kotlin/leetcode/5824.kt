package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaximumNumber {
    fun maximumNumber(num: String, change: IntArray): String {
        val n = num.length
        val cur = num.toMutableList()
        var i = 0
        while (i < n && change[cur[i] - '0'] <= cur[i] - '0')
            i += 1
        while (i < n && change[cur[i] - '0'] >= cur[i] - '0') {
            cur[i] = '0' + change[cur[i] - '0']
            i += 1
        }

        return cur.joinToString("")
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
