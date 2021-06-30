package leetcode

import org.junit.jupiter.api.Assertions
import org.junit.jupiter.api.Test

class LargestOddNumber {
    fun largestOddNumber(num: String): String {
        var i = num.length - 1
        while (i >= 0 && num[i] in "02468") i -= 1
        return num.substring(0, i + 1)
    }

    @Test
    fun test1() {
        Assertions.assertEquals("5", largestOddNumber("52"))
        Assertions.assertEquals("", largestOddNumber("4206"))
        Assertions.assertEquals("35427", largestOddNumber("35427"))
    }
}