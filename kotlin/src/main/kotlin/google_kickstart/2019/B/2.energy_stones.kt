package google_kickstart.`2019`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        data class Stone(val s: Int, val e: Int, val l: Int) : Comparable<Stone> {
            override fun compareTo(other: Stone) = compareValues(s * other.l, other.s * l)
        }
        val (N) = input_ints()
        val ss = mutableListOf<Stone>()
        var sum = 0
        repeat(N) {
            val (s, e, l) = input_ints()
            ss.add(Stone(s, e, l))
            sum += s
        }

        sum += 100
        ss.sort()
        val max_energy = MutableList(sum + 1) { 0 }
        for (i in N - 1 downTo 0) {
            val (s, e, l) = ss[i]
            for (t in 0..sum - s)
                max_energy[t] = maxOf(max_energy[t], max_energy[t + s] + maxOf(0, e - t * l))
        }
        println("Case #${it}: ${max_energy[0]}")
    }
}
