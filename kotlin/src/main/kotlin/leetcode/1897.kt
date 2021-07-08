package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MakeEqual {
    fun makeEqual(words: Array<String>): Boolean {
        val n = words.size
        val count = hashMapOf<Char, Int>()
        for (w in words)
            for (ch in w)
                count[ch] = count.getOrDefault(ch, 0) + 1
        for ((k, v) in count)
            if (v % n != 0)
                return false
        return true
    }
}
