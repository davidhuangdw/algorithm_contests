fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        var (N, D, C, M) = input_ints()
        val line = readLine()!!
        fun check(): String {
            var (d, c) = D.toLong() to C.toLong()
            var last = N
            for ((i, ch) in line.withIndex()) {
                if (ch == 'D') {
                    d -= 1
                    c += M
                } else {
                    c -= 1
                }
                if (d < 0 || c < 0) {
                    last = i
                    break
                }
            }
            for (i in last until N)
                if (line[i] == 'D')
                    return "NO"
            return "YES"
        }
        println("Case #${it}: ${check()}")
    }
}
