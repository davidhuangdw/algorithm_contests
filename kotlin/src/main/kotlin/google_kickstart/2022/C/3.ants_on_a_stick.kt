package google_kickstart.`2022`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (N, L) = input_ints()
        val A = (1..N).map { input_ints() }
        fun set3() {
            val ind = (0 until N).sortedBy { A[it][0] }
            val f = A.map { (p, d) ->
                val t = if (d == 0) p else L - p
                t to d
            }.sortedBy { it.first }
            var (l, r) = 0 to N - 1
            val ord = mutableListOf<Pair<Int, Int>>()
            for ((t, d) in f) {
                if (d == 0) {
                    ord.add(t to ind[l] + 1)
                    l++
                } else {
                    ord.add(t to ind[r] + 1)
                    r--
                }
            }
            val line = ord.sortedWith(compareBy({ it.first }, { it.second })).map { it.second }.joinToString(" ")
            println(line)
        }
        print("Case #${it}: ")
        set3()
    }
}
