package google_kickstart.`2019`.H

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val A = (1..N).map { readLine()!!.trim() }
        var nid = 0
        val id = hashMapOf<Pair<Int, Int>, Int>()
        val adj = List(4 * N) { mutableListOf<Pair<Int, Int>>() }
        for (i in 0 until N)
            for ((j, v) in A[i].withIndex()) {
                val x = 0 to i + j
                val y = 1 to i - j
                if (x !in id) {
                    id[x] = nid
                    nid += 1
                }
                if (y !in id) {
                    id[y] = nid
                    nid += 1
                }
                val (ix, iy) = id[x]!! to id[y]!!
                val diff = if (v == '#') 0 else 1
                adj[ix].add(iy to diff)
                adj[iy].add(ix to diff)
            }

        var sum = 0
        val color = hashMapOf<Int, Int>()
        var tol = 0
        var cnt = 0
        fun dfs(u: Int, c: Int) {
            if (u in color) return
            color[u] = c
            cnt += c
            tol += 1
            for ((v, diff) in adj[u])
                dfs(v, c xor diff)
        }
        for (i in 0 until nid)
            if (i !in color) {
                tol = 0
                cnt = 0
                dfs(i, 0)
                sum += minOf(cnt, tol - cnt)
            }

        println("Case #${it}: $sum")
    }
}
