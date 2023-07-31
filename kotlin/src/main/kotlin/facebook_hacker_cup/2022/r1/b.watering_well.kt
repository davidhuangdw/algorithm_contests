fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val M = 1_000_000_007L
    fun sum(a: List<Long>): Long {
        var s = 0L
        for (v in a)
            s = (s + v % M) % M
        return s
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val A = (1..N).map { input_longs() }
        val (Q) = input_ints()
        val B = (1..Q).map { input_longs() }
        var s = 0L
        for (i in 0..1) {
            val a = A.map { it[i] }
            val aq = a.map { it * it % M }
            val b = B.map { it[i] }
            val bq = b.map { it * it % M }
            s = sum(listOf(s, Q * sum(aq), N * sum(bq), (M-sum(a)) * sum(b) * 2))
        }
        println("Case #${it}: $s")
    }
}
