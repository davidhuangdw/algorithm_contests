package google_kickstart.`2021`.acpc

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val S = readLine()!!.trim()
        val n = S.length
        val pre = S.scan(0L) { acc, ch -> acc + if (ch == 'B') 1 else 0 }
        var (l, r) = input_longs()
        l -= 1
        r -= 1
        val (i, j) = l / n to r / n
        var (oi, oj) = (l % n).toInt() to (r % n).toInt()
        val cnt = pre[oj + 1] + (j - i) * pre.last() - pre[oi]

        println("Case #${it}: $cnt")
    }
}
