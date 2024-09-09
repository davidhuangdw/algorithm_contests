package atcoder.dp

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (N) = input_ints()
    val V = input_ints()
    val DP = Array(N) { LongArray(N) { 0 } }
    for (d in 0..N - 1)
        for (i in 0 until N - d) {
            val j = i + d
            DP[i][j] = if (d == 0) V[i] + 0L else maxOf(V[j] - DP[i][j - 1], V[i] - DP[i + 1][j])
        }

    println(DP[0][N-1])
}
