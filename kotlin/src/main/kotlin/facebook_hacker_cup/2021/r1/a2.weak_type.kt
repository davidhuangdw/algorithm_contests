fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val MOD = 1000_000_007
        val (N) = input_ints()
        val s = readLine()!!
        val ed_sum = MutableList(s.length + 1) { 0L }

        var pre = 'F'
        var ed = 0
        var last_ed = 0
        var sum = 0L
        for ((i, ch) in s.withIndex()) {
            val l = i + 1
            if (ch in "XO") {
                if (ch != pre) {
                    last_ed = ed
                }
                ed = l
            }
            ed_sum[l] = (ed_sum[last_ed] + last_ed) % MOD
            sum = (sum + ed_sum[l]) % MOD
            if (ch in "XO")
                pre = ch
        }
        println("Case #${it}: $sum")
    }
}
