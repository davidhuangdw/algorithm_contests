package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class AreOccurrencesEqual5804 {
    fun areOccurrencesEqual(s: String): Boolean {
        val count = hashMapOf<Char, Int>()
        for (ch in s)
            count[ch] = count.getOrDefault(ch, 0) + 1

        return count.values.toHashSet().size <= 1
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
