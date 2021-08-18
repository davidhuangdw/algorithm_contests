package google_kickstart.`2018`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        val A = input_longs().sorted()
        var pre = A.scan(0L, Long::plus)
        var e = 1.0 * A.sum() / N
        repeat(K) {
            var (l, r) = 0 to N - 1
            while (l <= r) {
                val md = (l + r) / 2
                if (A[md] <= e) l = md + 1
                else r = md - 1
            }
            e = 1.0 * l / N * e + 1.0 * (pre[N] - pre[l]) / N
        }
        println("Case #${it}: $e")
    }
}
