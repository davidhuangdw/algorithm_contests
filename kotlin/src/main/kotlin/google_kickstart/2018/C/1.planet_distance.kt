package google_kickstart.`2018`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val adj = List(N + 1) { mutableListOf<Int>() }
        repeat(N) {
            val (x, y) = input_ints()
            adj[x].add(y)
            adj[y].add(x)
        }

        val path = mutableListOf<Int>()
        val pset = hashSetOf<Int>()
        var done = false
        val dis = MutableList(N + 1) { -1 }
        val que = ArrayDeque<Int>()
        fun dfs(u: Int, p: Int) {
            path.add(u)
            pset.add(u)
            for (v in adj[u]) {
                if (v == p) continue
                if (v in pset) { // cycle
                    do {
                        val x = path.removeLast()
                        dis[x] = 0
                        que.add(x)
                    } while (x != v)
                    done = true
                    return
                }
                dfs(v, u)
                if (done) return
            }
            pset.remove(u)
            path.removeLast()
        }
        dfs(1, -1)

        // bfs
        while (que.isNotEmpty()) {
            val u = que.removeFirst()
            for (v in adj[u]) {
                if (dis[v] == -1) {
                    dis[v] = dis[u] + 1
                    que.add(v)
                }
            }
        }
        println("Case #${it}: ${(1..N).map { dis[it] }.joinToString(" ")}")
    }
}
