package google_kickstart.`2019`.Practice

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (A, B) = input_ints()
        input_ints()
        var (l, r) = A + 1 to B
        while (l <= r) {
            val md = (l + r) ushr 1
            println(md)
            val resp = readLine()!!
            if ("CORRECT" in resp)
                break
            if ("WRONG" in resp)
                break
            if ("SMALL" in resp)
                l = md + 1
            else r = md - 1
        }
//        println("Case #${it}: ")
    }
}
