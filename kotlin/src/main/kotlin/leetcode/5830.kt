package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class IsThree5830 {
    fun isThree(n: Int): Boolean {
        var cnt = 0
        for (p in 1..n) {
            if (n % p == 0)
                cnt += 1
        }
        return cnt == 3
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
