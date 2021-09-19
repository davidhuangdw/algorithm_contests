package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class LongestSubsequenceRepeatedK5878 {
    fun longestSubsequenceRepeatedK(s: String, k: Int): String {
        val n = s.length
        var firstEnd = hashMapOf<String, Int>()
        val cnt = hashMapOf<Char, Int>()
        val after = List(n) { hashSetOf<Char>() }
        for ((i, ch) in s.withIndex()) {
            cnt[ch] = cnt.getOrDefault(ch, 0) + 1
            if (cnt[ch]!! == k)
                firstEnd[ch.toString()] = i

            for (j in i + 1 until n) {
                val ch = s[j]
                after[i].add(ch)
            }
        }


        for (l in 2..n / k) {
            val nxt = hashMapOf<String, Int>()
            for ((pre, ed) in firstEnd) {
                for (ch in after[ed]) {
                    val cur = pre + ch
                    val m = cur.length
                    if (cur.substring(1, cur.length) !in firstEnd)
                        continue

                    var ed = -1
                    var r = 0
                    for ((i, ch) in s.withIndex()) {
                        if (ch == cur[r % m])
                            r += 1
                        if (r == m * k) {
                            ed = i
                            break
                        }
                    }
                    if (ed >= 0)
                        nxt[cur] = ed
                }
            }
            if (nxt.isEmpty()) break
            firstEnd = nxt
        }
        if (firstEnd.isEmpty()) return ""
        return firstEnd.keys.maxOrNull()!!
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
