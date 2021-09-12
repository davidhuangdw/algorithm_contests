package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class ReversePrefix {
    fun reversePrefix(word: String, ch: Char): String {
        var k = 0
        for((i, c) in word.withIndex())
            if(c == ch) {
                k = i + 1
                break
            }
        return word.substring(0, k).reversed() + word.substring(k, word.length)
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
