package google_kickstart.`2017`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, M, K) = input_ints()
        val G = List(N) {
            val (row) = split_input()
            row.map { it == '#' }
        }

        fun testset2(): Int {
            var height = List(N + 1) { MutableList(M) { 0 } }
            for (i in N - 1 downTo 0)
                for (j in 0 until M) {
                    if (!G[i][j]) {
                        continue
                    }
                    height[i][j] = 1
                    if (j - 1 >= 0 && j + 1 < M) {
                        height[i][j] += ((j - 1)..(j + 1)).minOf { height[i + 1][it] }
                    }
                }

            var dp = List(N + 1) { i -> MutableList(M) { j -> height[i][j] * height[i][j] } }
            for (_t in 2..K) {
                var cur = List(N + 1) { MutableList(M) { 0 } }
                for (i in 0 until N)
                    for (j in 0 until M) {
                        for (h in 1..height[i][j]) { // bug: greedily use single[i][j] is wrong!!
                            val (l, r) = j - (h - 1) to j + (h - 1)
                            val sub = (l..r).maxOf { dp[i + h][it] }
                            if (sub == 0)
                                continue
                            cur[i][j] = maxOf(cur[i][j], h * h + sub)
                        }
                    }
                dp = cur
            }
            return dp.maxOf { it.maxOrNull()!! }
        }
        println("Case #${it}: ${testset2()}")
    }
}
