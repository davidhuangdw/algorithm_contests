fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val s = readLine()!!
        var pre = 'F'
        var cnt = 0
        for (ch in s)
            if (ch in "XO") {
                if (ch != pre)
                    cnt += 1
                pre = ch
            }
        if (cnt > 0)
            cnt -= 1
        println("Case #${it}: $cnt")
    }
}
