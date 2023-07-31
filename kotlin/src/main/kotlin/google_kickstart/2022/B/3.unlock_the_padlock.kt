package google_kickstart.`2022`.B

import kotlin.math.absoluteValue

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, D) = input_ints()
        val V = input_ints()
        fun ds(d: Int) = minOf(d.absoluteValue, D - d.absoluteValue)
        fun set2(): Long {
            var dp = List(N + 1) { List(N + 1) { MutableList(D) { 0L } } }
            for (d in 1..N) {
                for (l in 0..N - d) {
                    for (k in 0 until D) {
                        dp[d][l][k] =
                            (listOf(l to V[l + d - 1], l + 1 to V[l])).minOf { p ->
                                val (l, b) = p
                                // bug: must all go to b first! cannot change b independently
                                dp[d - 1][l][b] + ds(b - k)
                            }
                    }
                }
            }
            return dp[N][0][0]
        }

        fun set3(): Long {
            var dp = List(N + 1) { List(N) { MutableList(2) { 0L } } }
            for (d in 2..N) {
                for (l in 0..N - d) {
                    for (dir in 0..1) {
                        var (pl, b) = if (dir == 0) {
                            l to V[l + d - 1]
                        } else {
                            l + 1 to V[l]
                        }
                        dp[d][l][dir] = minOf(
                            dp[d - 1][pl][0] + ds(b - V[pl + d - 2]),
                            dp[d - 1][pl][1] + ds(b - V[pl]))
                    }
                }
            }
            return minOf(dp[N][0][0] + ds(V[N - 1]), dp[N][0][1] + ds(V[0]))
        }

        val res = set3()
        println("Case #${it}: $res")
    }
}
