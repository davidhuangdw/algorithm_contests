package google_kickstart.`2022`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (I) = split_input()
        val (P) = split_input()
        val m = P.length
        var j = 0
        var fail = false
        for (a in I) {
            while (j < m && P[j] != a)
                j++
            if (j >= m) {
                fail = true
                break
            } else {
                j++
            }
        }
        val res = if (fail) "IMPOSSIBLE" else "${P.length - I.length}"
        println("Case #${it}: $res")
    }
}
