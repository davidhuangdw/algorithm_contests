package google_kickstart.`2017`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val M = 1000_000_007
    val P = (1..10_000).scan(1L) { acc, _ -> acc * 2 % M }
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val A = input_ints()
        var res = 0L
        for ((i, a) in A.withIndex()) {
            res = (res + a * (P[i] + M - P[N - i - 1]) % M) % M
        }

        println("Case #${it}: $res")
    }
}
