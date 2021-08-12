package google_kickstart.`2019`.F

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        val A = input_ints()
        var m = 0
        val vid = hashMapOf<Int, Int>()
        val id = MutableList(N) { 0 }
        for ((i, v) in A.withIndex()) {
            if (v !in vid) {
                vid[v] = m
                m += 1
            }
            id[i] = vid[v]!!
        }

        fun method1(): Int {      // O(n*n*k)
            val dp = List(2) { List(K + 1) { MutableList(m) { 0 } } }
            for ((i, id_i) in id.withIndex()) {
                val pre = dp[i and 1 xor 1]
                val cur = dp[i and 1]
                for (k in 0..K) {
                    val pre_level_min = if (k == 0) Int.MAX_VALUE else pre[k - 1].minOrNull()!!
                    for (x in 0 until m)
                        cur[k][x] = (if (id_i == x) 0 else 1) + minOf(pre_level_min, pre[k][x])
                }
            }
            return dp[(N - 1) and 1][K].minOrNull()!!
        }

        fun method2(): Int {    // O(n*n *log(n)), todo
            return 1
        }

        println("Case #${it}: ${method1()}")
    }
}
