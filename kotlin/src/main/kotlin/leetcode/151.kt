package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class ReverseWords151 {
    fun reverseWords(s: String): String {
        var (i, n) = 0 to s.length
        return buildString {
            while (i < n) {
                while (i < n && s[i] == ' ')
                    i += 1
                if (i < n) {
                    if (length != 0)
                        append(' ')
                    var j = i
                    while (j < n && s[j] != ' ')
                        j += 1
                    for (k in j - 1 downTo i)
                        append(s[k])
                    i = j + 1
                }
            }
        }.reversed()
    }

    @Test
    fun test1() {
        assertEquals("blue is sky the", reverseWords("the sky is blue"))
        assertEquals("world hello", reverseWords("  hello world  "))
        assertEquals("example good a", reverseWords("a good   example"))
    }
}
