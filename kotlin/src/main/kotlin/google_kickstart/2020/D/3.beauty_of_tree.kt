package google_kickstart.`2020`.D

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, A, B) = input_ints()
        val line = input_ints()

        val ch = List(N + 1) { mutableListOf<Int>() }
        for ((i, par) in line.withIndex()) {
            val u = i + 2
            ch[par].add(u)
        }

        val path = mutableListOf<Int>()
        val sub_a = DoubleArray(N + 1)
        val sub_b = DoubleArray(N + 1)

        var res = .0
        fun calc(u: Int) {
            path.add(u)
            for (v in ch[u]) {
                calc(v)
            }
            sub_a[u] += 1.0 / N
            sub_b[u] += 1.0 / N
            res += sub_a[u] + sub_b[u] - sub_a[u] * sub_b[u]

            val pa = path.size - 1 - A
            if (pa >= 0) sub_a[path[pa]] += sub_a[u]
            val pb = path.size - 1 - B
            if (pb >= 0) sub_b[path[pb]] += sub_b[u]
            path.removeLast()
        }

        calc(1)
        println("Case #${it}: $res")
    }
}
