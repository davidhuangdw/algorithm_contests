package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class GetLucky {
    fun getLucky(s: String, k: Int): Int {
        val cur = s.map { (it - 'a' + 1).toString() }.joinToString("")
        fun calc(x: String, k: Int): Int {
            if (k == 0) return x.toInt()
            return calc(x.map { it - '0' }.sum().toString(), k - 1)
        }
        return calc(cur, k)
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
