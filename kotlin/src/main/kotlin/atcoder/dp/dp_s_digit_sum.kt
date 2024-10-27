package atcoder.dp

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val MOD = 1e9.toInt() + 7
    val KS = split_input()[0].map { it.digitToInt() }.reversed()
    val (D) = input_ints()
    val N = KS.size
    val dp = List(N + 1) { List(D) { MutableList(2) { 0L } } }
    for (i in 0 until N)
        for (d in 0 until D)
            for (t in 0 until 2) {
                dp[i][d][t] = (0 until 10).sumOf { x ->
                    val y = (d + 10*D - x) % D
//                    println("----- ${i} ${d} ${t}: ${x} ${y}")
                    if (t == 1 && x > KS[i]) {
                        0L
                    } else if (t == 1 && x == KS[i]) {
                        if(i==0) (if(y==0) 1 else 0) else dp[i-1][y][1]
                    }else {
                        if(i==0)(if(y==0) 1 else 0) else dp[i-1][y][0]
                    }
                } % MOD
//                println("----- ${i} ${KS[i]} ${d} ${t}: ${dp[i][d][t]}")
            }

    println((dp[N-1][0][1] + MOD -1) % MOD)
}
