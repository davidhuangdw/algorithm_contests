package google_kickstart.`2020`.E

import java.util.*

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val MAX = 1e12.toLong()
        val (N, M, S, R) = input_ints()
        val C = N * S
        val node = List(N + 1) { IntArray(S + 1) }
        val adj = List(C) { mutableListOf<Pair<Int, Int>>() }
        val trans = List(C) { mutableListOf<List<Int>>() }
        val que = ArrayDeque<Int>()
        val queued = hashSetOf<Int>()
        val cost = LongArray(C) { MAX }

        for (i in 1..N)
            for (j in 1..S)
                node[i][j] = S * (i - 1) + j - 1

        repeat(M) {
            val (a, b) = input_ints()
            for (s in 1..S) {
                val (u, v) = node[a][s] to node[b][s]
                adj[u].add(v to 1)
                adj[v].add(u to 1)
            }
        }
        for (i in 1..N) {
            val line = input_ints()
            for (s in line.subList(1, line.size)) {
                val u = node[i][s]
                cost[u] = 0
//                println("$u: ${cost[u]}")
                que.add(u)
                queued.add(u)
            }
        }
        repeat(R) {
            val line = input_ints()
            for (i in 1..N) {
                val tr = line.subList(1, line.size).map { s -> node[i][s] }
                for (u in tr.subList(0, tr.size - 1))
                    trans[u].add(tr)
            }
        }

        while (que.isNotEmpty()) {
            val u = que.removeFirst()
            for ((v, d) in adj[u]) {
                if (cost[u] + d < cost[v]) {
                    cost[v] = cost[u] + d
//                    println("$v: ${cost[v]}: $u -> $v")
                    if (v !in queued) {
                        queued.add(v)
                        que.add(v)
                    }
                }
            }
            for (tr in trans[u]) {
                val v = tr.last()
                val d = tr.subList(0, tr.size - 1).sumOf { cost[it] }
                if (d < cost[v]) {
                    cost[v] = d
//                    println("$v: ${cost[v]}: $tr")
                    if (v !in queued) {
                        queued.add(v)
                        if (cost[v] == cost[u])
                            que.addFirst(v)
                        else
                            que.add(v)
                    }
                }
            }
            queued.remove(u)
        }

        var res = (1..N).minOfOrNull { cost[node[it][1]] }!!
        if (res == MAX) res = -1

        println("Case #${it}: $res")
    }
}
