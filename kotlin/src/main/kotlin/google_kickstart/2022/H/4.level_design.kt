package google_kickstart.`2022`.H

import java.util.Deque

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (N) = input_ints()
        val P = input_ints()
//        println("------- $N ${P.joinToString(" ")}")
        val done = MutableList(N) { false }
        val cycles = mutableListOf<Int>()
        fun dfs(u: Int): Int {
            if (done[u]) return 0
            done[u] = true
            return dfs(P[u] - 1) + 1
        }
        for (i in 0 until N)
            if (!done[i]) {
                cycles.add(dfs(i))
            }
        cycles.sortDescending()

        fun set1(): List<Int> {
            val dp = MutableList(N + 1) { 1e8.toInt() }
            dp[0] = 0
            val exist = hashSetOf(0)  // bug: declared as mutableList, not hashSet
            for (v in cycles) {
                for (x in exist.toList()) {
                    exist.add(v + x)
                    dp[v + x] = minOf(dp[v + x], dp[x] + 1)
                }
            }
            return dp
        }

        fun set2(): List<Int>{
            val cnt = MutableList(N + 1) { 0 }
            for (v in cycles) {
                cnt[v]++
            }
            var dp = MutableList(N + 1) { 1e8.toInt() }
            dp[0] = 0
            for ((v, c) in cnt.withIndex()) {
                if (v == 0 || c == 0) continue
                var cur = MutableList(N + 1) { dp[it] }  // bug: need reuse old for 0 case
                for (st in 0 until v) {
                    val win = ArrayDeque<Int>()
                    for (i in 0..(N - st) / v) {
                        // deque
                        if (win.isNotEmpty() && i - win.first() > c)
                            win.removeFirst()
                        // compute
                        if (win.isNotEmpty())
                            cur[st + i * v] = minOf(cur[st + i * v], dp[st + win.first() * v] - win.first() + i)
                        // enque
                        while (win.isNotEmpty() && dp[st + win.last() * v] - win.last() >= dp[st + i * v] - i)
                            win.removeLast()
                        win.addLast(i)  // bug: forget add
                    }
                }
                dp = cur
            }
            return dp
        }

        val dp = set2()
        val acc = cycles.scan(0, Int::plus)
        fun bfs(x: Int): Int {
            var (l, r) = 0 to acc.size - 1
            while (l <= r) {
                val m = (l + r) / 2
                if (acc[m] >= x)
                    r = m - 1
                else l = m + 1
            }
            return l
        }

        val res = (1..N).map {
            minOf(bfs(it), dp[it] - 1)
        }.joinToString(" ")

        println("Case #${it}: $res")
    }
}
