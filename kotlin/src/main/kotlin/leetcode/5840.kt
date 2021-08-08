package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinSwaps5840 {
    fun minSwaps(s: String): Int {
        var diff = 0
        var fail = 0
        for (ch in s) {
            if (ch == '[')
                diff += 1
            else if (diff == 0)
                fail += 1
            else
                diff -= 1
        }
        return (fail + 1) / 2
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
