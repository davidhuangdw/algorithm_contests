package google_code_jam.`2021`.r2

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }


    val M = 1000_000_007
    val F = MutableList(100_001) { 1L }
    val INV = MutableList(100_001) { 1L }
    val IF = MutableList(100_001) { 1L }
    for (i in 2..100_000) {
        INV[i] = (M - (M / i.toLong() * INV[M % i] % M)) % M
        F[i] = F[i - 1] * i % M
        IF[i] = IF[i - 1] * INV[i] % M
    }
    fun C(n: Int, m: Int): Long = F[n] * IF[m] % M * IF[n - m] % M

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val V = input_ints()

        var H = (1..N).firstOrNull { i -> (1 shl i) > N }!!
        var LM = List(N) { i -> MutableList(H) { i } }
        for (h in 1 until H) {
            for (i in 0..N - (1 shl h)) {
                val l = LM[i][h - 1]
                val r = LM[i + (1 shl (h - 1))][h - 1]
                LM[i][h] = if (V[l] < V[r]) l else r
            }
        }

        fun lastMinPos(l: Int, r: Int): Int {
            val d = r + 1 - l
            val h = (1..d).firstOrNull {i -> (1 shl i) > d }!! - 1
            val a = LM[l][h]
            val b = LM[r + 1 - (1 shl h)][h]
            return if (V[a] < V[b]) a else b
        }

        fun count(l: Int, r: Int, st: Int): Long {
            if (l > r) return 1
            if (l == r) return if (V[l] == st) 1 else 0


            val lastMin = lastMinPos(l, r)
            if (V[lastMin] != st) return 0
            return C(r - l, lastMin - l) * count(l, lastMin - 1, st) % M * count(lastMin + 1, r, st + 1) % M
        }

        val res = count(0, N - 1, 1)
        println("Case #${it}: $res")
    }
}
