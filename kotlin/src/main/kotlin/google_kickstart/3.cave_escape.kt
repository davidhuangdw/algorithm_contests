import java.util.*

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val OB = -100000
    val DIR = listOf(0 to 1, 0 to -1, -1 to 0, 1 to 0)
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val line = input_ints()
        val (N, M, E) = line
        var (si, sj, ti, tj) = line.subList(3, line.size).map { it - 1 }
        val V = List(N) { input_ints() }

        fun testset1(): Int {
            val dp = List(N) { MutableList(M) { -1000000 } }
            val done = hashSetOf<Pair<Int, Int>>()
            dp[si][sj] = E
            val que = PriorityQueue<Pair<Int, Pair<Int, Int>>>(compareByDescending { it.first })
            que.add(E to (si to sj))

            fun valid(i: Int, j: Int) = i in 0 until N && j in 0 until M && i to j !in done && V[i][j] != OB
            while (que.isNotEmpty()) {
                val (x, pos) = que.poll()
                if (pos in done)
                    continue
                done.add(pos)
                val (i, j) = pos
                for ((di, dj) in DIR) {
                    val ni = i + di
                    val nj = j + dj
                    if (valid(ni, nj)) {
                        val nx = x + V[ni][nj]
                        if (nx > dp[ni][nj]) {
                            dp[ni][nj] = nx
                            que.add(nx to (ni to nj))
                        }
                    }
                }
            }
//        for(i in 0 until N)
//            println(dp[i])
            return maxOf(dp[ti][tj], -1)
        }

        // todo: testset2
        println("Case #${it}: ${testset1()}")
    }
}
