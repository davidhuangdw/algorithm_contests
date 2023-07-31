package google_code_jam.`2018`.qua

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (_D, S) = split_input()
        val D = _D.toLong()
        var cnt = mutableListOf(0)
        for (ch in S) {
            if (ch == 'S')
                cnt[cnt.size-1]++
            else
                cnt.add(0)
        }
        var cur = cnt.withIndex().sumOf { (i, c) -> (1L shl i) * c }
        var move = 0
        var acc = 0
        var i = cnt.size - 1
        while (cur > D && i > 0) {
            acc += cnt[i]
            for (_j in 0 until acc) {
                if (cur > D) {
                    move++
                    cur -= 1L shl (i-1)
                } else break
            }
            i--
        }
        val res = if (cur > D) "IMPOSSIBLE" else move.toString()

        println("Case #${it}: $res")
    }
}
