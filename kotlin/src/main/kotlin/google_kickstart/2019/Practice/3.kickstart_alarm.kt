package google_kickstart.`2019`.Practice

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val MOD = (1e9 + 7).toLong()

    val MAXN = (1e6 + 1).toInt()
    val INV = LongArray(MAXN) { 1L }
    for (i in 2 until MAXN)
        INV[i] = (MOD - (MOD / i) * INV[(MOD % i).toInt()] % MOD) % MOD

    fun powMod(p: Long, k: Int): Long {
        var u = p
        var res = 1L
        var r = k
        while (r > 0) {
            if (r and 1 == 1) {
                res = res * u % MOD
            }
            u = u * u % MOD
            r /= 2
        }
        return res
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val line = input_ints()
        val (N, K, X1, Y1, C) = line.subList(0, 5)
        val (D, E1, E2, F) = line.subList(5, line.size)

        val A = MutableList(N) { 0L }
        var (x, y) = X1.toLong() to Y1.toLong()
        repeat(N) {
            A[it] = (x + y) % F
            x = (C * x + D * y + E1) % F.also {
                y = (D * x + C * y + E2) % F
            }
        }
//        println(A)

        val SK = LongArray(N + 1)
        SK[1] = K.toLong()
        for (q in 2..N) {
            SK[q] = (powMod(q * 1L, K + 1) + MOD - q) % MOD * INV[q - 1] % MOD + SK[q - 1]
            SK[q] %= MOD
//            println("$q $K: ${SK[q]}")
        }

        var res = 0L
        for (i in 1..N) {
            val a = A[i - 1]
            res += a * (N + 1 - i) % MOD * SK[i] % MOD
            res %= MOD
//            println("---$i $a: $a * $sub -- $res")
        }

        println("Case #${it}: $res")
    }
}
