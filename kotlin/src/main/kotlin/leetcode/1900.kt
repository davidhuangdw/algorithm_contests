package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class EarliestAndLatest {
    fun earliestAndLatest(n: Int, firstPlayer: Int, secondPlayer: Int): IntArray {
        data class Node(val n: Int, val a: Int, val b: Int)

        val dp = hashMapOf<Node, IntArray>()
        fun solve(node: Node): IntArray {
            if (node !in dp) {
                fun calc(): IntArray {
                    val (n, a, b) = node
                    val k = n / 2
                    if (a == n - 1 - b) return intArrayOf(1, 1)
                    var (min, max) = Int.MAX_VALUE to Int.MIN_VALUE
                    for (bits in 0 until (1 shl k)) {
                        var cnt = 0
                        var (na, nb) = -1 to -1
                        for (i in 0 until n) {
                            val lose =
                                (i < k && bits and (1 shl i) == 0) || (n - 1 - i < k && bits and (1 shl (n - 1 - i)) > 0)
                            if (!lose) {
                                if (i == a) na = cnt
                                else if (i == b) nb = cnt
                                cnt += 1
                            }
                        }
                        if (na == -1 || nb == -1) continue
                        val x = solve(Node(n - k, na, nb))
                        min = minOf(min, x[0] + 1)
                        max = maxOf(max, x[1] + 1)
                    }
                    return intArrayOf(min, max)
                }
                dp[node] = calc()
            }
            return dp[node]!!
        }

        val x = solve(Node(n, firstPlayer - 1, secondPlayer - 1))
        val y = solve(Node(n, secondPlayer - 1, firstPlayer - 1))
        return intArrayOf(minOf(x[0], y[0]), maxOf(x[1], y[1]))
    }


    @Test
    fun test1() {
        assertEquals(listOf(3, 4), earliestAndLatest(11, 2, 4).toList())
        assertEquals(listOf(1, 1), earliestAndLatest(5, 1, 5).toList())
    }
}
