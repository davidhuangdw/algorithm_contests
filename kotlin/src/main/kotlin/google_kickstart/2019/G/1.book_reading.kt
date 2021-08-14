package google_kickstart.`2019`.G

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, M, Q) = input_ints()
        val ms = input_ints().toHashSet()
        val count = hashMapOf<Int, Int>()
        for (r in input_ints())
            count[r] = count.getOrDefault(r, 0) + 1

        var sum = 0L
        for ((r, c) in count) {
            if (r == 1)
                sum += (N - M) * 1L * c
            else {
                for (x in r..N step r)
                    if (x !in ms)
                        sum += c
            }
        }

        println("Case #${it}: $sum")
    }
}
