package google_code_jam.`2018`.r2

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    // set2: trick: precompute for all
    val N = 35
    val M = 501
    val dp = List(N + 1) { List(M) { MutableList(M) { 0 } } }
    for (i in 0..N)
        for (j in 0 until M) {
            for (k in 0 until M) {
                for (t in 0..M) {
                    val dj = i * t
                    val dk = t * (t - 1) / 2
                    if (j - dj < 0 || k - dk < 0)
                        break
                    dp[i][j][k] = maxOf(dp[i][j][k], (if(i-1 >= 0)dp[i-1][j - dj][k - dk] else 0) + t)
                }
            }
        }
    val (T) = input_ints()
    (1..T).forEach {
        val (R, B) = input_ints()

        fun set1(): Int {
            val cache = hashMapOf<List<Int>, Int>()
            val choose = hashMapOf<List<Int>, Int>()
            fun dp(i: Int, n: Int, m: Int): Int {
                if (i > n) return 0
                val id = listOf(i, n, m)
                if (id !in cache) {
                    var max = 0
                    var ch = -1
                    for (k in 0..m) {
                        val dn = (k + 1) * i
                        val dm = k * (k + 1) / 2
                        if (dn > n || dm > m) break
                        val cur = dp(i + 1, n - dn, m - dm) + k + 1
                        if (cur > max) {
                            max = cur
                            ch = k
                        }
                    }
                    cache[id] = max
                    choose[id] = ch
                }
                return cache[id]!!
            }
            return dp(0, R, B) - 1
        }

        fun set2(): Int {
            return dp[N][R][B] - 1
        }

        val r = set2()
        println("Case #${it}: $r")
    }
}
