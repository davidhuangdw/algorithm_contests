package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinInteger {
    class Fenwick(val n: Int) {
        val tree = MutableList(n + 1) { 0 } // query range [0, n-1] -> map to [1, n]
        fun preSum(i: Int): Int {   // sum[0..i]
            var (j, sum) = i + 1 to 0
            while (j > 0) {
                sum += tree[j]
                j -= j and -j
            }
            return sum
        }

        fun suffixSum(l: Int) = preSum(n - 1) - preSum(l - 1) // sum[l..(n-1)]
        fun add(i: Int) {
            var j = i + 1
            while (j <= n) {
                tree[j] += 1
                j += j and -j
            }
        }
    }

    fun minInteger(num: String, k: Int): String {
        val n = num.length
        val pos = List(10) { ArrayDeque<Int>() }
        for (i in n - 1 downTo 0)
            pos[num[i] - '0'].add(i)

        val removed = Fenwick(n)    // removed indexes
        var rem = k
        return buildString {
            for (i in 0 until n)
                for (d in 0..9) {
                    if (pos[d].isEmpty()) continue
                    val j = pos[d].last()
                    val jright_removed = removed.suffixSum(j)
                    val dist = j + jright_removed - i
                    if (dist <= rem) {
                        rem -= dist
                        pos[d].removeLast()
                        removed.add(j)
                        append(d)
                        break
                    }
                }
        }
    }

    @Test
    fun test1() {
        assertEquals("1342", minInteger("4321", 4))
        assertEquals("010", minInteger("100", 1))
        assertEquals("0345989723478563548", minInteger("9438957234785635408", 23))
    }
}
