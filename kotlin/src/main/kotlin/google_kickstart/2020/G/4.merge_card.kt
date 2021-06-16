package google_kickstart.`2020`.G

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }
    fun input_longs() = split_input().map { it.toLong() }

    val MAX_N = 5001
//    val dp = List(MAX_N) { DoubleArray(MAX_N) { .0 } }
//    for (n in 2 until MAX_N)
//        for (i in 0 until n) {
//            if (i > 0)
//                dp[n][i] += (dp[n - 1][i - 1] * i + 1) / (n - 1)
//            if (i + 1 < n)
//                dp[n][i] += (dp[n - 1][i] * (n - 1 - i) + 1) / (n - 1)
//        }
//

    val H = (1..MAX_N).scan(.0){p, v -> p + 1.0/v}
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (n) = input_ints()
        val a = input_longs()

//        var res = (0 until n).sumOf { i -> a[i] * dp[n][i] }
        var res = (0 until n).sumOf { i -> a[i] * (H[i] + H[n-1-i]) }
        println("Case #${it}: ${res}")


//        val pre = a.scan(.0, Double::plus)
//        val e = MutableList(n) { MutableList(n) { .0 } }
//        val pe = MutableList(n) { MutableList(n) { .0 } }
//        val se = MutableList(n) { MutableList(n) { .0 } }
//
//
//        for (d in 1 until n) {
//            for (i in 0 until n - d) {
//                val j = i + d
////                e[i][j] = (i until j).sumOf { k ->
////                    e[i][k] + e[k + 1][j] + pre[j + 1] - pre[i]
////                } / d
//                e[i][j] = (pe[i][j - 1] + se[i + 1][j]) / d + pre[j + 1] - pre[i]
//                pe[i][j] = pe[i][j - 1] + e[i][j]
//                se[i][j] = e[i][j] + se[i + 1][j]
//            }
//        }
//        println("Case #${it}: ${e[0][n - 1]}")
    }
}
