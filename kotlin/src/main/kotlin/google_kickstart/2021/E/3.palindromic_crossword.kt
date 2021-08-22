fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, M) = input_ints()
        val A = (0 until N).map { readLine()!!.trim().toMutableList() }
        val rows = hashMapOf<Pair<Int, Int>, Pair<Int, Int>>()
        val cols = hashMapOf<Pair<Int, Int>, Pair<Int, Int>>()
        val que = mutableListOf<Pair<Int, Int>>()
        for (i in 0 until N) {
            var pre = 0
            for (j in 0..M) {
                if (j == M || A[i][j] == '#') {
                    var (l, r) = pre to j - 1
                    while (l <= r) {
                        rows[i to l] = i to r
                        rows[i to r] = i to l
                        l += 1
                        r -= 1
                    }
                    pre = j + 1
                } else if (j != M && A[i][j] != '.') {
                    que.add(i to j)
                }
            }

        }
        for (j in 0 until M) {
            var pre = 0
            for (i in 0..N) {
                if (i == N || A[i][j] == '#') {
                    var (l, r) = pre to i - 1
                    while (l <= r) {
                        cols[l to j] = r to j
                        cols[r to j] = l to j
                        l += 1
                        r -= 1
                    }
                    pre = i + 1
                }
            }
        }

        var cnt = 0
        while (que.isNotEmpty()) {
            val pos = que.removeLast()
            val (i, j) = pos
            for (other in listOf(rows[pos]!!, cols[pos]!!))
                if (other != pos) {
                    val (ti, tj) = other
                    if (A[ti][tj] == '.') {
                        cnt += 1
                        A[ti][tj] = A[i][j]
                        que.add(other)
                    }
                }
        }

        println("Case #${it}: $cnt")
        for (i in 0 until N)
            println(A[i].joinToString(""))
    }
}
