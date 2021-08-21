package google_kickstart.`2018`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val MOD = 1000_000_000 + 7
    val inv = MutableList(1000_000 + 1) { 1L }
    for (x in 2..1000_000)
        inv[x] = (MOD - 1L * MOD / x * inv[MOD % x] % MOD) % MOD

    fun pow(b: Int, k: Int): Long {
        var mul = b % MOD * 1L
        var p = 1L
        var x = k
        while (x > 0) {
            if (x and 1 > 0) {
                p = p * mul % MOD
            }
            mul = mul * mul % MOD
            x /= 2
        }
        return p
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val line = input_ints()
        val (N, K, X, Y) = line
        val (C, D, E1, E2, F) = line.subList(4, line.size)
        var (x, y) = X.toLong() to Y.toLong()

        val A = MutableList(N) { 0L }
        repeat(N) { i ->
            A[i] = (x + y) % F
            x = ((C * x + D * y + E1) % F).also {
                y = (D * x + C * y + E2) % F
            }
        }

        fun power_sum(b: Int): Long {  //b^1 + b^2 + .. + b^K
            if (b == 1) return K.toLong()
            return (MOD + (pow(b, K + 1) - b) * inv[b - 1] % MOD) % MOD
        }

        var sum = 0L
        var pre = 0L
        for ((i, v) in A.withIndex()) {
            pre = (pre + power_sum(i + 1)) % MOD
            sum = (sum + v * (N - i) % MOD * pre % MOD) % MOD   // bug: triple mul
//            println("$i $pre $sum")
        }

        println("Case #${it}: $sum")
    }
}
