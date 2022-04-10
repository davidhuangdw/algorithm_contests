package google_code_jam.`2022`.`1a`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (E, W) = input_ints()
        val X = List(E) { input_ints() }
        fun set1(): Int {
            fun comb(ws: List<Int>): MutableList<List<Int>> {
                var res = mutableListOf<List<Int>>()
                var cur = mutableListOf<Int>()
                fun swap(k: Int, j: Int) {
                    val t = cur[j]
                    cur[j] = cur[k]
                    cur[k] = t
                }
                for ((i, k) in ws.withIndex()) {
                    (1..k).forEach { _ ->
                        cur.add(i)
                    }
                }
                val n = cur.size
                fun dfs(j: Int) {
                    if (j == n) {
                        res.add(cur.toList())
                    } else {
                        var done = MutableList(W) { false }
                        (j until n).forEach { k ->
                            if (!done[cur[k]]) {
                                done[cur[k]] = true
                                swap(j, k)
                                dfs(j + 1)
                                swap(j, k)  // bug: forget to revert when backtrace
                            }
                        }
                    }
                }
                dfs(0)
                return res
            }

            var pre = listOf(listOf<Int>() to 0)
            for (ws in X) {
                pre = comb(ws).map { y ->
                    var min = Int.MAX_VALUE
                    for ((x, k) in pre) {
                        var i = 0
                        while (i < minOf(x.size, y.size) && x[i] == y[i]) {
                            i++
                        }
                        min = minOf(min, x.size - i + y.size - i + k)
                    }
                    y to min
                }
            }
            return pre.minOf { it.second } + pre[0].first.size
        }

        fun set2(): Int {
            val pre = List(E) { MutableList(E) { 0 } }
            for (i in 0 until E) {
                val remain = X[i].toMutableList()
                pre[i][i] = remain.sum()
                for (j in i + 1 until E) {
                    for ((k, w) in X[j].withIndex()) {
                        if (w < remain[k]) {
                            remain[k] = w
                        }
                    }
                    pre[i][j] = remain.sum()
                }
            }

            val dp = List(E) { MutableList(E) { -1 } } // (l, r): given pre[l,r], cost to transfer from l to r
            fun f(l: Int, r: Int): Int {
                if (dp[l][r] >= 0) return dp[l][r]
                var min = Int.MAX_VALUE
                if (l == r) {
                    min = 0
                } else {
                    for (x in l until r) { // bug: not l..r
                        min = minOf(
                            min,
                            f(l, x) + 2 * (pre[l][x] - pre[l][r]) + f(x + 1, r) + 2 * (pre[x + 1][r] - pre[l][r])
                        )
                    }
                }
                dp[l][r] = min
                return min
            }
            return f(0, E - 1) + 2 * pre[0][E - 1]
        }

        val res = set2()
        println("Case #${it}: $res")
    }
}
