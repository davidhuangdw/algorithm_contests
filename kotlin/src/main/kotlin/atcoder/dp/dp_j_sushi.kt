package atcoder.dp

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (N) = input_ints()
    val all = input_ints()
    val (X, Y, Z) = (1..3).map { v -> all.count { it == v } }
    val DP = Array(N + 1) { Array(N + 1) { DoubleArray(N + 1) { -1.0 } } }
    DP[0][0][0] = 0.0

    fun dp(a: Int, b: Int, c: Int): Double {
        if (a < 0 || b < 0 || c < 0) return .0
        val s = a + b + c
        if (DP[a][b][c] < 0) {
            DP[a][b][c] = (N + dp(a - 1, b, c) * a + dp(a + 1, b - 1, c) * b + dp(a, b + 1, c - 1) * c) / s
        }
        return DP[a][b][c]
    }
    println(dp(X, Y, Z))
}

//    for (s in 1 ..N)
//        for (a in s downTo 0)
//            for (c in 0 ..minOf(Z, s-a)) {
//                val b = s - a - c
//                if(b > Y+Z) continue
//                DP[a][b][c] = (1 + (if (a - 1 >= 0) DP[a - 1][b][c] * a / N else .0) +
//                        (if (b - 1 >= 0) DP[a + 1][b-1][c] * b / N else .0) +
//                        (if (c - 1 >= 0) DP[a][b + 1][c - 1] * c / N else .0)) * N / s
////                println("$s: $a $b $c: ${DP[a][b][c]}")
//            }
//    println(DP[X][Y][Z])
//}
