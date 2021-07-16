package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CheckInclusion567 {
    fun checkInclusion(s1: String, s2: String): Boolean {
        val (n, m) = s1.length to s2.length
        if (m < n) return false
        val s1_cnt = hashMapOf<Char, Int>()
        for (ch in s1)
            s1_cnt[ch] = s1_cnt.getOrDefault(ch, 0) + 1

        val s2_cnt = hashMapOf<Char, Int>()
        fun eq() = s1_cnt.all { it.value == s2_cnt.getOrDefault(it.key, 0) }
        for ((i, ch) in s2.withIndex()) {
            s2_cnt[ch] = s2_cnt.getOrDefault(ch, 0) + 1
            if (i >= n - 1) {
                if (eq()) return true
                val x = s2[i - n + 1]
                s2_cnt[x] = s2_cnt[x]!! - 1
            }
        }
        return false
    }


    @Test
    fun test1() {
        assertEquals(true, checkInclusion("ab", "eidbaooo"))
        assertEquals(false, checkInclusion("ab", "eidboaooo"))
    }
}
