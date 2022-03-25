package google_kickstart.`2022`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = split_input()
        var x = 0
        for (ch in N) {
            x = (x + (ch - '0')) % 9
        }
        x = (9 - x) % 9

        var xc = '0' + x
        val n = N.length
        var k = n
        for ((i, ch) in N.withIndex()) {
            if (i == 0 && xc == '0')     // bug: not skip the leading-zero case
                continue
            if (ch > xc) {
                k = i
                break
            }
        }

        println("Case #${it}: ${N.substring(0, k) + xc + N.substring(k)}")
    }
}
