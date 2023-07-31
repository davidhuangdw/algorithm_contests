package google_kickstart.`2022`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    val MOD = 1000_000_007
    val M = 400
    val fact = MutableList(M + 1) { 1L }
    val inv = MutableList(M + 1) { 1L }
    val inv_fact = MutableList(M + 1) { 1L }
    for (i in 2..M) {
        fact[i] = fact[i - 1] * i % MOD
        inv[i] = (MOD - MOD / i) * inv[MOD % i] % MOD
        inv_fact[i] = inv_fact[i - 1] * inv[i] % MOD
    }
    (1..T).forEach {
        val (N) = input_ints()
        val S = split_input()[0]
        fun set1(): Long {
            fun isPar(s: String): Boolean {
                val n = s.length
                var (l, r) = 0 to n - 1
                while (l < r) {
                    if (s[l] != s[r])
                        return false
                    l++
                    r--
                }
                return true
            }

            var res = 0L
            for (i in 0 until (1 shl minOf(N, 10)) - 1) {
                val s = mutableListOf<Char>()
                for (j in 0 until N) {
                    if ((1 shl j) and i > 0) {
                        s.add(S[j])
                    }
                }
                val n = s.size
                if (isPar(s.joinToString(""))) {
                    res = (res + fact[n] * fact[N - n] % MOD) % MOD
                }
            }
            res = res * inv_fact[N] % MOD
            return res
        }

        fun set2(): Long {
            var res = fact[0] * fact[N] % MOD
            res = (res + N * fact[1] * fact[N - 1] % MOD) % MOD
            val dp = List(N) { l ->
                List(N) { i ->
                    MutableList(N) { j ->
                        if (l == 0) 1L
                        else if (l == 1)
                            maxOf(0L, j - i + 1L)
                        else 0L
                    }
                }
            }
            for (l in 2 until N) {
                val cur = dp[l]
                for (d in l..N) {
                    for (i in 0 until N + 1 - d) {
                        val j = i + d - 1
                        if (S[i] == S[j])
                            cur[i][j] += if (i + 1 <= j - 1) dp[maxOf(l - 2, 0)][i + 1][j - 1] else 1
                        cur[i][j] = (cur[i][j] + cur[i + 1][j] + cur[i][j - 1] - cur[i + 1][j - 1]) % MOD
                    }
                }
                res = (res + cur[0][N - 1] * fact[l] % MOD * fact[N - l] % MOD) % MOD
            }
            res = res * inv_fact[N] % MOD
            return res
        }

        val r = set2()
        println("Case #${it}: $r")
    }
}
