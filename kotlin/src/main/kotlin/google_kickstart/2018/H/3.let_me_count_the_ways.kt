fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val U = 200_001
    val MOD = 1000_000_007
    val P = MutableList(U) { 1L }
    val F = MutableList(U) { 1L }
    val IF = MutableList(U) { 1L }
    val INV = MutableList(U) { 1L }
    P[1] = 2L
    for (x in 2 until U) {
        P[x] = P[x - 1] * 2 % MOD
        F[x] = F[x - 1] * x % MOD
        INV[x] = (MOD - (MOD / x) * INV[MOD % x] % MOD) % MOD
        IF[x] = IF[x - 1] * INV[x] % MOD
    }
    fun C(n: Int, k: Int) = F[n] * IF[n - k] % MOD * IF[k] % MOD
    fun mul(vararg xs: Long) = xs.fold(1L) { a, b -> a * b % MOD }
//    for(k in 0..5)
//        println("$k: ${C(6, k)}")

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, M) = input_ints()
        var sum = 0L
        var neg = 1
        for (k in 0..M) {
            sum = (sum + neg * mul(C(M, k), C(2 * N - k, k), P[k], F[k], F[2 * N - 2 * k]) + MOD) % MOD
            neg *= -1
        }
        println("Case #${it}: $sum")
    }
}
