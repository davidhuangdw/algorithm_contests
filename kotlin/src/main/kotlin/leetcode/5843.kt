package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NumOfStrings5843 {
    fun numOfStrings(patterns: Array<String>, word: String): Int {
        var cnt = 0
        for (p in patterns)
            if (p in word)
                cnt += 1
        return cnt
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
