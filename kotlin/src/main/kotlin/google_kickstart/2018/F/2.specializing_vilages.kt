fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (V, E) = input_ints()
        val adj = List(V + 1) { hashMapOf<Int, Int>() }
        val zeros = mutableListOf<Int>()
        repeat(E) {
            val (a, b, l) = input_ints()
            if (l == 0) {
                zeros.add(a)
                zeros.add(b)
            }
            adj[a][b] = l
            adj[b][a] = l
        }
        val INF = 100_000_000
        val dis = List(V + 1) { i ->
            MutableList(V + 1) { j ->
                if (i == j) 0
                else adj[i].getOrDefault(j, INF)
            }
        }
        for (k in 1..V)
            for (i in 1..V)
                for (j in 1..V)
                    dis[i][j] = minOf(dis[i][j], dis[i][k] + dis[k][j])

        fun testset1(): Int {
            var min_sum = INF
            var cnt = 0
            for (bits in 0 until (1 shl V)) {
                val s = List(2) { hashSetOf<Int>() }
                for (i in 1..V) {
                    if ((1 shl (i - 1)) and bits > 0)
                        s[0].add(i)
                    else
                        s[1].add(i)
                }
                var sum = 0
                var connected = true
                for ((i, x) in s.withIndex()) {
                    for (a in x) {
                        var mi = INF
                        for (b in s[i xor 1])
                            mi = minOf(mi, dis[a][b])
                        if (mi == INF) {
                            connected = false
                            break
                        }
                        sum += mi
                    }
                }
                if (!connected) continue
                if (sum < min_sum) {
                    cnt = 1
                    min_sum = sum
                } else if (sum == min_sum)
                    cnt += 1
            }
            return cnt
        }

        fun testset2(): Long {
            val pa = MutableList(V + 1) { it }
            fun find(x: Int): Int {
                if (pa[x] != x)
                    pa[x] = find(pa[x])
                return pa[x]
            }

            var cnt = V
            fun union(x: Int, y: Int) {
                val rx = find(x)
                val ry = find(y)
                if (rx != ry) {
                    cnt -= 1
                    pa[rx] = ry
                }
            }

            val min_conn = MutableList(V + 1) { x ->
                if (x == 0) return@MutableList 0
                var y = if (x != 1) 1 else 2
                for (v in 1..V)
                    if (v != x && dis[x][v] < dis[x][y])
                        y = v
                if (y !in zeros)
                    union(x, y)         // don't connect zeros
                y
            }
            if (zeros.isNotEmpty())         // zeros pairs has 2 choices(cannot be the same)
                cnt -= 1

            return 1L shl cnt
        }

        println("Case #${it}: ${testset2()}")
    }
}
