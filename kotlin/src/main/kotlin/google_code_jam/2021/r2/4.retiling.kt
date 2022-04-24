package google_code_jam.`2021`.r2

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, C, F, S) = input_ints()
        val A = List(R) { readLine()!!.trim() }
        val B = List(R) { readLine()!!.trim() }

        fun set1(): Int {
            val P = List(2) { mutableListOf<Pair<Int, Int>>() }
            val ind = hashMapOf<Pair<Int, Int>, Int>()
            for (i in 0 until R)
                for (j in 0 until C)
                    if (A[i][j] != B[i][j]) {
                        val k = (i + j) % 2
                        P[k].add(i to j)
                        ind[i to j] = P[k].size - 1
                    }

            val n = P[0].size
            val m = P[1].size
            val adj = P[0].map {
                val a = mutableListOf<Int>()
                val (ci, cj) = it
                for ((i, j) in listOf(ci - 1 to cj, ci + 1 to cj, ci to cj - 1, ci to cj + 1)) {
                    if (i to j in ind && B[ci][cj] != B[i][j]) {
                        a.add(ind[i to j]!!)
                    }
                }
                a
            }

            fun maxMatch(): Int {
                var sum = 0
                val match = MutableList(m) { -1 }
                var vis = MutableList(m) { false }
                fun dfs(i: Int): Boolean {
                    for (j in adj[i]) {
                        if (!vis[j]) {
                            vis[j] = true
                            if (match[j] == -1 || dfs(match[j])) {
                                match[j] = i
                                return true
                            }
                        }
                    }
                    return false
                }
                for (i in 0 until n) {
                    vis = MutableList(m) { false }  // bug: forget to reset
                    if (dfs(i))
                        sum++
                }
                return sum
            }

            return P[0].size + P[1].size - maxMatch()
        }

        val res = set1()
        // todo: set2
        println("Case #${it}: $res")
    }
}
