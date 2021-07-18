package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CanBeTypedWords {
    fun canBeTypedWords(text: String, brokenLetters: String): Int {
        val b = brokenLetters.toSet()
        fun valid(word: String): Boolean {
            for (ch in word)
                if (ch in b)
                    return false
            return true
        }

        var res = 0
        for (w in text.split(" "))
            if (valid(w)) res += 1

        return res
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
