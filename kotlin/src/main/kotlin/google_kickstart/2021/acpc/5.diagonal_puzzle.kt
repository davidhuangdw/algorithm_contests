package google_kickstart.`2021`.acpc

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val A = (1..N).map { readLine()!! }
        var nid = 0
        val id = hashMapOf<Pair<Int, Int>, Int>()
        val adj = List(N * 4) { mutableListOf<Pair<Int, Int>>() }
        for (i in 0 until N)
            for (j in 0 until N) {
                val l = 0 to i + j
                val r = 1 to i - j
                for (p in listOf(l, r)) {
                    if (p !in id) {
                        id[p] = nid
                        nid += 1
                    }
                }
                val (ld, rd) = id[l]!! to id[r]!!
                val diff = if (A[i][j] == '#') 0 else 1
                adj[ld].add(rd to diff)
                adj[rd].add(ld to diff)
            }

        var sum = 0
        val color = hashMapOf<Int, Int>()
        var total = 0
        var cnt = 0
        fun dfs(u: Int, c: Int) {
            if (u in color) return
            color[u] = c
            total += 1
            cnt += c
            for ((v, diff) in adj[u])
                dfs(v, c xor diff)
        }
        for (u in 0 until nid)
            if (u !in color) {
                total = 0
                cnt = 0
                dfs(u, 0)
                sum += minOf(cnt, total - cnt)
            }
        println("Case #${it}: $sum")
    }
}
