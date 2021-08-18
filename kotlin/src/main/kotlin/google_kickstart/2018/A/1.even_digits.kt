package google_kickstart.`2018`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_longs()
        var diff = 0L

        var (i, b) = 0 to 1L
        var r = 0L
        while (i <= 16) {
            val x = N / b % 10
            if (x % 2L == 1L) {
                val part = N % (b * 10)
                val up = (x + 1) * b
                val down = (x - 1) * b + r
                diff = if (x < 9) minOf(up - part, part - down) else part - down
            }
            b *= 10L
            i += 1
            r = r * 10L + 8
        }

        println("Case #${it}: $diff")
    }
}
